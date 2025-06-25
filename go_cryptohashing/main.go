package main

import (
	"fmt"

	"github.com/fatih/color"
)

func cryptic(a string) string { //Caesar's Cipher
	b := []byte(a) //converting to mutable string
	for j := range len(a) {
		//first make sure all are in lower case (lower_case encryption)
		if b[j] >= 65 && b[j] <= 90 {
			b[j] += 32
		} else if b[j] < 97 || b[j] > 122 {
			continue
		}
		ch := b[j] - 'a'
		ch++
		ch %= 26
		b[j] = ch + 'a'
	}
	return string(b)
}

func main() {
	//Main fn
	var s string
	fmt.Print("Please Enter String to be Encrypted: ")
	fmt.Scan(&s)
	ans := cryptic(s)
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("For the given input %s the encrypted string is %s", red(s), green(ans))
}
