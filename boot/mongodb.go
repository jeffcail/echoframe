package boot

import (
	"context"
	"fmt"

	"github.com/echoframe/common/db"

	"github.com/echoframe/conf"

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
