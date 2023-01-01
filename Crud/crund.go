package main 



import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"crud/servidor"
)


func main() {

		// CREATE = POST
		// READ = GET
		// UPDATE = PUT
		// DELETE  = DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) 
	fmt.Println("Escutando na porta 500")
	log.Fatal(http.ListenAndServe(":5000", router))

}