package migrations

import (
	"fmt"

	migrate "github.com/eminetto/mongo-migrate"
	"github.com/globalsign/mgo"
)

func init() {
	migrate.Register(func(db *mgo.Database) error { //Up
		fmt.Println(*db, "Apply migrations here...")
		return nil
	}, func(db *mgo.Database) error { //Down
		fmt.Println(*db, "Apply migrations here...")
		return nil
	})
}
