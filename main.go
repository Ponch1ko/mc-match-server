package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Server struct {
	players     []string
	maxPerMatch int
}

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

// AddUser добавляет игрока в очередь.
func (s *Server) AddUser(name string) {
	s.players = append(s.players, name)
	fmt.Printf("Игрок '%s' добавлен (всего: %d)\n", name, len(s.players))
}

// функция AllUsers выводит всех игроков на сервере
func (s *Server) AllUsers() {

	if len(s.players) == 0 {
		fmt.Println("На сервере нет игроков")
		return
	}
	for i, name := range s.players {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
}

// NewMatch собирает матч из первых maxPerMatch игроков.
func (s *Server) NewMatch() {
	if len(s.players) < s.maxPerMatch {
		fmt.Printf("Ошибка: нужно не менее %d игроков. Сейчас на сервере: %d игроков\n", s.maxPerMatch, len(s.players))
		return
	}
	matchPlayers := s.players[:s.maxPerMatch]
	fmt.Println("Матч начинается! Игроки: ")
	for i, name := range matchPlayers {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
	s.players = s.players[s.maxPerMatch:]
	fmt.Printf("В очереди осталось: %d игроков\n", len(s.players))
}

// Stats выводит количество игроков в очереди.
func (s *Server) Stats() {
	fmt.Printf("Количество игроков: %d человек \n", len(s.players))
}

// Remove удаляет игрока из очереди по имени.
func (s *Server) Remove(name string) {
	for i, player := range s.players {
		if player == name {
			s.players = append(s.players[:i], s.players[i+1:]...)
			fmt.Printf("Игрок %q удалён\n", name)
			return
		}
	}
	fmt.Printf("Игрок %q не найден\n", name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	server := Server{
		players:     []string{},
		maxPerMatch: 4,
	}
	for {
		ansUser := readLine(reader, "Команда: ")
		ansUser = strings.ToLower(ansUser)
		switch ansUser {
		case "quit":
			fmt.Println("Выход")
			return
		case "add":
			server.AddUser(name)
		case "list":
			server.AllUsers()
		case "match":
			server.NewMatch()
		case "stats":
			server.Stats()
		case "remove":
			name := readLine(reader, "Имя человека, которого надо удалить: ")
			server.Remove(name)
		default:
			fmt.Printf("Неизвестная команда: %q. Доступные: add, list, match, quit, stats, remove\n", ansUser)
		}
	}
}

