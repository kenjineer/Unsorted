import "fmt"

func main() {
	var t int
	fmt.Scanln(&t)

	var result []int
	for i := 0; i < t; i++ {

		var n int
		fmt.Scanln(&n)

		var m int
		fmt.Scanln(&m)

		words := make([][]string, n)
		for i := range words {
			words[i] = make([]string, m)
		}

		for i := 0; i < n; i++ {
			var str string
			fmt.Scanln(&str)

			for idx, r := range str {
				words[i][idx] = string(r)
			}
		}

		var searchWord string
		fmt.Scanln(&searchWord)

		ct := findWordVertical(words, searchWord, n, m)
		ct = ct + findWordHorizontal(words, searchWord, n, m)
		ct = ct + findWordDiagonal(words, searchWord, n, m)
		ct = ct + findWordDiagonalToUp(words, searchWord, n, m)

		result = append(result, ct)
	}

	for i := 0; i < t; i++ {
		str := fmt.Sprintf("Case %v: %v", (i + 1), result[i])
		fmt.Println(str)
	}
}

func findWordHorizontal(words [][]string, searchWord string, n int, m int) int {
	ct := 0
	lenStr := len(searchWord)

	for u := 0; u < n; u++ {
		i := 0
		for i+lenStr <= m {
			str := ""

			for k := 0; k < lenStr; k++ {
				str = str + words[u][i+k]
			}

			if str == searchWord {
				ct++
			}

			if str == swap(searchWord) {
				ct++
			}

			i++
		}
	}
	return ct
}

func findWordVertical(words [][]string, searchWord string, n int, m int) int {
	ct := 0
	lenStr := len(searchWord)

	for u := 0; u < m; u++ {
		i := 0
		for i+lenStr <= n {
			str := ""

			for k := 0; k < lenStr; k++ {
				str = str + words[i+k][u]
			}

			if str == searchWord {
				ct++
			}

			if str == swap(searchWord) {
				ct++
			}

			i++
		}
	}
	return ct
}

func findWordDiagonal(words [][]string, searchWord string, n int, m int) int {
	ct := 0
	lenStr := len(searchWord)

	for u := 0; u < n; u++ {
		for i := 0; i < m; i++ {
			if u+lenStr <= n && i+lenStr <= m {
				str := ""

				for k := 0; k < lenStr; k++ {
					str = str + words[u+k][i+k]
				}

				if str == searchWord {
					ct++
				}

				if str == swap(searchWord) {
					ct++
				}
			}
		}
	}

	return ct
}

func findWordDiagonalToUp(words [][]string, searchWord string, n int, m int) int {
	ct := 0
	lenStr := len(searchWord)

	for u := 0; u < n; u++ {
		for i := 0; i < m; i++ {
			if (u+1)-lenStr >= 0 && i+lenStr <= m {
				str := ""

				// get diagonal word
				for k := 0; k < lenStr; k++ {
					str = str + words[u-k][i+k]
				}

				if str == searchWord {
					ct++
				}

				if str == swap(searchWord) {
					ct++
				}
			}
		}
	}

	return ct
}

func swap(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	} else {
		return swap(split(str)) + string(str[0])
	}
}

func split(str string) string {
	i := 0
	new_str := ""

	for i < (len(str) - 1) {
		new_str = new_str + string(str[i+1])
		i = i + 1
	}

	return new_str
}