//package main

//import "fmt"

//func main() {
//var a [5]int

//a[4] = 100
//fmt.Println("get:", a[2])
//fmt.Println("len:", len(a))

//b := [5]int{1, 2, 3, 4, 5}
//fmt.Println(b)

// var twoD [2][3]int
// for i := 0; i < 2; i++ {
// for j := 0; j < 3; j++ {
// twoD[i][j] = i + j
//
//		}
//	}
//
//		fmt.Println("2d: ", twoD)
//	}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is", secretNumber)

	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input.Please enter an integer value")
			continue
		}
		fmt.Println("Your guess is ", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.Please try again")
		} else {
			fmt.Println("Correct,you Legend!")
			break
		}
	}

}
