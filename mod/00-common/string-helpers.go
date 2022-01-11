package common

import (
	"fmt"
)

func LabelMessage(message string) {
	var hyphens int = 50
	var mark int = (hyphens - len(message)) / 2
	for i := 0; i < mark; i++ {
		fmt.Print("-")
	}
	fmt.Printf(" %v ", message)
	for i := 0; i < mark; i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}
