package exercise1

import (
	"fmt"
	"math/rand"
)

func GenerateRandom(ch chan int, count int, random_range int) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(random_range)
	}
	close(ch)
}

func AverageValue(ch chan int, avg chan float64, count int) {
	var sum int
	for i := range ch {
		sum = sum + i
	}
	avg <- float64(sum) / float64(count)
	fmt.Printf("sum=%d, len=%d\n", sum, count)
}

func PrintAverage(ch chan float64) {
	fmt.Printf("AvarageValue=%g", <-ch)
	close(ch)
}
