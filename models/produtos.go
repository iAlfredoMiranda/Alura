package models

import "curso.com/index/db"

type Product struct {
	Id, Quantity      int
	Name, Discription string
	Price             float64
}

func SearchAllProducts() []Product {

	db := db.ConnectDB()

	selectAllProducts, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = nome
		p.Discription = descricao
		p.Price = preco
		p.Quantity = quantidade

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	insertDataDB, err := db.Prepare("insert into produtos(nome, discricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataDB.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

func DeletaProduto(id string) {

	db := db.ConnectDB()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)

	defer db.Close()

}
