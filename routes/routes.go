package routes

import (
	"net/http"

	"curso.com/index/controllers"
)

func LoudRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}