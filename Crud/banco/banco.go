package banco

import (
	"database/sql"
     "log"
   _ "github.com/go-sql-driver/mysql"
)

/// ABRE A CONEXAO COM O BANCO DE DADOS

func Connect() (*sql.DB , error) {
	strinDe := "root:123@/joel2?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql" , strinDe)
	if erro != nil {
		return nil , erro
	}

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	

	return db, nil
}