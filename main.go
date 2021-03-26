package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"github.com/fatih/color"
)

func main() {
	fname := "todo.md"

	//open "todo.md" file where the list will be saved & show the content
	textin, err := ioutil.ReadFile(fname)
	if err == nil {
		cyan := color.New(color.FgCyan).PrintFunc()
		cyan(string(textin))
	}
	
	var userInput []string
	
	scanner := bufio.NewScanner(os.Stdin)
	
	//take in user input and and stop taking input when "." is entered on a new line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "." {
			break
		}
		if line == "-" {
			continue
		}
		userInput = append(userInput, line+"\n")
	}
	
	//write the scanned text to the file and add empty check box to the item
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()
	
	if err == nil {
		fmt.Println("Saving Your Tasks & Closing")
		for _, eachLine := range userInput {
			l := eachLine
			index := 0
			io.WriteString(f, "[ ] " + l[index:])
		}
	}
	
	//open "todo.md" again, showing the newly added content
	textin, err = ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	fmt.Println( "Your TODO List : \n")
	green := color.New(color.FgGreen).PrintFunc()
	green(string(textin))
}
