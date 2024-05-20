package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// romanToArabic представляет карту для преобразования римских цифр в арабские
var romanToArabic = map[string]int{
    "I":  1,
    "II": 2,
    "III": 3,
    "IV":  4,
    "V":   5,
    "VI":  6,
    "VII": 7,
    "VIII": 8,
    "IX":   9,
    "X":   10,
}

// arabicToRoman представляет массив для преобразования арабских цифр в римские
const maxRomanResult = "C"

var arabicToRoman = []string{
    "", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
    "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
    "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
    "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
    "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
    "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
    "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
    "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
    "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
    "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
}

// isArabic проверяет, является ли переданная строка арабской цифрой
func isArabic(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

// isRoman проверяет, является ли переданная строка римской цифрой
func isRoman(s string) bool {
    _, exists := romanToArabic[s]
    return exists
}

// toArabic преобразует переданную строку в арабское число
func toArabic(s string) int {
    if val, err := strconv.Atoi(s); err == nil {
        return val
    }
    return romanToArabic[s]
}

// toRoman преобразует переданное арабское число в римскую цифру
func toRoman(n int) string {
    if n <= 0 || n >= len(arabicToRoman) {
        panic("Результат выходит за пределы диапазона")
    }
    return arabicToRoman[n]
}

// evaluate выполняет арифметическую операцию над двумя числами
func operations(a, b int, operator string) int {
    switch operator {
    case "-":
        return a - b
    case "+":
        return a + b
    case "/":
        return a / b
    case "*":
        return a * b
    default:
        panic("Некорректное действие")
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Калькулятор")
    fmt.Println("Введите число, напрмер: 1 + 2, либо I + II, используя данные действия: -, +, /, * или exit для выхода") 
    
    for {
        if !scanner.Scan() {
            break
        }
        expression := scanner.Text()
        
        // Выход из программы при вводе "exit"
        if strings.ToLower(expression) == "exit" {
            break
        }

        tokens := strings.Fields(expression)
        if len(tokens) != 3 {
            panic("Неверный формат ввода")
        }

        operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]

       
        // Обработка римских цифр
        if isRoman(operand1) && isRoman(operand2) {
            a := toArabic(operand1)
            b := toArabic(operand2)
            if a < 1 || a > 10 || b < 1 || b > 10 {
                panic("Римские цифры вне диапазона от I до X")
            }
            result := operations(a, b, operator)
            if result <= 0 {
                panic("Результат выходит за пределы диапазона")
            }
            fmt.Printf("Результат: %s\n", toRoman(result))
        }

         // Обработка арабских цифр
         if isArabic(operand1) && isArabic(operand2) {
            a := toArabic(operand1)
            b := toArabic(operand2)
            if a < 1 || a > 10 || b < 1 || b > 10 {
                panic("Арабские цифры вне диапазона от 1 до 10")
            }
            result := operations(a, b, operator)
            fmt.Printf("Результат: %d\n", result)
        }

        // Обработка некорректных входных данных
        if !isArabic(operand1) && !isRoman(operand1) || !isArabic(operand2) && !isRoman(operand2) {
            panic("Некорректные входные данные")
        }
    }

    if scanner.Err() != nil {
        fmt.Println("Ошибка при чтении входных данных:", scanner.Err())
    }
}