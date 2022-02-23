package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var nums = map[string]int{
	"zero": 0,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
	"ten": 10,
	"eleven": 11,
}

func NumParse(num string) int {
	if string(num[0]) == "'" {
		num = num[1:]
	}
	if string(num[len(num)-1]) == "." {
		num = num[:len(num)-1]
	}

	num = strings.ToLower(num)

	i, err := strconv.ParseInt(num, 10, 64)

	if err != nil {
		i, ok := nums[num]
		if !ok {
			panic("Wtf is this?? " + num)
		}
		return i
	}
	
	return int(i)
}

func NumToSpace(num int, size int) string {
	str := fmt.Sprint(num)

	if size < len(str) {
		panic("Too big!")
	}

	if size == len(str) + 1 {
		panic("Too small!")
	}
	str += "."
	suffix := ""
	for (len(str) + len(suffix)) != size {	
		str += "0"
	}

	for {
		if str[len(str)-2:] == "00" && str[len(str)-3:len(str)-2] != "." {
			str = str[:len(str)-2]
			str = " " + str
			suffix += " "
		} else {
			break
		}
	}

	return str + suffix
}

func nearestAtX(num float64, multiple float64) int {
	return int(math.Ceil(num/multiple))
}
