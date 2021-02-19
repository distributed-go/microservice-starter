package migrations

import (
	"fmt"

	migrate "github.com/jobbox-tech/mongomigrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		fmt.Println(*db, "Apply migrations here...")
		return nil
	}, func(db *mongo.Database) error { //Down
		fmt.Println(*db, "Apply migrations here...")
		return nil
	})
}
