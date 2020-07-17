package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*MongoCN es uin objeto de coneccion a la base de datos*/
var MongoCN = ConectarDB()
var clientOptins = options.Client().ApplyURI("mongodb+srv://Carlos:Cntja12xxx@cluster0.l3iea.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority")
//ConectarBD
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptins)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexicon valida")
	return client

}

/*CheckConnection ping bd*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}

