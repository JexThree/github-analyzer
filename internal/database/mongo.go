package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func ConnectMongo() {

	log.Println("Iniciando conexión MongoDB...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	client, err := mongo.Connect(
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Error haciendo ping a MongoDB:", err)
	}

	DB = client.Database(dbName)

	log.Println("MongoDB conectado correctamente")
}
