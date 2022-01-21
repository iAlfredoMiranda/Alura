package models

import "curso.com/index/db"

type Product struct {
	Id, Quantity      int
	Name, Discription string
	Price             float64
}

func SearchAllProducts() []Product {

	db := db.ConnectDB()

	selectAllProducts, err := db.Query("SELECT * FROM produtos order by id asc")

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

func EditaProduto(id string) Product {
	db := db.ConnectDB()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Product{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Name = nome
		produtoParaAtualizar.Discription = descricao
		produtoParaAtualizar.Price = preco
		produtoParaAtualizar.Quantity = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {

	db := db.ConnectDB()
	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, discricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
