package main

import (

	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {

	fmt.Println("What Do You Need To Do? ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("[ ] ", scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

