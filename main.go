package katacalcular

import (
	"fmt"
	"regexp"
	"strconv"
)

// Функция для преобразования римских чисел в арабские
func RomanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	arabic := 0
	prevValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			arabic -= value
		} else {
			arabic += value
		}
		prevValue = value
	}

	return arabic, nil
}

// Функция для преобразования арабских чисел в римские
func ArabicToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syb[i]
		}
	}
	return roman
}

// Функция для обработки римских чисел
func CheckOperationRoman(data string, regexp *regexp.Regexp) {
	operands := regexp.FindStringSubmatch(data)
	firstOperand, err1 := RomanToArabic(operands[1])
	twoOperator := operands[2]
	secondOperand, err2 := RomanToArabic(operands[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Неверный формат чисел:", data)
		return
	}

	var result int
	switch twoOperator {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		if secondOperand == 0 {
			fmt.Println("Деление на ноль невозможно")
			return
		}
		result = firstOperand / secondOperand
	}
	if result < 1 {
		panic("The developer is panicking")
	}
	fmt.Println(ArabicToRoman(result))
}

// Функция для обработки арабских чисел
func CheckOperation(data string, regexp *regexp.Regexp) {
	operands := regexp.FindStringSubmatch(data)

	firstOperandStr := operands[1]
	secondOperandStr := operands[3]

	// Проверяем, что оба операнда целые числа
	if _, err := strconv.ParseFloat(firstOperandStr, 64); err != nil {
		panic(fmt.Sprintf("Первый операнд не является числом: %s", firstOperandStr))
	}
	if _, err := strconv.ParseFloat(secondOperandStr, 64); err != nil {
		panic(fmt.Sprintf("Второй операнд не является числом: %s", secondOperandStr))
	}

	firstOperand, err1 := strconv.Atoi(firstOperandStr)
	secondOperand, err2 := strconv.Atoi(secondOperandStr)

	if err1 != nil || err2 != nil {
		panic("The developer is panicking")
	}

	if firstOperand > 10 || firstOperand < 1 || secondOperand > 10 || secondOperand < 1 {
		panic("The developer is panicking")
	}

	twoOperator := operands[2]

	switch twoOperator {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		fmt.Println(firstOperand / secondOperand)
	}
}

// func main() {

// // Регулярное выражение для римских чисел и арифметических операций
// var regexRomanAddition = regexp.MustCompile(`^\s*([IVXLCDM]+)\s*([+\-*\/])\s*([IVXLCDM]+)\s*$`)

// // Регулярные выражения для операций
// var regexAddition = regexp.MustCompile(`^\s*(\d+)\s*([+\-*\/])\s*(\d+)\s*$`)

// fmt.Println(regexRomanAddition)
// fmt.Println(regexAddition)
// for {
// 	var data string

// 	fmt.Println("Введите числовые данные (или 'exit' для выхода):")
// 	fmt.Fscan(os.Stdin, &data)

// 	if data == "exit" {
// 		break // завершаем цикл, если введено 'exit'
// 	}

// 	switch {
// 	case regexRomanAddition.MatchString(data): // Проверка строки на соответствие с каждым регулярным выражением
// 		CheckOperationRoman(data, regexRomanAddition)
// 	case regexAddition.MatchString(data):
// 		CheckOperation(data, regexAddition)
// 	default:
// 		panic("The developer is panicking")
// 	}
// }
// }
