package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type doc struct {
	location string
	data     string
}

func validateargs(args []string) {
	if len(args) < 4 {
		fmt.Printf("ERROR: Not enough args\n")
		fmt.Printf("Usage: %s FILEPATH \"FILENAME REGEX\" \"SEARCH REGEX\" [NUM LINES AROUND MATCH]\n", args[0])
		fmt.Printf("Go regex: https://regex-golang.appspot.com/assets/html/index.html\n")
		fmt.Println(args)
		os.Exit(1)
	}
}

func main() {
	validateargs(os.Args)
	searchDir := os.Args[1]

	docRegex := os.Args[2]
	dataRegex := os.Args[3]

	numAroundMatch := 0
	if len(os.Args) == 5 {
		numAroundMatch, _ = strconv.Atoi(os.Args[4])
	}

	r, _ := regexp.Compile(docRegex)

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fname := f.Name()
		if r.MatchString(fname) {
			fileList = append(fileList, path)
		}
		return nil
	})

	docs := []doc{}
	for _, filename := range fileList {
		cmd := exec.Command("antiword", filename)
		out, _ := cmd.CombinedOutput()

		d := doc{location: filename, data: string(out)}
		docs = append(docs, d)
	}

	re, _ := regexp.Compile(dataRegex)

	numMatches := 0
	for _, doc := range docs {
		if re.MatchString(doc.data) {
			fmt.Printf("Match found in: %s\n", doc.location)
			lines := strings.Split(doc.data, "\n")
			for lineno, line := range lines {
				if re.MatchString(line) {
					numMatches++
					if numAroundMatch == 0 {
						fmt.Printf("\033[1;34m%d:\t%s\n\n\033[0m", lineno+1, line)
					} else {
						fmt.Println("...")
						for i := numAroundMatch; i > 0; i-- {
							fmt.Printf("%d:\t%s\n", (lineno+1)-i, lines[lineno-i])
						}
						//fmt.Printf("%d:\t%s\n\n", lineno+1, line)
						fmt.Printf("\033[1;34m%d:\t%s\n\033[0m", lineno+1, line)
						for i := 1; i <= numAroundMatch; i++ {
							fmt.Printf("%d:\t%s\n", lineno+1+i, lines[lineno+i])
						}
						fmt.Println("...")
					}

				}
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("Total matches across %d files: %d\n", len(docs), numMatches)

}
