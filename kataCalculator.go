package kataCalculator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Функция для преобразования римских чисел в арабские
func RomanToArabic(Roman string) (int, error) {
	RomanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	Arabic := 0
	PrevValue := 0
	for i := len(Roman) - 1; i >= 0; i-- {
		value := RomanNumerals[rune(Roman[i])]
		if value < PrevValue {
			Arabic -= value
		} else {
			Arabic += value
		}
		PrevValue = value
	}

	return Arabic, nil
}

// Функция для преобразования арабских чисел в римские
func ArabicToRoman(Num int) string {
	Val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	Syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	Roman := ""

	for i := 0; i < len(Val); i++ {
		for Num >= Val[i] {
			Num -= Val[i]
			Roman += Syb[i]
		}
	}
	return Roman
}

// Функция для обработки римских чисел
func CheckOperationRoman(Data string, Regexp *regexp.Regexp) {
	Operands := Regexp.FindStringSubmatch(Data)
	FirstOperand, err1 := RomanToArabic(Operands[1])
	TwoOperator := Operands[2]
	SecondOperand, err2 := RomanToArabic(Operands[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Неверный формат чисел:", Data)
		return
	}

	var Result int
	switch TwoOperator {
	case "+":
		Result = FirstOperand + SecondOperand
	case "-":
		Result = FirstOperand - SecondOperand
	case "*":
		Result = FirstOperand * SecondOperand
	case "/":
		if SecondOperand == 0 {
			fmt.Println("Деление на ноль невозможно")
			return
		}
		Result = FirstOperand / SecondOperand
	}
	if Result < 1 {
		panic("The developer is panicking")
	}
	fmt.Println(ArabicToRoman(Result))
}

// Функция для обработки арабских чисел
func CheckOperation(Data string, Regexp *regexp.Regexp) {
	Operands := Regexp.FindStringSubmatch(Data)

	FirstOperandStr := Operands[1]
	SecondOperandStr := Operands[3]

	// Проверяем, что оба операнда целые числа
	if _, err := strconv.ParseFloat(FirstOperandStr, 64); err != nil {
		panic(fmt.Sprintf("Первый операнд не является числом: %s", FirstOperandStr))
	}
	if _, err := strconv.ParseFloat(SecondOperandStr, 64); err != nil {
		panic(fmt.Sprintf("Второй операнд не является числом: %s", SecondOperandStr))
	}

	FirstOperand, err1 := strconv.Atoi(FirstOperandStr)
	SecondOperand, err2 := strconv.Atoi(SecondOperandStr)

	if err1 != nil || err2 != nil {
		panic("The developer is panicking")
	}

	if FirstOperand > 10 || FirstOperand < 1 || SecondOperand > 10 || SecondOperand < 1 {
		panic("The developer is panicking")
	}

	TwoOperator := Operands[2]

	switch TwoOperator {
	case "+":
		fmt.Println(FirstOperand + SecondOperand)
	case "-":
		fmt.Println(FirstOperand - SecondOperand)
	case "*":
		fmt.Println(FirstOperand * SecondOperand)
	case "/":
		fmt.Println(FirstOperand / SecondOperand)
	}
}

func startCalculator() {

	// Регулярное выражение для римских чисел и арифметических операций
	var regexRomanAddition = regexp.MustCompile(`^\s*([IVXLCDM]+)\s*([+\-*\/])\s*([IVXLCDM]+)\s*$`)

	// Регулярные выражения для операций
	var regexAddition = regexp.MustCompile(`^\s*(\d+)\s*([+\-*\/])\s*(\d+)\s*$`)

	fmt.Println(regexRomanAddition)
	fmt.Println(regexAddition)
	for {
		var data string

		fmt.Println("Введите числовые данные (или 'exit' для выхода):")
		fmt.Fscan(os.Stdin, &data)

		if data == "exit" {
			break // завершаем цикл, если введено 'exit'
		}

		switch {
		case regexRomanAddition.MatchString(data): // Проверка строки на соответствие с каждым регулярным выражением
			CheckOperationRoman(data, regexRomanAddition)
		case regexAddition.MatchString(data):
			CheckOperation(data, regexAddition)
		default:
			panic("The developer is panicking")
		}
	}
}
