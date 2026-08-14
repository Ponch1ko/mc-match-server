package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var listOfUsers []string //Список всех игроков
var reader *bufio.Reader // Глобальный reader

const maxPlayersPerMatch = 4 // Максимальное колисество игроков на одном сервере

// Читает строку из stdin и возвращает её без пробелов по краям.
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}

// функция newUser добавляет нового игрока в слайс всех игроков на сервере
func newUser(name string) {
	listOfUsers = append(listOfUsers, name)
	fmt.Printf("Игрок '%s' добавлен (всего: %d)\n", name, len(listOfUsers))
}

// функция allUsers выводит всех игроков на сервере
func allUsers() {
	sizeOfList := len(listOfUsers)

	if sizeOfList == 0 {
		fmt.Println("На сервере нет игроков")
		return
	}
	for _, name := range listOfUsers {
		fmt.Println(name)
	}
}

// функция newMatch создаёт новый матч если на сервере больше или равно 4-м игроккам,
// в противном случае выводит ошибку
func newMatch() {
	if len(listOfUsers) < maxPlayersPerMatch {
		fmt.Printf("Ошибка: нужно не менее %d игроков. Сейчас на сервере: %d игроков\n", maxPlayersPerMatch, len(listOfUsers))
		return
	}
	matchPlayers := listOfUsers[:maxPlayersPerMatch]
	fmt.Println("Матч начинается! Игроки: ")
	for i, name := range matchPlayers {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
	listOfUsers = listOfUsers[maxPlayersPerMatch:]
	fmt.Printf("В очереди осталось: %d игроков\n", len(listOfUsers))
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	listOfUsers = []string{}
	for {
		ansUser := readLine(reader, "Команда: ")
		ansUser = strings.ToLower(ansUser)
		switch ansUser {
		case "quit":
			fmt.Println("Выход")
			return
		case "add":
			name := readLine(reader, "Введите имя нового игрока: ")
			newUser(name)
		case "list":
			allUsers()
		case "match":
			newMatch()
		default:
			fmt.Printf("Неизвестная команда: %q. Доступные: add, list, match, quit\n", ansUser)
		}
	}
}
