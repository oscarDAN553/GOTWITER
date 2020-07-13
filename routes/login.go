package routes

import (
	"encoding/json"
	"net/http"
	"time"

	//"github.com/oscarDAN553/GOTWITER/bd"
	"github.com/oscarDAN553/GOTWITER/bd"
	"github.com/oscarDAN553/GOTWITER/jwt"
	"github.com/oscarDAN553/GOTWITER/models"
)

/*Login regresa el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "usuario y/o contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "usuario o contraseña invalido", 400)
		return
	}
	jwtkey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar generar el token"+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
