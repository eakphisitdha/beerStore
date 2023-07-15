package transaction

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --with-expecter --name "ITransaction" --output $PWD/mocks
type ITransaction interface {
	Log(fields []string, values []interface{}, user string, id int, action string) error
}

type Transaction struct {
	collection *mongo.Collection
}

func NewTransaction(mongoDB *mongo.Database) ITransaction {
	return &Transaction{collection: mongoDB.Collection("logs")}
}

func (t *Transaction) Log(fields []string, values []interface{}, user string, id int, action string) error {

	// create log in MongoDB
	doc := bson.D{{Key: "time", Value: time.Now().Format("2-Jan-06 03:04PM")}, {Key: "user", Value: user}, {Key: "action", Value: action}, {Key: "id", Value: id}}
	if fields != nil {
		for i, field := range fields {
			doc = append(doc, bson.E{Key: field, Value: values[i]})
		}
	}

	_, err := t.collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
