package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rain struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (rain Rain) slope() float64{
	return float64(rain.y2-rain.y1) / float64(rain.x2-rain.x1)
}

func main() {
	b, err := os.Open("input-rain.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var values = make(map[float64]int)

	for r.Scan() {
		line := r.Text()

		rain := Rain{}
		fmt.Sscanf(line,"(%d,%d);(%d,%d)\n", &rain.x1, &rain.y1, &rain.x2, &rain.y2)

		values[rain.slope()]++
	}

	mostLines := 0

	for _, count := range values{
		if count > mostLines{
			mostLines = count
		}
	}

	fmt.Printf("Svar: %v\n", mostLines)

}



