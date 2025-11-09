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

	lines, err := readFile(filePath)
	if err != nil {
		panic(err)
	}

	var (
		seenBrowsers   []string
		uniqueBrowsers int
		foundUsers     string
	)

	var users []User
	for _, line := range lines {
		var user User
		// fmt.Printf("%v %v\n", err, line)
		err = json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers := user.Browsers

		var browser string
		for _, browser = range browsers {

			if ok := strings.Contains(browser, "Android"); ok {
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

		for _, browser = range browsers {

			if ok := strings.Contains(browser, "MSIE"); ok {
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

// readFile читает файл построчно и возвращает слайс строк
func readFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	var closeErr error
	defer func(file *os.File) {
		closeErr = file.Close()
	}(file)

	var (
		lines []string
	)

	// по умолчанию разбивает входной поток по символу новой строки \n
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
	}

	return lines, closeErr
}
