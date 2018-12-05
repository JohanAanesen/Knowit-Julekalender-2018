package main

import (
	"fmt"
)

var numbers = []string{"2", "3", "4", "5", "6", "7", "8", "7", "6", "5", "4", "3", "2", "1"}
var signs = [3]string{"+", "-", ""}

var sequences []string
var exactNums int

func main() {
	exactNums = 0
	//bruteforce(0, 0, "START", "", 1)
	//brute(0, 0, 0, "+", "")
	//brute(0, 0, 0, "-", "")
	//brute(0, 0, 0, " ", "")
	createAllSequences()

	fmt.Printf("Answer: %v\n", exactNums)
}

func createAllSequences(){

	buildString := "1"

	for i, val := range numbers{
		sign := signs[i%3]
		buildString = buildString+sign+val
	}

}
/*
func brute(n int, sum int, oldNr int, sign string, oldSign string){

	if n < len(numbers){
		if sign == "+"{
			sum2 := sum
			sum2 += numbers[n]

			brute(n+1, sum2, numbers[n], "+", sign)
			brute(n+1, sum2, numbers[n], "-", sign)

			brute(n+1, sum, numbers[n], " ", sign)
		}else if sign == "-"{
			sum2 := sum - numbers[n]

			brute(n+1, sum2, numbers[n], "+", sign)
			brute(n+1, sum2, numbers[n], "-", sign)

			brute(n+1, sum, numbers[n], " ", sign)
		}else if sign == " "{
			combinedNum := oldNr*padding(numbers[n])+numbers[n]

			sum2 := sum
			if oldSign == "+"{
				sum2 += combinedNum
			}else if oldSign == "-"{
				sum2 -= combinedNum
			}
			if oldSign == ""{
				brute(n+1, combinedNum, combinedNum, "+", oldSign)
				brute(n+1, combinedNum, combinedNum, "-", oldSign)
			}else{
				brute(n+1, sum2, combinedNum, "+", oldSign)
				brute(n+1, sum2, combinedNum, "-", oldSign)
			}


			brute(n+1, sum, combinedNum, " ", oldSign)
		}



	}else if n == len(numbers){
		if sum == 42{
			exactNums++
		}
	}


}


func bruteforce(number int, sum int, nextSign string, prevSign string, prevNumber int){

	if number < len(numbers){

		if nextSign == "START"{
			newSum := numbers[number]

			bruteforce(number+1, newSum, "+", "", numbers[number])
			bruteforce(number+1, newSum, "-", "", numbers[number])

			bruteforce(number+1, sum, " ", "", numbers[number])

		}else if nextSign == "+"{
			newSum := sum + numbers[number]

			bruteforce(number+1, newSum, "+", "+", numbers[number])
			bruteforce(number+1, newSum, "-", "+", numbers[number])

			bruteforce(number+1, sum, " ", "+", numbers[number])

		}else if nextSign == "-"{
			newSum := sum - numbers[number]

			bruteforce(number+1, newSum, "+", "-", numbers[number])
			bruteforce(number+1, newSum, "-", "-", numbers[number])

			bruteforce(number+1, sum, " ", "-", numbers[number])

		}else if nextSign == " "{
			combinedNum := prevNumber*padding(numbers[number])+numbers[number]
			//combinedNum = numbers[number] * padding(prevNumber) + prevNumber

			newSum := sum
			if prevSign == "+" {
				newSum += combinedNum
			} else if prevSign == "-" {
				newSum -= combinedNum
			}

			bruteforce(number+1, newSum, "+", prevSign, combinedNum)
			bruteforce(number+1, newSum, "-", prevSign, combinedNum)

			bruteforce(number+1, sum, " ", prevSign, combinedNum)

		}
	}else{

		if nextSign == " "{
			if prevSign == "+"{
				sum += prevNumber
			}else if prevSign == "-"{
				sum -= prevNumber
			}
		}

		if sum == 42{
			exactNums++
		}

	}

}*/


func padding(n int) int {
	var p int = 1
	for p < n {
		p *= 10
	}
	return p
}