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
	mod := []mongo.IndexModel{
		{
			Keys: bson.M{"AccountID": 1},
		},
		{
			Keys:    bson.M{"TokenUUID": 1},
			Options: options.Index().SetUnique(true),
		},
	}
	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(viper.GetString("db.access_tokens_collection")).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(viper.GetString("db.access_tokens_collection")).Indexes().DropOne(ctx, "AccountID_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(viper.GetString("db.access_tokens_collection")).Indexes().DropOne(ctx, "TokenUUID_1")
		return err
	})
}
