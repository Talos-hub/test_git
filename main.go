package main

import (
	"fmt"
	"strconv"
)

func main() {
	Start()
}

func Start() {
	var input string // input users

	for {
		fmt.Println("Вас приветствует  программа конвектор температуры")
		fmt.Println("Выберите то что хотите хотите конвертировать: 1 = F в С, 2: C в F,3 = Exite")

		fmt.Scan(&input)

		out, _ := strconv.Atoi(input)

		switch out {
		case 1:
			var value int
			fmt.Println("Введите ваше число: ")
			fmt.Scan(&value)

			reseelt := ConvectorToFarengate(value)

			fmt.Println("................................")
			fmt.Println("Ваш результат", reseelt)

		case 2:
			var value int
			fmt.Println("Введите ваше число: ")
			fmt.Scan(&value)

			reseelt := ConvertToCelsia(value)

			fmt.Println("................................")
			fmt.Println("Ваш результат", reseelt)

		case 3:
			break
		default:
			fmt.Println("Нет такой команды")

		}

	}
}

// ConvectorToFarengate
//
// Parameters:
//   - value int
//
// Returns:
//   - int
func ConvectorToFarengate(value int) int {
	//(0 °C × 9/5) + 32 = 32 °F

	Result := (value * 9 / 5) + 32

	return Result
}

// ConvertToCelsia
//
// Parameters:
//   - value int
//
// Returns:
//   - int
func ConvertToCelsia(value int) int {
	//(70 °F − 32) × 5/9 = 21,111 °C

	Result := (value - 32) * 5 / 9
	return Result
}
