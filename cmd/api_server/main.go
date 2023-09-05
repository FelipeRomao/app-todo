package main

import (
	"fmt"
	"net/http"

	"github.com/FelipeRomao/todo/cmd/api_server/routes"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	routes.SetRoutes(r)

	fmt.Println("running server on port 8080")
	http.ListenAndServe(":8080", r)

}
