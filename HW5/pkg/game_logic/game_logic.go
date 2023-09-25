package game_logic

import (
	"fmt"
	"hw5/internal/structs"
	"strconv"
	"time"
)

func Intoduction() (users [2]string) {
	var answer string
	fmt.Println("Welcome to tic tac toe")
Make_choise_first:
	fmt.Println("Enter name for X Player")
	fmt.Scanf("%s", &users[0])
	fmt.Printf("First player is: %s. Press 1 to confirm. Any other key to repeat.\n", users[0])
	fmt.Scanf("%s", &answer)
	if answer != "1" {
		time.Sleep(3 * time.Second)
		fmt.Print("Try again\n\n")
		goto Make_choise_first
	}
Make_choise_second:
	fmt.Println("Enter name for 0 Player")
	fmt.Scanf("%s", &users[1])
	fmt.Printf("Second player is: %s. Press 1 to confirm. Any other key to repeat.\n", users[1])
	fmt.Scanf("%s", &answer)
	if answer != "1" {
		time.Sleep(3 * time.Second)
		fmt.Print("Try again\n\n")
		goto Make_choise_second
	}
	return
}

func MakeMark(user structs.User, field map[int]string) {
	var turn, answer string

Make_Move:
	fmt.Printf("User: %s, make your turn\n", user.Name)
	fmt.Scanf("%s", &turn)
	fmt.Printf("Your turn is: %s. Press 1 to confirm. Any other key to repeat.\n", turn)
	fmt.Scanf("%s", &answer)
	if answer != "1" {
		time.Sleep(3 * time.Second)
		fmt.Print("Try again\n\n")
		goto Make_Move
	}
	mark, _ := strconv.Atoi(turn)
	if field[mark] != " " {
		fmt.Print("Cell is occupied, Try again\n\n")
		goto Make_Move
	}

	field[int(mark)] = user.Mark
}

func ShowStatus(field map[int]string) {
	fmt.Println("_______")
	for j := 0; j < 3; j++ {

		fmt.Println("|" + field[j*3+1] + "|" + field[j*3+2] + "|" + field[j*3+3] + "|")
	}
	fmt.Println("_______")
}

func CheckWinner(field map[int]string) (winner string) {
	winner = "n/a"
	if TestkWinner(field, "X") {
		winner = "X"
	}
	if TestkWinner(field, "0") {
		winner = "0"
	}
	return
}

func TestkWinner(field map[int]string, user string) (winner bool) {
	for i := 1; i < 4; i++ {
		if (field[i] == user && field[i+3] == user && field[i+6] == user) || (field[i] == user && field[i+1] == user && field[i+2] == user) {
			winner = true
		}
	}
	if (field[1] == user && field[5] == user && field[9] == user) || (field[3] == user && field[5] == user && field[7] == user) {
		winner = true
	}
	return
}
