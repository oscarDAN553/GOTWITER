package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/oscarDAN553/GOTWITER/bd"
	"github.com/oscarDAN553/GOTWITER/models"
)

/*GraboTweet permite grabar tweet en BD*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	_, status, err := bd.IncertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar incertal el registro intenta otra vez"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro incertar el tweet", 400)
	}
	w.WriteHeader(http.StatusCreated)
}
