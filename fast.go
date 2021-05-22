package main

import (
	"bufio"
	"fmt"
	"github.com/ferrozmatic/bench/models"
	"io"
	"os"
	"strings"
)

// FastSearch - вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	seenBrowsers := make(map[string]bool)
	foundUsers := ""
	index := 0

	for fileScanner.Scan() {
		user := models.User{}
		if err := user.UnmarshalJSON(fileScanner.Bytes()); err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				seenBrowsers[browser] = true
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				seenBrowsers[browser] = true
			}
		}

		if isAndroid && isMSIE {
			email := strings.ReplaceAll(user.Email, "@", " [at] ")
			foundUsers += fmt.Sprintf("[%d] %s <%s>\n", index, user.Name, email)
		}

		index++
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
