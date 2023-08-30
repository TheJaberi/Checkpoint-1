package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}
	num,err := Atoi(os.Args[1])
	if err == "error" {
		fmt.Printf("ERROR: cannot convert to roman digit\n")
		return
	}
	Rom := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	res1 := ""
	res2 := ""
	for num != 0 {
		remove := ToRemove(num)
		num -= remove
		if len(Rom[remove]) == 1 {
			res1 += Rom[remove] + "+"
		} else {
			res1 += "("+string(Rom[remove][1]) + "-" + string(Rom[remove][0])+ ")" + "+"
		}
		res2 += Rom[remove]
	}
	fmt.Printf("%v\n",res1[:len(res1)-1])
	fmt.Printf("%v\n",res2) 
}

func ToRemove(num int) int {
	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i< len(keys); i++ {
		if keys[i] <= num {
			return keys[i]
		}
	}
	return 0
}

func Atoi(s string) (int, string) {
	num := 0
	for _, r := range s {
		num *= 10
		num += int(r-'0')
	}
	if num == 0 || num >= 4000 {
		return -1 , "error"
	}
	return num , ""
}

