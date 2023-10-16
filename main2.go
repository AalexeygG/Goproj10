






/* 1-10 через строки
package main

import "fmt"

func rep(str1, str2 *string) {
	for _, elem1 := range *str1 {
		for _, elem2 := range *str2 {
			if elem1 == elem2 {
				fmt.Printf("%c ", elem1)
			}
		}
	}
}

func main() {
	var str1, str2 string
	_, err := fmt.Scan(&str1, &str2)
	if err == nil {
		rep(&str1, &str2)
	} else {
		fmt.Println("wtf")
	}
}
*/

/*package main

import (
	"fmt"
	"math"
)

func M(p, v float64) float64 {
	return p * v
}
func W(k, p, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}
func T(k, p, v float64) float64 {
	return 6 / W(k, p, v)
}

func main() {
	var k, p, v float64
	_, err := fmt.Scan(&k, &p, &v)
	if err == nil {
		fmt.Println(T(k, p, v))
	} else {
		fmt.Println("wtf")
	}
}
*/

/*package main

import (
	"fmt"
	"strconv"
)

func quadr(str *string) {
	for _, elem := range *str {
		val, _ := strconv.Atoi(string(elem)) // ***
		fmt.Print(val * val)
	}
}

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err == nil {
		quadr(&str)
	} else {
		fmt.Println("mistake")
	}
}
*/
/*
либо вот так:

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		c -= '0'
		fmt.Print(c * c)
	}
}
*/

/*
***
	arrStr := strings.Split(*str, "")
	for _, elem := range arrStr {
		val, _ := strconv.Atoi(elem)
		fmt.Print(val * val)
	}
***
*/

/*package main

import (
	"fmt"
)

func findMin(str *string) rune {
	var min rune = 57
	for _, elem := range *str {
		if elem < min {
			min = elem
		}
	}
	return min
}

func findMax(str *string) rune {
	var max rune = 48
	for _, elem := range *str {
		if elem > max {
			max = elem
		}
	}
	return max
}

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err == nil {
		fmt.Printf("%c", findMax(&str))
	}
}
*/

/*package main

import (
	"fmt"
)

func add(str *string) {
	runeStr := []rune(*str)
	for i := 0; i < len(runeStr)-1; i++ {
		fmt.Printf("%c*", runeStr[i])
	}
	fmt.Printf("%c", runeStr[len(runeStr)-1])
}

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err == nil {
		add(&str)
	} else {
		fmt.Println("wtf")
	}
}
*/

/*package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("ошибка")
		return -1, err
	}
	return a / b, nil
}

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err == nil {
		ans, err := divide(a, b)
		if err != nil {
			fmt.Println(err) //можно просто "ошибка", без errors.New
		} else {
			fmt.Println(ans)
		}
	} else {
		fmt.Println("ошибка")
	}
}

*/

/*package main

import (
	"fmt"
	"unicode"
)

func checkPass(str *string) string {
	runeStr := []rune(*str)
	if len(runeStr) >= 5 {
		for i := range runeStr {
			if unicode.Is(unicode.Latin, runeStr[i]) || unicode.Is(unicode.ASCII_Hex_Digit, runeStr[i]) { //unicode.Is(unicode.Digit, runeStr[i] можно так
				continue
			}
			return "Wrong password"
		}
		return "Ok"
	}
	return "Wrong password"
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkPass(&str))
}

*/

/*package main

import (
	"fmt"
	"strings"
)

func rebulid(str *string) {
	for _, elem := range *str {
		if strings.Count(*str, string(elem)) == 1 {
			fmt.Print(string(elem))
		}
	}
}

func main() {
	var stroka string
	fmt.Scan(&stroka)
	rebulid(&stroka)
}
*/

/*
package main

import (
	"fmt"
)

func rebuild(str *string) {
	runeStr := []rune(*str)
	ansStr := []rune("")
	for i := 1; i <= len(runeStr)-1; i += 2 {
		ansStr = append(ansStr, runeStr[i])
	}
	fmt.Println(string(ansStr))
}

func main() {
	var stroka string
	fmt.Scan(&stroka)
	rebuild(&stroka)

}
*/

/*
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	stroka, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runStr := []rune(stroka)
	if unicode.IsUpper(runStr[0]) && runStr[len(runStr)-1] == '.' { //можно rune('.')
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}


func check(s *string) bool {
	runeStr := []rune(*s)
	for i := 0; i < len(runeStr)/2; i++ {
		if runeStr[i] != runeStr[len(runeStr)-i-1] {
			return false
		}
	}
	return true
}
*/

/*
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func firstLetterCheck(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

func checkPointAvailability(s string) bool {
	return strings.HasSuffix(s, ".")
}

func checkAllRules(s string) bool {
	return firstLetterCheck(s) && checkPointAvailability(s)
}

func boolToStringAnswer(b bool) string {
	if b {
		return "Right"
	}
	return "Wrong"
}

func mainFunction(buf io.Reader) (string, error) {
	s, err := bufio.NewReader(buf).ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return boolToStringAnswer(
		checkAllRules(
			s,
		),
	), nil
}

func main() {
	s, err := mainFunction(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Print(s)
}
*/
