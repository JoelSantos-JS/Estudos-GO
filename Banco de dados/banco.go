package main

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)





func main() {
	urlConexao := "root:123@/joel?charset=utf8&parseTime=True&loc=Local"

	db , erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conecao Abeerta")

   linhas, erro :=	db.Query("select *from books")
   if erro != nil {
	log.Fatal(erro)
   }

   defer linhas.Close()

   fmt.Println(linhas)
}