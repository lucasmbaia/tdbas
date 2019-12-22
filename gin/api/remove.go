package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var (
		hosts []string
		rem   []string
	)

	f1, err := os.Open("moirai.txt")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open("rem.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		if scanner.Text() != "" {
			hosts = append(hosts, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}


	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		if scanner2.Text() != "" {
			rem = append(rem, strings.Replace(scanner2.Text(), ".bkp.nuvem-interal.local", "", -1))
		}
	}

	if err := scanner2.Err(); err != nil {
		panic(err)
	}

	var exists bool

	for _, a := range rem {
		exists = false

		for _, b := range hosts {
			if strings.TrimSpace(a) == strings.TrimSpace(b) {
				exists = true
				break
			}
		}

		if !exists {
			fmt.Println(a)
		}
	}
}
