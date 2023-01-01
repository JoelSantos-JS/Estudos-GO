package servidor
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"crud/banco"
)

type usuarios struct {
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
 }

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao , erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corppo da requisicao"))
		return
	}

	var usuario usuarios

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := banco.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao converter ao Banco de dados"))
	}

	statement ,erro := db.Prepare("insert into usuarios (nome ,email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!") )
		return
	}

	defer statement.Close()

	insercao ,erro := statement.Exec(usuario.Nome , usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!") )
		return
	}

	idInserido , erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter id inserido"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso ID: %d", idInserido)) )

}