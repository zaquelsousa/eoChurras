package controller

import (
	"churras/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool { return true},
}
//lista de conexoes ja que vait er os gaçons os churrasqueiros vendo os pedios em tempo real ne
var clients = make(map[*websocket.Conn]bool)

func UpdatePedidos(w http.ResponseWriter, r *http.Request){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	clients[ws] = true
}

func BroadcastNovoPedido(pedido models.Pedido) {
	msg, _ := json.Marshal(pedido)
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, msg) 
		if err != nil{
			fmt.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}


