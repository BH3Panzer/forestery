package main

import "fmt"
import "github.com/inancgumus/screen"

var VERSION string = "0.0.1"

var STATE string = "main_menu"

var LocNumber int

func skipLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Println()
	}

}

func main_menu() {
	screen.MoveTopLeft()
	fmt.Println("Forestery v" + VERSION)
	fmt.Println("*****************")
	skipLines(5)
	fmt.Println("1) Nouvelle partie")
	fmt.Printf("Choix : ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		STATE = "new_game_menu_number"
	}
	screen.Clear()
}

func new_game_menu_number() int {
	screen.MoveTopLeft()
	fmt.Println("Creer une nouvelle partie")
	fmt.Println("*****************")
	skipLines(5)
	fmt.Printf("Nombre de lieux : ")
	var choice int
	fmt.Scanln(&choice)
	screen.Clear()
	return choice
}

type Location struct {
	id int
	name string
	temp int
	height int
	typ string
}

func generateNewGame(nb int) {
	locations := make([]Location, nb)
	for i := 0; i < nb; i++ {
		loc := Location{
			id: i,
			name: "",
		}
		locations[i] = loc
	}
}

func main() {
	screen.Clear()
	if STATE == "main_menu" {
		main_menu()
	}
	if STATE == "new_game_menu_number" {
		LocNumber = new_game_menu_number()
		STATE = "generate_new_game"
		generateNewGame(LocNumber)
	}
}


