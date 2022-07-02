package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	executionMilliseconds := 0
	password := "0000"
	possibleNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	for _, n1 := range possibleNumbers {
		for _, n2 := range possibleNumbers {
			for _, n3 := range possibleNumbers {
				for _, n4 := range possibleNumbers {

					time.Sleep(50 * time.Millisecond)
					executionMilliseconds += 50

					kick := fmt.Sprintf("%d%d%d%d", n1, n2, n3, n4)
					fmt.Println(kick)

					if kick == password {
						timePassed := executionMilliseconds / 1000
						fmt.Println("Senha encontrada: " + kick)
						fmt.Printf("Tempo corrido: %ds\n", timePassed)
						os.Exit(2)
					}
				}
			}
		}
	}
}