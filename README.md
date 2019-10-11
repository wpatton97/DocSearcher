# DocSearcher
A program to search multiple a directory nested with nested doc files, and return any matching text to the result string

## Usage
`./DocSearcher FILEPATH "FILENAME REGEX" "SEARCH REGEX" [NUM LINES AROUND MATCH]`
 * `FILENAME REGEX` - Refers to the regex that will tell the program which doc files to search
 * `SEARCH REGEX` - Refers to the regex that will be searched through each file
 * `NUM LINES AROUND MATCH`[OPTIONAL] - Will print N extra lines before and after each match

## Requirements
'antiword' must be installed to your local machine

## Installing 
```
brew install antiword # if you are on mac/missing antiword
sudo apt install antiword # if you are on debian/missing antiword
sudo pacman -Syu antiword # if you are on arch missing antiword

go get github.com/Ragex04/DocSearcher
go install github.com/Ragex04/DocSearcher
```

Most this code depends on Antiword, credit to antiword goes to __Adri van Os(http://www.winfield.demon.nl/)__
