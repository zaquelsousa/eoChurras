package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool { return true},
}

var clients = make(map[*websocket.Conn]bool)

type Mensagem struct {
	Tipo string      `json:"tipo"` 	
	Dado interface{} `json:"dado"`
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Erro ao conectar WebSocket:", err)
		return
	}
	clients[ws] = true

	go func() {
		defer func() {
			ws.Close()
			delete(clients, ws)
		}()
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				break
			}
		}
	}()
}


func BroadcastMensagem(tipo string, dado interface{}) {
	msg := Mensagem{
		Tipo: tipo,
		Dado: dado,
	}

	bytes, _ := json.Marshal(msg)

	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			fmt.Println("Erro ao enviar:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
