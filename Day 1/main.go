package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	b, err := os.Open("input-vekksort.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	prevNumber := 0
	var numbers []int

	for r.Scan() {
		path := r.Text()


		nr, _ := strconv.Atoi(path)

		if nr >= prevNumber{
			prevNumber = nr
			numbers = append(numbers, nr)
		}
	}

	sum := 0

	for i := 0; i < len(numbers); i++{
		sum += numbers[i]
	}

	fmt.Printf("Svar: %v\n",sum)
}
