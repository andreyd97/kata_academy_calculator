package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var riman = map[string]int{
	"I":     1,
	"II":    2,
	"III":   3,
	"IV":    4,
	"V":     5,
	"VI":    6,
	"VII":   7,
	"VIII":  8,
	"IX":    9,
	"X":     10,
	"XI":    11,
	"XII":   12,
	"XIII":  13,
	"XIV":   14,
	"XV":    15,
	"XVI":   16,
	"XVII":  17,
	"XVIII": 18,
	"XIX":   19,
	"XX":    20,
}

var c int

func sumAB(a, b int) int {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		c = a + b
	} else {
		fmt.Println(errors.New("Числа должны быть с пределах от 1 до 10 включительно"))
	}
	return c
}

func subAB(a, b int) int {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		c = a - b

	} else {
		fmt.Println(errors.New("Введенные числа не подходят по условиям задачи"))
	}
	return c
}
func multiAB(a, b int) int {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		c = a * b

	} else {
		fmt.Println(errors.New("Введенные числа не подходят по условиям задачи"))
	}
	return c
}

func divAB(a, b int) int {
	if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
		c = a / b
	} else {
		fmt.Println(errors.New("Введенные числа не подходят по условиям задачи"))
	}
	return c

}

func main() {
	var romA, romB int
	var testA = true
	var testB = true

	reader := bufio.NewReader(os.Stdin)
	scanner, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(errors.New("Данные передаются в одну строку!!!"))
	}

	scanner = strings.TrimSpace(scanner)

	result := strings.Split(scanner, " ")

	a, eraA := strconv.Atoi(result[0])
	if eraA != nil {
		testA = false
		_ = a
	}

	b, eraB := strconv.Atoi(result[2])
	if eraB != nil {
		testB = false
		_ = b
	}

	if testA == false && testB == false {
		for key, value := range riman {
			if result[0] == key {
				romA = value
			}
			if result[2] == key {
				romB = value
			}
		}
		switch result[1] {
		case "+":
			{
				c = sumAB(romA, romB)
				for key, value := range riman {
					if c == value {
						fmt.Println(key)
					}
				}
			}
		case "-":
			{
				c = subAB(romA, romB)
				for key, value := range riman {
					if c == value {
						fmt.Println(key)
					}
				}
			}
		case "*":
			{
				c = multiAB(romA, romB)
				for key, value := range riman {
					if c == value {
						fmt.Println(key)
					}
				}
			}
		case "/":
			{
				c = divAB(romA, romB)
				for key, value := range riman {
					if c == value {
						fmt.Println(key)
					}
				}
			}
		}
	}

	if eraA == nil && eraB == nil {
		switch result[1] {
		case "+":
			{
				c = sumAB(a, b)
				fmt.Println(c)
			}
		case "-":
			{
				c = subAB(a, b)
				fmt.Println(c)
			}
		case "*":
			{
				c = multiAB(a, b)
				fmt.Println(c)
			}
		case "/":
			{
				c = divAB(a, b)
				fmt.Println(c)
			}
		}
	}

}
