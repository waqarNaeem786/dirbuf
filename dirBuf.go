package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//path
	if os.Args[1] != "" && os.Args[2] != "" && os.Args[1] != "-h" {

		file, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			path := scanner.Text()
			// fmt.Println(path)
			//http
			url := os.Args[1]
			resp, err := http.Get(url + path)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(path, ":", resp.Status)

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("{URL}", "path to wordlist")
	}

}
