package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Server struct {
	players     []string
	playerSet   map[string]bool
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
	if s.playerSet[name] {
		fmt.Printf("Игрок %s уже существует\n", name)
		return
	}

	s.players = append(s.players, name)
	s.playerSet[name] = true
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
	for _, name := range matchPlayers {
		delete(s.playerSet, name)
	}

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
			delete(s.playerSet, name)
			s.players = append(s.players[:i], s.players[i+1:]...)
			fmt.Printf("Игрок %q удалён\n", name)
			return
		}
	}
	fmt.Printf("Игрок %q не найден\n", name)
}

// Сохраняет очередь в players.json
func (s *Server) Save(filename string) error {
	data, err := json.MarshalIndent(s.players, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Загружает очередь из players.json
func (s *Server) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var players []string
	if err := json.Unmarshal(data, &players); err != nil {
		return err
	}
	s.players = players
	s.playerSet = make(map[string]bool)
	for _, name := range players {
		s.playerSet[name] = true
	}
	return nil
}

func main() {
	filename := "players.json"
	reader := bufio.NewReader(os.Stdin)

	server := Server{
		players:     []string{},
		maxPerMatch: 4,
		playerSet:   make(map[string]bool),
	}

	if err := server.Load(filename); err == nil {
		fmt.Printf("Загружено игроков: %d\n", len(server.players))
	}
	for {
		ansUser := readLine(reader, "Команда: ")
		ansUser = strings.ToLower(ansUser)
		switch ansUser {
		case "quit":
			fmt.Println("Выход")
			return
		case "add":
			name := readLine(reader, "Введите имя нового игрока: ")
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
		case "load":
			if err := server.Load(filename); err != nil {
				fmt.Printf("Ошибка при загрузке файла: %v\n", err)
			} else {
				fmt.Println("Очередь загружена")
			}
		case "save":
			if err := server.Save(filename); err != nil {
				fmt.Printf("Ошибка при сохранении: %v\n", err)
			} else {
				fmt.Println("Очередь сохранена")
			}
		default:
			fmt.Printf("Неизвестная команда: %q. Доступные: add, list, match, quit, stats, remove, save, load\n", ansUser)
		}
	}
}
