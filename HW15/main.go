package main

import (
	"HW15/observer"
	"HW15/pub_sub"
	"fmt"
)

func exercise1() {
	p := &pub_sub.Publisher{
		Subscribers: make(map[pub_sub.Subscriber]struct{}),
	}

	directory := "/tmp/test"

	go pub_sub.DetectChangesINDirectory(p, directory)

	sub := make(pub_sub.Subscriber)
	p.AddSubscriber(sub)

	go func() {
		for {
			select {
			case msg := <-sub:
				fmt.Println("getting changes:", msg)
			}
		}
	}()

	fmt.Println("Enter for exit")
	fmt.Scanln()

	p.RemoveSubscriber(sub)
}

func exercise2() {

	game := &observer.Game{
		Observers: make(map[observer.Observer]struct{}),
	}

	player1 := &observer.Player{Name: "Gamer 1"}
	player2 := &observer.Player{Name: "Gamer 2"}
	player3 := &observer.Player{Name: "Gamer 3"}
	player4 := &observer.Player{Name: "Gamer 4"}
	player5 := &observer.Player{Name: "Gamer 5"}

	game.Register(player1)
	game.Register(player2)
	game.Register(player3)
	game.Register(player4)
	game.Register(player5)

	game.Notify("Let's start!")

	game.Notify("Gamer 1 made move!")
	game.Notify("Gamer 2 is moving!")
	game.Notify("Gamer 3 is blocking!")
	game.Unregister(player4)
	game.Unregister(player5)

	game.Notify("good bye to Gamer 4 and 5")
}

func main() {
	// exercise1()
	exercise2()

}
