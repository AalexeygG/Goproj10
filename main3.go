// 3.6.2 мое решение
// проверка для гита
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type myStruct []struct {
	GlobalId int64 `json:"global_id"`
}

func main() {
	file, err := os.Open("data-20190514T0100.json")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var ans myStruct
	if err := json.Unmarshal(data, &ans); err != nil {
		panic(err)
	}
	var sum int64
	for i := range ans {
		sum += ans[i].GlobalId
	}
	fmt.Println(sum)
}

/* 3.6.2 - решение с декодом и енкодом
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var items []struct {
		ID uint64 `json:"global_id"`
	}
	f, _ := os.Open("data-20190514T0100.json")
	json.NewDecoder(f).Decode(&items)
	var sum uint64
	for _, item := range items {
		sum += item.ID
	}
	fmt.Println(sum)
}
*/

/*
3.6.1
package main

import (
	"encoding/json"
	"io"
	"os"
)

type studentStruct struct {
	Students []struct {
		Rating []float64
	}
}

type assign struct {
	Average float64
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var newData studentStruct
	if err := json.Unmarshal(data, &newData); err != nil {
		panic(err)
	}
	var number int
	for i := range newData.Students {
		number += len(newData.Students[i].Rating)
	}
	var avr assign
	avr.Average = float64(number) / float64(len(newData.Students))
	ansData, err := json.MarshalIndent(avr, "", "    ")
	io.WriteString(os.Stdout, string(ansData))
}
*/

/* 3.5.3
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ans := bufio.NewReader(file)
	var count int
	for {
		line, err := ans.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("something went wrong")
				break
			}
		}
		if line == "0;" {
			fmt.Println(count + 1)
			break
		}
		count++
	}
}
*/

/* 3.5.2
через открытие zip файла
package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	z, _ := zip.OpenReader("task.zip")
	defer z.Close()
	for _, f := range z.File {
		r, _ := f.Open()
		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(f.Name, rows[4][2])
		}
		r.Close()
	}
}
*/

/* 3.5.2 распаковал архив
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func myWalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	openedFile := csv.NewReader(file)
	data, _ := openedFile.ReadAll()
	if len(data) > 1 {
		fmt.Println(data[4][2])
	}
	return nil
}

func main() {
	root := "C:/Users/yapod/Downloads/task"

	if err := filepath.Walk(root, myWalkFunc); err != nil {
		fmt.Printf("ошибка %v\n", err)
	}
}
*/

/* 3.5.1
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	str := bufio.NewScanner(os.Stdin)
	var sum int
	for str.Scan() {
		forsum, err := strconv.Atoi(str.Text())
		if err != nil {
			panic(err)
		}
		sum += forsum
	}
	n := bufio.NewWriter(os.Stdout) // io.WriteString(os.Stdout, strconv.Itoa(amount)) - можно вместо всего, но надо обработать ошибку
	_, err := n.WriteString(strconv.Itoa(sum))
	if err != nil {
		panic(err)
	}
	n.Flush()
}
*/

/*3.4.2
более менее лаконичное решение
package main

import (
	"fmt"
	"strings"
) // пакет используется для проверки ответа, не удаляйте его

type accumulator struct {
	battery string
}

func (a accumulator) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", strings.Count(a.battery, "1")))
}

func main() {
	var batteryForTest accumulator
	fmt.Scan(&batteryForTest.battery)
}

// batteryForTest - не забывайте об имени
// } Скобка, закрывающая функцию main() вам не видна, но она есть
*/

/*3.4.2
Мое первое решение

package main

import (
	"fmt"
	"strings"
) // пакет используется для проверки ответа, не удаляйте его

type accumulator struct {
	battery string
}

func (a accumulator) String() string {
	var ansStr string = "["
	var count int
	for _, elem := range a.battery {
		if elem == '0' {
			count++
		}
	}
	ansStr += strings.Repeat(" ", count)
	ansStr += strings.Repeat("X", 10-count)
	ansStr += "]"
	return ansStr
}

func main() {
	var batteryForTest accumulator
	fmt.Scan(&batteryForTest.battery)
}
*/

/* 3.4.1
мое решение:

package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func readTask() (interface{}, interface{}, interface{}) {
	var val1, val2, oper interface{}
	_, err := fmt.Scan(&val1, &val2, &oper)
	if err != nil {
		panic(err)
	}
	return val1, val2, oper
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	switch v1 := value1.(type) {            // все полученные значения имеют тип пустого интерфейса
	case float64:
		switch v2 := value2.(type) {
		case float64:
			switch operation {
			case "+":
				fmt.Printf("%.4f", v1+v2)
			case "-":
				fmt.Printf("%.4f", v1-v2)
			case "/":
				fmt.Printf("%.4f", v1*v2)
			case "*":
				fmt.Printf("%.4f", v1*v2)
			default:
				fmt.Println("неизвестная операция")
			}
		default:
			fmt.Printf("value=%v: %T", v2, v2)
		}
	default:
		fmt.Printf("value=%v: %T", v1, v1)
	}
}
*/

/* 3.4.1

прикольная идея с map и функциями:

package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)

func main() {
    vi1, vi2, operation := readTask()
    vi := [2]interface{}{vi1, vi2}
    var v1, v2 float64
    vf := [2]*float64{&v1, &v2}
    var ok bool
    for i, v:= range vi{
        if *vf[i], ok = v.(float64); !ok {
            fmt.Printf("value=%v: %T", v, v)
            return
        }
    }
    ops:= map[string]func() float64{
        "+": func() float64 {return v1+v2 },
        "-": func() float64 {return v1-v2 },
        "*": func() float64 {return v1*v2 },
        "/": func() float64 {return v1/v2 },
    }

    if oper, ok := operation.(string); ok {
        if fun, ok:= ops[oper]; ok{
            fmt.Printf("%.4f", fun())
            return
        }
    }
    fmt.Print("неизвестная операция")

}

*/

/*3.4.1
Лаконичное решение +-
package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)

func CheckFloat(a ...interface{}) bool {
	for _, k := range a {
		if _, ok := k.(float64); !ok { fmt.Printf("value=%v: %T", k, k); return false }
	}
	return true
}

func main() {
	value1, value2, operation := readTask()
	if CheckFloat(value1, value2) {
		switch operation.(string) {
		case "+": fmt.Printf("%.4f", value1.(float64) + value2.(float64))
		case "-": fmt.Printf("%.4f", value1.(float64) - value2.(float64))
		case "*": fmt.Printf("%.4f", value1.(float64) * value2.(float64))
		case "/": fmt.Printf("%.4f", value1.(float64) / value2.(float64))
		default: fmt.Print("неизвестная операция")
		}
	}
}
*/

/* 3.3.1
package main

import (
	"strconv"
)

func main() {
	fn := func(dig uint) uint {
		strDig := strconv.FormatUint(uint64(dig), 10)
		var ansStr string
		for _, elem := range strDig {
			if elem == '0' || elem%2 != 0 {
				continue
			}
			ansStr += string(elem)
		}
		ans, err := strconv.ParseUint(ansStr, 10, 64)
		if err == nil {
			return uint(ans)
		}
		return 100
	}
}

/* 3.3.1
 вариант через арфиметику
fn := func(n uint) uint {
    var i, res uint
	for i = 1; n > 0; n /= 10 {
		c := n % 10
		if c != 0 && c&1 == 0 {
			res += uint(c) * i
			i *= 10
		}
	}
    if res == 0 {res = 100}
	return res
}
*/

/*3.3.5 - вопрос

//func(int, int) в начале выполнилось из-за наличия тела фукнции(тела функции)

package main

import "fmt"

func main() {
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)
}

*/

/*3.2.3
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dev(str *string) {
	*str = strings.Replace(strings.Replace(*str, ",", ".", -1), " ", "", -1)
	arrstr := strings.Split(*str, ";")
	dig1, err1 := strconv.ParseFloat(arrstr[0], 64)
	dig2, err2 := strconv.ParseFloat(arrstr[1], 64)
	if err1 == nil && err2 == nil {
		fmt.Printf("%.4f", dig1/dig2)
	} else {
		fmt.Println("u r wrong as usual)")
	}
}

func main() {
	str, err := bufio.NewReader(os.Stdin).ReadString('\r')
	if err != nil && err != io.EOF {
		fmt.Println("try again")
	} else {
		dev(&str)
	}
}
*/

/*3.2.2
package main

import "strconv"

func adding(str1, str2 string) int64 {
	var newStr1 string = ""
	var newStr2 string = ""
	for _, elem := range str1 {
		if elem >= '0' && elem <= '9' {
			newStr1 += string(elem)
		}
	}
	for _, elem := range str2 {
		if elem >= '0' && elem <= '9' {
			newStr2 += string(elem)
		}
	}
	newInt1, err1 := strconv.ParseInt(newStr1, 10, 64)
	newInt2, err2 := strconv.ParseInt(newStr2, 10, 64)
	if err1 == nil && err2 == nil {
		return newInt1 + newInt2
	}
	return 0
}
*/

/*package main

import "fmt"

func main() {
	a := "12"
	b := []byte(a)
	fmt.Println(b)
	a += string(b) //переводит в строку
	fmt.Println(a)
}
*/

/* 3.2.1

func convert(x int64) uint16 {
	return uint16(x)
}
*/

/* 3.1.2
package main

func main() {
	groupCity := map[int][]string{
		10:   []string{"..."}, // города с населением 10-99 тыс. человек
		100:  []string{"..."}, // города с населением 100-999 тыс. человек
		1000: []string{"..."}, // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"город": 123,
	}
	for key1 := range cityPopulation {
		for key2, val2 := range groupCity {
			if key2 == 100 {
				continue
			}
			for _, j := range val2 {
				if j == key1 {
					delete(cityPopulation, j)
				}
			}
		}
	}
	//норм решение
	//for _, city := range append(groupCity[10], groupCity[1000]...) {
    //	delete(cityPopulation, city)
}
*/

/* 3.1.1
package main

import "fmt"

func work(x int) int

func main() {
	m := make(map[int]int)
	var number int
	for i := 0; i < 10; i++ {
		fmt.Scan(&number)
		_, inMap := m[number]
		if !inMap {
			m[number] = work(number)
		}
		fmt.Print(m[number])
	}
}
*/
