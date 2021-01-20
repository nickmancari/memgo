package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

const (
	check string = "\u2611"
	uncheck string = "\u2610"
)

func readFile( filename string ) []string {
	file, err := os.Open( filename )
	if err != nil {
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner( bufio.NewReader( file ))

	var tmp []string
	for scanner.Scan() {
		tmp = append( tmp, scanner.Text())
	}
	
	return tmp
}

func main() {
	fmt.Println(check + " Done")
	fmt.Println(uncheck)
}
