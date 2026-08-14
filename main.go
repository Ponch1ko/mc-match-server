package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var list_of_users []string //Список всех игроков
var reader *bufio.Reader   // Глобальный reader

const MAX_PLAYER_IN_MATCH = 4 // Максимальное колисество игроков на одном сервере

// Читает число из stdin.
func get_ans_from_user(reader *bufio.Reader) string {
	fmt.Print("Команда: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}

// функция new_user добавляет нового игрока в слайс всех игроков на сервере
func new_user() {
	fmt.Print("Введите имя игрока: ")
	user_name := get_ans_from_user(reader)
	list_of_users = append(list_of_users, user_name)
	fmt.Printf("Игрок '%s' добавлен (всего: %d)\n", user_name, len(list_of_users))
}

// функция all_users выводит всех игроков на сервере
func all_users() {
	size_of_list := len(list_of_users)

	if size_of_list == 0 {
		fmt.Println("На сервере нет игроков")
		return
	}
	for i := 0; i < size_of_list; i++ {
		fmt.Println(list_of_users[i])
	}
}

// функция new_match создаёт новый матч если на сервере больше или равно 4-м игроккам,
// в противном случае выводит ошибку
func new_match() {
	if len(list_of_users) < MAX_PLAYER_IN_MATCH {
		fmt.Printf("Ошибка!!! На сервере меньше %d-х игроков. Матч невозмодно начать", MAX_PLAYER_IN_MATCH)
	} else {
		list_of_users = list_of_users[MAX_PLAYER_IN_MATCH:]
		fmt.Println("Матч успешно запущен")
	}
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	list_of_users = []string{}
	for {
		ans_user := get_ans_from_user(reader)
		if ans_user == "quit" {
			fmt.Println("Выход")
			break
		}
		if ans_user == "ADD" {
			new_user()
		} else if ans_user == "list" {
			all_users()
		} else {
			new_match()
		}

	}
}
