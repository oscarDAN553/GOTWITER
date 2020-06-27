package bd

import (
	"context"
	"time"

	"github.com/oscarDAN553/GOTWITER/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es  la parada final de la BD para incertar los datos del user*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("gotwiter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
