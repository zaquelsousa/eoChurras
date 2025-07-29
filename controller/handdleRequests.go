package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", getUsers).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	
	//rotas pro produto
	router.HandleFunc("/produtos", getProdutos).Methods("GET")
	router.HandleFunc("/produtos/{id}", getProduto).Methods("GET")
	router.HandleFunc("/produtos", createProduto).Methods("POST")
	router.HandleFunc("/produtos/{id}", updateProduto).Methods("PUT")
	router.HandleFunc("/produtos/{id}", deleteProduto).Methods("DELETE")
	
	//rotas pro pedidos
	router.HandleFunc("/pedidos", getPedidos).Methods("GET")
	router.HandleFunc("/pedidos/{id}", getPedido).Methods("GET")
	router.HandleFunc("/pedidos", createPedido).Methods("POST")
	router.HandleFunc("/pedidos/{id}", updatePedido).Methods("PUT")
	router.HandleFunc("/pedidos/{id}", deletePedido).Methods("DELETE")
	
	//rotas pro comandas
	router.HandleFunc("/comandas", getComandas).Methods("GET")
	router.HandleFunc("/comandas/{id}", getComanda).Methods("GET")
	router.HandleFunc("/comandas", createComanda).Methods("POST")
	router.HandleFunc("/comandas/{id}", updateComanda).Methods("PUT")
	router.HandleFunc("/comandas/{id}", deleteComanda).Methods("DELETE")

	
	//websocket routes
	router.HandleFunc("/ws", UpdatePedidos)

	log.Fatalln(http.ListenAndServe(":8081",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(router),
	))

}

