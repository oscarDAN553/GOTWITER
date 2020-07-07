package models

/*RespuestaLogin tiene el token que se regresa con el login */
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
