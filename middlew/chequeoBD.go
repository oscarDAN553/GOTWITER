package middlew

import (
	"net/http"

	"github.com/oscarDAN553/GOTWITER/bd"
)

/*ChequeoBD es el middlewer que pernite conocer el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
