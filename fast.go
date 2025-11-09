package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Browsers []string
	Company  string
	Country  string
	Email    string
	Job      string
	Name     string
	Phone    string
}

func FastSearch(out io.Writer) {

	users, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	var (
		seenBrowsers   []string
		uniqueBrowsers int
		foundUsers     string
		ok             bool
	)

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers
		if browsers == nil {
			//log.Println("cant cast browsers")
			continue
		}

		for _, browser := range browsers {
			if browser == "" {
				//log.Println("cant cast browser to string")
				continue
			}

			if ok = strings.Contains(browser, "Android"); ok {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browser := range browsers {
			if !ok {
				//log.Println("cant cast browser to string")
				continue
			}
			if ok = strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])

		email := strings.ReplaceAll(user.Email, "@", " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	_, err = fmt.Fprintln(out, "found users:\n"+foundUsers)
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
	if err != nil {
		return
	}
}

// readFile читает файл построчно и возвращает слайс юзеров
func readFile(filepath string) ([]User, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	var closeErr error
	defer func(file *os.File) {
		closeErr = file.Close()
	}(file)

	var (
		users []User
		user  User
	)

	// по умолчанию разбивает входной поток по символу новой строки \n
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Printf("%v %v\n", err, line)
		err = json.Unmarshal(scanner.Bytes(), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users, closeErr
}
