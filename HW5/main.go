package main

import (
	"fmt"
	"hw5/internal/structs"
	"hw5/pkg/game_logic"
	"os"
)

func main() {
	field := map[int]string{
		1: " ",
		2: " ",
		3: " ",
		4: " ",
		5: " ",
		6: " ",
		7: " ",
		8: " ",
		9: " ",
	}

	Users := []structs.User{}
	users := game_logic.Intoduction()
	Users = append(Users, structs.User{Name: users[0], Mark: "X"}, structs.User{Name: users[1], Mark: "0"})
	fmt.Printf("Name=%s\n", Users[0].Name)
	game := structs.Game{Field: field}
	fmt.Println(game.Field[1])

	switch_var := true
	for i := 1; i < 10; i++ {
		if switch_var {
			game_logic.MakeMark(Users[0], field)
			game_logic.ShowStatus(field)
			winner := game_logic.CheckWinner(field)
			if winner != "n/a" {
				fmt.Printf("Winner is: %s\n", winner)
				os.Exit(0)
			}
			switch_var = !switch_var
		} else {

			game_logic.MakeMark(Users[1], field)
			game_logic.ShowStatus(field)
			winner := game_logic.CheckWinner(field)
			if winner != "n/a" {
				fmt.Printf("Winner is: %s\n", winner)
				os.Exit(0)
			}
			switch_var = !switch_var
		}

	}
	fmt.Printf("Game is finished. No one won")
	// winner := game_logic.CheckWinner(field)

	// fmt.Printf("Winner is: %s\n", winner)

}
