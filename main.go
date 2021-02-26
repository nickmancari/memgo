package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fname := "todo.md"

	//take user input and print it to "todo.md" file
	textin, err := ioutil.ReadFile(fname)
	if err == nil {
		fmt.Println(string(textin))
	}
	
	var userInput []string
	
	scanner := bufio.NewScanner(os.Stdin)
	
	//stop the program and write the input to file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "." {
			break
		}
		userInput = append(userInput, line+"\n")
	}
	
	//append text to file
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()
	
	if err == nil {
		fmt.Println("Saving Your Tasks & Closing")
		for _, eachLine := range userInput {
			io.WriteString(f, eachLine)
		}
	}
	
	//print text 
	textin, err = ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s You're TODO List : \n", fname)
	fmt.Println(string(textin))
}
