package main

import "fmt"
import "github.com/inancgumus/screen"
import "math/rand"
import "time"

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
	id     int
	name   string
	temp   int
	height int
	typ    string
}

func getRandomNumber(nb int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(nb)
}

func getRandomType() string {
	typeList := [...]string{"mou", "mea", "val", "mar"}
	choice := getRandomNumber(4)
	return typeList[choice]
}

func generateName(lenght int) string {
	cons := [...]string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Z"}
	voy := [...]string{"A", "E", "I", "O", "U", "Y"}
	var name string
	var choice int
	for i := 0; i < lenght; i++ {
		choice = getRandomNumber(20)
		name = name + string(cons[choice])
		choice = getRandomNumber(6)
		name = name + string(voy[choice])
	}
	return name
}

func generateNewGame(nb int) {
	locations := make([]Location, nb)
	for i := 0; i < nb; i++ {
		loc := Location{
			id:     i,
			name:   generateName(3),
			typ:    getRandomType(),
			height: getRandomNumber(500),
			temp:   getRandomNumber(40),
		}
		locations[i] = loc
		fmt.Println(loc)
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
