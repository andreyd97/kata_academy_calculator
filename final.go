package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	//var testA = true
	//var testB = true

	reader := bufio.NewReader(os.Stdin)
	scanner, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	scanner = strings.TrimSpace(scanner)

	result := strings.Split(scanner, " ")

	a, err = strconv.Atoi(result[0])
	if err != nil {
		log.Fatal(err)
	}
	b, err = strconv.Atoi(result[2])
	if err != nil {
		log.Fatal(err)
	}

}
