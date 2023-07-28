package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Запрос выражения у пользователя
	fmt.Println("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cleanInput := strings.TrimSpace(strings.ToUpper(strings.TrimRight(input, "\n")))

	// Определение оператора в выражении для Split и Switch
	var operator string
	operators := [...]string{"+", "-", "*", "/"}
	for _, op := range operators {
		if strings.Contains(cleanInput, op) {
			operator = op
			break
		} else if op == operators[len(operators)-1] {
			fmt.Println("Оператор не найден, строка не является математической операцией")
			return
		}
	}

	// Сепарация ввода на операнды по оператору
	operands := strings.Split(cleanInput, operator)
	if len(operands) > 2 {
		fmt.Println("Неверный формат выражения, не более двух операндов в выражении")
		return
	}
	if len(operands) < 2 {
		fmt.Println("Cтрока не является математической операцией")
		return
	}

	var OperandSystem1, OperandSystem2 string

	if matched, _ := regexp.MatchString("^(I{1,3}|IV|V|VI{1,3}|IX|X)$", operands[0]); matched {
		OperandSystem1 = "roman"
		fmt.Println("OperandSystem1 = roman")
	} else if matched, _ := regexp.MatchString("^(10|[1-9])$", operands[0]); matched {
		OperandSystem1 = "arabic"
		fmt.Println("OperandSystem1 = arabic")
	} else {
		OperandSystem1 = "wrong"
		fmt.Println("Калькулятор принимает значения от 1 до 10, ошибка в первом операнде")
	}

	if matched, _ := regexp.MatchString("^(I{1,3}|IV|V|VI{1,3}|IX|X)$", operands[1]); matched {
		OperandSystem2 = "roman"
		fmt.Println("OperandSystem2 = roman")
	} else if matched, _ := regexp.MatchString("^(10|[1-9])$", operands[1]); matched {
		OperandSystem2 = "arabic"
		fmt.Println("OperandSystem2 = arabic")
	} else {
		fmt.Println("Калькулятор принимает значения от 1 до 10, ошибка в втором операнде")
	}

	var result int
	var num1, num2 int
	if OperandSystem1 == "roman" && OperandSystem2 == "roman" {
		fmt.Println("Римское счисление.")
		num1 = Roman2Arabic(operands[0])
		num2 = Roman2Arabic(operands[1])
		fmt.Println(num1, num2)
		switch operator {
		case "+":
			result = num1 + num2

		case "-":
			result = num1 - num2

		case "*":
			result = num1 * num2

		case "/":
			if num2 == 0 {
				fmt.Println("Мы не в техническом ВУЗе")
			}
			result = num1 / num2
		default:
			fmt.Println("Неверный оператор.")
		}
		if result < 0 {
			fmt.Println("В римской системе нет отрицательных чисел")
		} else {
			fmt.Println(Arabic2Roman(result))
		}
	} else if OperandSystem1 == "arabic" && OperandSystem2 == "arabic" {
		fmt.Println("Арабское счисление.")
		_, err := fmt.Sscanf(operands[0], "%d", &num1)
		if err != nil {
			fmt.Println("Неверный формат первого операнда.")
		}

		_, err = fmt.Sscanf(operands[1], "%d", &num2)
		if err != nil {
			fmt.Println("Неверный формат второго операнда.")
		}

		// Определение математической операции

		switch operator {
		case "+":
			result = num1 + num2

		case "-":
			result = num1 - num2

		case "*":
			result = num1 * num2

		case "/":
			if num2 == 0 {
				fmt.Println("Мы не в техническом ВУЗе")
			}
			result = num1 / num2
		default:
			fmt.Println("Неверный оператор.")
		}
		fmt.Println(result)
	} else if OperandSystem1 == "roman" && OperandSystem2 == "arabic" || OperandSystem2 == "roman" && OperandSystem1 == "arabic" {
		fmt.Println("Введите только римские или арабские числа.")
	}
}

// Вычисляем арабское значение римскоо числа
func Roman2Arabic(num string) int {
	result := 0
	DicR2A := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}

	for i := 0; i < len(num)-1; i++ {
		if DicR2A[string(num[i])] < DicR2A[string(num[i+1])] {
			result -= DicR2A[string(num[i])]
		} else {
			result += DicR2A[string(num[i])]
		}
	}
	result += DicR2A[string(num[len(num)-1])]
	return result
}

// Вычисляем римское значение арабского числа
func Arabic2Roman(num int) string {
	DicA2R := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}
	var result string
	var keys []int

	for k := range DicA2R {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, value := range keys {
		for num >= value {
			StringNumeral := DicA2R[value]
			result += StringNumeral
			num -= value
		}
	}
	return result
}
