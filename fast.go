package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func FastSearch(out io.Writer) {

	lines, err := readFile(filePath)
	//file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	//fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var (
		seenBrowsers   []string
		uniqueBrowsers int
		foundUsers     string
	)

	//lines := strings.Split(string(fileContents), "\n")

	users := make([]map[string]interface{}, 0)
	for _, line := range lines {
		user := make(map[string]interface{})
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

		browsers, ok := user["browsers"].([]interface{})
		if !ok {
			log.Println("cant cast browsers")
			continue
		}

		var browser string
		for _, browserRaw := range browsers {
			browser, ok = browserRaw.(string)
			if !ok {
				log.Println("cant cast browser to string")
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

		for _, browserRaw := range browsers {
			browser, ok = browserRaw.(string)
			if !ok {
				log.Println("cant cast browser to string")
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

		email := strings.ReplaceAll(user["email"].(string), "@", " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

// readFile читает файл построчно и возвращает слайс строк
func readFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		lines []string
	)

	// по умолчанию разбивает входной поток по символу новой строки \n
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 2048)
	scanner.Buffer(buf, 4096)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
	}

	return lines, nil
}
