package gopointer

import (
	"fmt"
	"reflect"
)

type TestKeys struct {
	First  string
	Second string
	Thrid  string
}

func PlusPointer(data *int) int {
	return *data + 100
}

func CheckKeys(keys []string, mapB map[string]string) bool {
	var keyExists []string
	var count int = 0
	for j, _ := range mapB {
		for k := 0; k < len(keys); k++ {
			if keys[k] == j {
				keyExists = append(keyExists, j)
				count += 1
			}
		}
	}
	if len(keys) == count {
		return true
	}
	return false
}

func TestGetKeysFormStruct() bool {
	a := &TestKeys{
		First:  "afoo",
		Second: "asdasdsa",
		Thrid:  "asdsads",
	}
	v := reflect.Indirect(reflect.ValueOf(a))
	// val := reflect.Indirect(reflect.ValueOf(a))
	fmt.Println("v :", v)
	// fmt.Println("val :", val)

	mapA := map[string]string{
		"first":  a.First,
		"second": a.Second,
		"thrind": a.Thrid,
	}

	mapB := map[string]string{
		"four":   a.First,
		"five":   a.Second,
		"thrind": a.Thrid,
	}

	var keys []string
	for i, _ := range mapA {
		// fmt.Printf("%s: %s\n", i, v)
		keys = append(keys, i)
	}

	checkKeys := CheckKeys(keys, mapB)
	return checkKeys
}

func FindPrimeNumber(n int) []int {
	var primes []int
	// modula := 2
	// primes = append(primes, modula)
	for i := 2; i <= n; i++ {
		p := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				p = false
			}
		}
		if p == true {
			primes = append(primes, i)
		}
	}
	return primes
}

func SortNumberAsc(data []int) []int {
	// fmt.Println("data ==", data)
	i := 0
	// fmt.Println("init i ==", i)
	for i < len(data)-1 {
		// fmt.Println("i ->", i)
		// fmt.Println("data[i] ->", data[i])
		if data[i] > data[i+1] {
			temp := data[i]
			data[i] = data[i+1]
			data[i+1] = temp
			i = -1
			// fmt.Println("after each loop value i ->", i)
		} else {
			// fmt.Println("after each loop value i ->", i)
		}
		i += 1
	}
	return data
}

func SortNumberDesc(data []int) []int {
	// fmt.Println("data ->", data)
	i := 0
	for i < len(data)-1 {
		// fmt.Println("i ->", i)
		if data[i] < data[i+1] {
			// fmt.Println("in if i ->", i)
			temp := data[i]
			data[i] = data[i+1]
			data[i+1] = temp
			i = -1
			// fmt.Println("after loop if i ->", i)
		}
		i += 1
	}
	return data
}

func Test() string {
	i := 0
	n := 5
	for i < n-1 {
		// fmt.Println("i ->", i)
		if i+1 > n-1 {
			// fmt.Println("i in if->", i)
			i = -1
			// break
		}
		// fmt.Println("i in else ->", i)
		i += 1
	}
	return "ok"
}

func ReversArray(ary []int) []int {
	// fmt.Println("ary :", ary)
	var revers []int
	for n := len(ary) - 1; n >= 0; n-- {
		revers = append(revers, ary[n])
	}
	return revers
}

// func ShiftLeftArray(data []int) []int {
// 	n := len(data)
// 	i := 0
// 	count := 0
// 	var temp int
// 	for i <= n-1 {
// 		temp = data[i]
// 		if i == n-1 {
// 			data[i] = temp
// 			if count == 1 {
// 				break
// 			} else {
// 				count += 1
// 			}
// 			i = -1
// 		} else {
// 			data[i] = data[i+1]
// 		}
// 		i += 1
// 	}
// 	return data
// }
