package boot

import (
	"context"
	"fmt"

	"github.com/echo-scaffolding/common/db"

	"github.com/echo-scaffolding/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitMongoDB
func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(conf.Config.MongoDB)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	fmt.Println("successfully connected and pinged.")
	db.SetMongo(client)
}
