package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func slicetostring(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func main() {
	var num int = 5647
	var sum int = 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10

	}
	var total int = 0
	if sum > 9 {
		for sum > 0 {
			total = total + sum%10
			sum = sum / 10
		}
		fmt.Println(total)
	} else {
		fmt.Println(sum)
	}
	var value []int
	number := 3456
	for number > 0 {
		value = append(value, number%10)
		number = number / 10
	}
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	fmt.Println(value)
	var a int = 999
	for (a / 10) > 0 {
		a = a/10 + a%10
	}
	fmt.Println(a)

	b := 1234
	fmt.Println(string(b))
	var input = 20
	output := strconv.Itoa(input)
	fmt.Println(output)
	fmt.Println("input Type:", reflect.TypeOf(input))
	fmt.Println("output Type", reflect.TypeOf(output))
	fmt.Println("ToString", slicetostring(value, ""))
}
