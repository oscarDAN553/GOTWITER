package models

/*Tweet captura el mensaje que llega del body */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
