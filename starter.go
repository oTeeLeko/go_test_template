package starter

import (
	"fmt"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func evenOrOdd(num int) string {
	evenOrOdd := "odd"
	if num%2 == 0 {
		evenOrOdd = "even"
	}
	return fmt.Sprintf("%v is an %v number", num, evenOrOdd)
}