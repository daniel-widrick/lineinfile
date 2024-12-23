package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	//Parse CLI ARGS
	if len(os.Args) != 3 {
		//Usage
		fmt.Printf("Usage: %s <filename> '<line>'\n\n",os.Args[0])
		fmt.Printf("Lineinfile is a simple command line utility to ensure\nthat a file contains a given line of text.\n\n")
		fmt.Printf("Examples:\n     lineinfile ~/.bashrc 'alias cdg=\"cd ~/git-repos\"'\n")
		os.Exit(-1)
	}
	filename := os.Args[1]
	line := os.Args[2]

	fmt.Printf("File: %s\n",filename)
	fmt.Printf("Line: %s\n",line)

	//Open File
	file, err := os.OpenFile(filename, os.O_RDWR, 0677)
	if err != nil {
		fmt.Printf("File does not exist: %s\n",filename)
		os.Exit(2)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == line {
			fmt.Println("Line Exists!")
			return
		}
	}
	//Line does not exist
	_, err = file.WriteString(line + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n",err)
		os.Exit(4)
	}


}
