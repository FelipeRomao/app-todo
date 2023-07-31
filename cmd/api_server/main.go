package main

import (
	"fmt"
	"net/http"

	"github.com/FelipeRomao/todo/cmd/api_server/routes"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := chi.NewRouter()

	routes.SetRoutes(r)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
