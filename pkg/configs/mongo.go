package configs

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

func Connect() error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://new_user1:ueWjO2knNNZVAW2H@cluster0.patmm.mongodb.net/max-trade?retryWrites=true&w=majority"))

	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("max-trade")

	if err != nil {
		return err
	}

	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
