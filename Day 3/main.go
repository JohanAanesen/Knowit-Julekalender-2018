package main

import (
	"fmt"
)

var juletall = make(map[float64]bool)
var primtall = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 293, 307, 311, 313, 317, 331, 337, 347, 349,
	353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457, 461, 463,
	467, 479, 487, 491, 499, 503, 509}
var max = 4294967296
var hitSameCount = 0

func main() {
	solveJuletall(float64(1), 1)
}

func solveJuletall(digit float64, number float64){


	for _, val := range primtall{
		nr := number
		nr *= float64(val)
		if int(nr) > max{
			break
		}


		if digit == 24{
			if juletall[nr]{
				hitSameCount++
			}else{
				hitSameCount = 0
				juletall[nr] = true
			}

			if hitSameCount > 10000000{ //after 10 milion numbers where no one is new, we should ish have our answer
				fmt.Printf("Answer: %d\n", len(juletall)) //wait a bit until it doesnt iterate anymore
			}
			break
		}else{
			solveJuletall(digit+1, nr)
		}

	}

}
