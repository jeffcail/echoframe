package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri        = "mongodb://127.0.0.1:27017/?maxPoolSize=20&w=majority"
	mon        *mongo.Client
	dataBase   = "echo-scaffolding" // 数据库
	collection = "restaurants"
)

type Restaurant struct {
	Name         string        `bson:"name,omitempty"`
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}

func NewMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	fmt.Println("successfully connected and pinged.")
	return client
}

// 单个写入
func TestInsertOne(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	newRestaurant := Restaurant{Name: "8282", Cuisine: "korean"}
	result, err := coll.InsertOne(context.TODO(), newRestaurant)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Document inserted with ID： %s\n", result.InsertedID)
}

// 批量写入
func TestInsertMultiple(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	newRestaurants := []interface{}{
		Restaurant{Name: "Rule of thirds", Cuisine: "Japanese"},
		Restaurant{Name: "Madame Vo", Cuisine: "Vietnamese"},
	}
	result, err := coll.InsertMany(context.TODO(), newRestaurants)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%d documents inseted with IDs:\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
}

// 单个查找
func TestFindOne(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	filter := bson.D{{"name", "Rule of thirds"}}
	var result Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("not match any documents")
		}
		t.Fatal(err)
	}
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", output)
}

// 批量查找
func TestFindMultiple(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	filter := bson.D{{"cuisine", "Japanese"}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		t.Fatal(err)
	}

	var results []Restaurant
	if err = cursor.All(context.TODO(), &results); err != nil {
		t.Fatal(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	}
}

// 单个更新
func TestUpdateOne(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	id, _ := primitive.ObjectIDFromHex("63a5696412b571d2b15db3e4")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"avg_rating", 4.4}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
}

// 统计条数
func TestCount(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	filter := bson.D{{"name", "8282"}}
	estCount, err := coll.EstimatedDocumentCount(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Estimated number of documents in the restaurants collection: %d\n", estCount)
	fmt.Printf("Number of name Japanese 8282: %d\n", count)
}

// 单个删除
func TestDelete(t *testing.T) {
	mon = NewMongoDB()
	coll := mon.Database(dataBase).Collection(collection)
	filter := bson.D{{"name", "8282"}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Documents deleted: %d\n", result.DeletedCount)
}
