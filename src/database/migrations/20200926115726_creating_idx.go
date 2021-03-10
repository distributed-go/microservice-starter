package migrations

import (
	"context"
	"time"

	migrate "github.com/jobbox-tech/mongomigrate"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"Email": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(viper.GetString("db.recruiters_collection")).Indexes().CreateOne(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(viper.GetString("db.recruiters_collection")).Indexes().DropOne(ctx, "Email_1")
		return err
	})
}
