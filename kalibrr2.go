package main

import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {

		var n int
		fmt.Scanln(&n)

		var m int
		fmt.Scanln(&m)

		words := make([][]string, n)
		for i := range words {
			words[i] = make([]string, m)
		}

		var army []string
		var armyLocX []int
		var armyLocY []int

		for i := 0; i < n; i++ {
			var str string
			fmt.Scanln(&str)

			for idx, r := range str {
				dt := string(r)
				words[i][idx] = dt

				if "#" != dt && "." != dt {
					army = append(army, dt)
					armyLocX = append(armyLocX, idx)
					armyLocY = append(armyLocY, i)
				}
			}
		}

		fmt.Println(army)
		fmt.Println(armyLocX)
		fmt.Println(armyLocY)
	}
	for i := 0; i < t; i++ {
		str := fmt.Sprintf("Case %v:", (i + 1))
		fmt.Println(str)
	}
}