package main

import (
	"fmt"

	lab1 "github.com/g-sht/algorithm-labs/lab1"
)

func main() {
	var problem string
	fmt.Scan(&problem)

	switch problem {
	case "3":
		lab1.ThreeMain()
	case "5":
		lab1.FiveMain()
	case "7":
		lab1.SevenMain()
	case "13":
		lab1.ThirteenMain()
	case "14":
		lab1.FourteenMain()
	case "18":
		lab1.EighteenMain()
	case "22":
		lab1.TwentytwoMain()
	case "23":
		lab1.TwentythreeMain()
	case "27":
		lab1.TwentysevenMain()
	}
}
