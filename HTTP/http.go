package main


import (
	"log"
	"net/http"
)

func main() {
		// HTTP É UM PROTOCOLO DE COMUNICAÇAO - BASE D A COMUNICAÇAO WEB

		// CLIETE FAZ ( A REQUISAÇAO ) - SERVIDOR (PROCESSA E ENVIA REQUISAO)

		// REQUEST - RESPONSE

		// ROTAS
		 	// URI - IDENTIFICADOR DE RECURSO
			// METODO - GET , POST , PUT , DELETE


			http.HandleFunc("/home" , func(w http.ResponseWriter, r *http.Request){
				w.Write([]byte("Ola mundo"))
			})
			http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request){
				w.Write([]byte("Pagina Principal"))
			})


			log.Fatal(http.ListenAndServe(":5000", nil))
}