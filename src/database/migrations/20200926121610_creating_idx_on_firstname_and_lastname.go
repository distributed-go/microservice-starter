package migrations

import (
	"context"
	"time"

	migrate "github.com/jobbox-tech/mongomigrate"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mod := []mongo.IndexModel{
		{
			Keys: bson.M{"FirstName": 1},
		},
		{
			Keys: bson.M{"LastName": 1},
		},
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(viper.GetString("db.recruiters_collection")).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(viper.GetString("db.recruiters_collection")).Indexes().DropOne(ctx, "FirstName_1")
		_, err = db.Collection(viper.GetString("db.recruiters_collection")).Indexes().DropOne(ctx, "LastName_1")
		return err
	})
}
