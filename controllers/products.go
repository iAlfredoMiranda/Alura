package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"curso.com/index/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {

			log.Println("Erro na conversão do preço:", err)

		}
		quantidadeConvertido, err := strconv.Atoi(quantidade)

		if err != nil {

			log.Println("Erro na conversão do quantidade:", err)

		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id para Int: ", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço para Float: ", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade para Int: ", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(w, r, "/", 301)

}
