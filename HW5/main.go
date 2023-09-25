package main

import (
	"fmt"
	"hw5/internal/structs"
	"hw5/pkg/game_logic"
)

func main() {
	field := map[int]string{
		1: "0",
		2: "0",
		3: "0",
		4: "-",
		5: " ",
		6: " ",
		7: "X",
		8: "8",
		9: "9",
	}

	Users := []structs.User{}
	users := game_logic.Intoduction()
	Users = append(Users, structs.User{Name: users[0], Mark: "X"}, structs.User{Name: users[1], Mark: "0"})
	fmt.Printf("Name=%s\n", Users[0].Name)
	game := structs.Game{Field: field}
	fmt.Println(game.Field[1])

	game_logic.MakeMark(Users[0], field)

	game_logic.ShowStatus(field)
	winner := game_logic.CheckWinner(field)
	fmt.Printf("Winner is: %s\n", winner)

}
