package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg *mongo.Database

func Connect() error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://new_user1:ueWjO2knNNZVAW2H@cluster0.patmm.mongodb.net/max-trade?retryWrites=true&w=majority"))

	if err != nil {
		fmt.Println("Error connecting 0")
		return err
	}

	// Defer disconnect when function exits
	// defer func() {
	// 	if err := client.Disconnect(context.Background()); err != nil {
	// 		fmt.Println("Error disconnecting:", err)
	// 	}
	// }()

	// Ping to test connection (optional)
	if err := client.Ping(context.Background(), nil); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}

	if err != nil {
		fmt.Println("Error connecting")
		return err
	}
	Mg = client.Database("max-trade")
	fmt.Println("Connected")

	return nil
}

func GetConnect() *mongo.Database {
	return Mg
}
