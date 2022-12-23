package db

import "go.mongodb.org/mongo-driver/mongo"

var (
	Mongo *mongo.Client
)

func SetMongo(m *mongo.Client) {
	Mongo = m
}
