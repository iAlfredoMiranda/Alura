package main

import (
	"net/http"

	"curso.com/index/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.LoudRoutes()
	http.ListenAndServe(":8000", nil)
}
