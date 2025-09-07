package server

import "github.com/gorilla/websocket"

type Player struct {
	Name string
	ID   string
	Conn *websocket.Conn
}

type GameStatus struct {
	ID      string
	Message string
	Len     int
	Players map[string]*Player //ID of the player is the key
}

func NewPlayer(id, name string, conn *websocket.Conn) *Player {
	return &Player{
		ID:   id,
		Name: name,
		Conn: conn,
	}
}

func CreateGame(message string) *GameStatus {
	return &GameStatus{
		ID:      "id",
		Message: message,
		Len:     len(message),
		Players: make(map[string]*Player),
	}
}

func (g *GameStatus) AddPlayer(p *Player) {
	g.Players[p.ID] = p
}
