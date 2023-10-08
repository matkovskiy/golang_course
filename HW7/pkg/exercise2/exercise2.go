package exercise2

import (
	"fmt"
	"math/rand"
	"sort"
)

func GenerateRandom(ch chan<- int, count int, random_range int, results chan int, wait chan bool) {
	for i := 0; i < count; i++ {
		ch <- rand.Intn(random_range)
	}
	close(ch)
	fmt.Printf("Min:%d\n", <-results)
	fmt.Printf("Max:%d\n", <-results)
	wait <- true
}

func MaxMin(ch chan int, result chan int) {
	var numbers []int
	for i := range ch {
		numbers = append(numbers, i)
	}
	sort.Ints(numbers)
	result <- numbers[0]
	result <- numbers[len(numbers)-1]
}
