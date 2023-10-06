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
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
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

	if len(result) == 3 {

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

		if eraA == nil && eraB != nil || eraA != nil && eraB == nil {
			fmt.Println(errors.New("Используются разные системы счисления"))
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
						if c == value && c > 1 {
							fmt.Println(key)
						} else {
							fmt.Println(errors.New("Отрицательное число"))
							break
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
						if c == value && c > 1 {
							fmt.Println(key)
						} else {
							fmt.Println(errors.New("Отрицательное число"))
							break
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

	} else {
		fmt.Println(errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
	}

}
