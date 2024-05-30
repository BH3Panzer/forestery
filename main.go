package main

import "fmt"
import "math/rand"
import "time"

var VERSION string = "0.0.1"

var STATE string = "main_menu"

var LocNumber int

var PlayerName string

func skipLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Println()
	}

}

func main_menu() {
	Clear()
	fmt.Println("lol")
	fmt.Println(GetAnsiColor("green") + "Forestery v" + VERSION + GetAnsiColor(""))
	fmt.Println("*****************")
	skipLines(5)
	fmt.Println("1) Nouvelle partie")
	fmt.Printf("Choix : ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 1 {
		STATE = "new_game_menu_number"
	}
}

func new_game_menu_number() int {
	Clear()
	fmt.Println("Creer une nouvelle partie")
	fmt.Println("*****************")
	skipLines(5)
	fmt.Printf("Nombre de lieux : ")
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func set_name_menu() string {
	Clear()
	fmt.Println("Creer une nouvelle partie")
	fmt.Println("*****************")
	skipLines(5)
	fmt.Printf("Nom du joueur : ")
	var name string
	fmt.Scanln(&name)
	return name
}

type Location struct {
	id     int
	name   string
	temp   int
	height int
	typ    string
}

type Player struct {
	name   string
	pos    int
	pv     int
	pa     int
	shield int
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

func generateNewGame(nb int, name string) {
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
		fmt.Printf(GetAnsiColor("yellow"))
		fmt.Println(loc)
		fmt.Printf(GetAnsiColor(""))
	}
	player := Player{
		name:   name,
		pos:    int(nb / 2),
		pv:     100,
		pa:     0,
		shield: 0,
	}
	fmt.Printf(GetAnsiColor("yellow"))
	fmt.Println(player)
	fmt.Printf(GetAnsiColor(""))
}

func main() {
	Clear()
	if STATE == "main_menu" {
		main_menu()
	}
	if STATE == "new_game_menu_number" {
		LocNumber = new_game_menu_number()
		PlayerName = set_name_menu()
		generateNewGame(LocNumber, PlayerName)
	}
}
