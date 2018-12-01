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
	sum := 0

	for r.Scan() {
		path := r.Text()


		nr, _ := strconv.Atoi(path)

		if nr >= prevNumber{
			prevNumber = nr
			sum += nr
		}
	}

	fmt.Printf("Svar: %v\n",sum)
}
