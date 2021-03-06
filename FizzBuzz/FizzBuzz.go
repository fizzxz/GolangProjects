package main

import "fmt"

func main() {
	fizzBuzz(15)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(fmt.Sprint(i) + " ")
		}
	}
}
