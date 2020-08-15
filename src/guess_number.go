package main

import (
	"fmt"
	"math/rand"
	"math"
)


func getRandomNumber(r int) int {
	return rand.Intn(r)
}


func isCloseAnswer(answer int, number int) bool {
	thresold := 5
	diff := int64(math.Abs(float64(answer - number)))
	if (diff <= int64(thresold)){
		return true
	}
	return false
}


func main(){
	MAX_ATTEMPTS := 5
	randomNumber := getRandomNumber(100)
	isAnswered := false
	for attempts := 0; attempts < MAX_ATTEMPTS; attempts++ {
		var guess int
		fmt.Println("Your guess")
		fmt.Scanln(&guess)
		if guess == randomNumber {
			fmt.Println("Correct Answer!")
			isAnswered = true
			break
		}
		if(isCloseAnswer(guess, randomNumber)){
			fmt.Println("You are very close to the answer")
		}
		fmt.Println("Incorrect guess. Continue")
	}
	if(!isAnswered){
		fmt.Println("All the chances are gone now.")
	}
	fmt.Println("Game Ends")
}
