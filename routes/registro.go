package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oscarDAN553/GOTWITER/bd"
	"github.com/oscarDAN553/GOTWITER/models"
)

/*Registro es el metodo que registra a un usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El imail del usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "debe especificar un pass de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un user con ese imail", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de user"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logro incertar el registro de user", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
