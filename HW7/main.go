package main

import (
	exercise1 "github.com/matkovskiy/golang_course/HW7/pkg/exercise1"
	"github.com/matkovskiy/golang_course/HW7/pkg/exercise2"
)

func exercise_1() {
	length := 100
	random_range := 100
	ch := make(chan int)
	avg := make(chan float64)
	display := make(chan float64)
	go exercise1.GenerateRandom(ch, length, random_range)
	go exercise1.AverageValue(ch, avg, length)
	display = avg
	exercise1.PrintAverage(display)

}

func exercise_2() {
	length := 100
	random_range := 100
	ch := make(chan int)
	results := make(chan int)
	wait := make(chan bool)
	go exercise2.GenerateRandom(ch, length, random_range, results, wait)
	go exercise2.MaxMin(ch, results)

	<-wait
}

func main() {
	// exercise_1()
	exercise_2()
	// Select2()
}
