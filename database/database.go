// Package database - Handles all interaction with ArangoDB
package database

import (
	"context"
	"fmt"
	"log"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// DBConnection is the structure that defined the database engine and collections
type DBConnection struct {
	Collection driver.Collection
	Database   driver.Database
	Context    context.Context
}

var initDone = false
var dbConnection DBConnection

func getEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

// InitializeDB is the function for connecting to the db engine, creating the database and collections
func InitializeDB() DBConnection {

	var db driver.Database
	var col driver.Collection
	var conn driver.Connection
	var client driver.Client
	var err error
	const databaseName = "examples_books"

	ctx := context.Background()

	if initDone {
		return dbConnection
	}

	dburl := getEnvDefault("ARGANGO_URL", "http://localhost:8529")

	if conn, err = http.NewConnection(http.ConnectionConfig{Endpoints: []string{dburl}}); err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	_, err = conn.SetAuthentication(driver.BasicAuthentication("root", "rootpassword"))

	if err == nil {
		if client, err = driver.NewClient(driver.ClientConfig{Connection: conn}); err != nil {
			log.Fatalf("Failed to create Client: %v", err)
		}

		exists := false
		dblist, _ := client.Databases(ctx)

		for _, dbinfo := range dblist {
			if dbinfo.Name() == databaseName {
				exists = true
				break
			}
		}

		if exists {
			if db, err = client.Database(ctx, databaseName); err != nil {
				log.Fatalf("Failed to create Database: %v", err)
			}
		} else {
			if db, err = client.CreateDatabase(ctx, databaseName, nil); err != nil {
				log.Fatalf("Failed to create Database: %v", err)
			}
		}

		exists, _ = db.CollectionExists(ctx, "books")
		if exists {
			if col, err = db.Collection(ctx, "books"); err != nil {
				log.Fatalf("Failed to use collection: %v", err)
			}
		} else {
			if col, err = db.CreateCollection(ctx, "books", nil); err != nil {
				log.Fatalf("Failed to create collection: %v", err)
			}
		}

		initDone = true

		dbConnection = DBConnection{
			Database:   db,
			Collection: col,
			Context:    ctx,
		}
	}
	return dbConnection
}

// PersistOnLTS interacts with the db abstraction microservice to
// store the cid/json data on NFT Storage or the OCI registry
func PersistOnLTS(cid2json map[string]string) {

	fmt.Printf("%+v\n", cid2json)
}

// FetchFromLTS interacts with the db abstraction microservice to
// fetch the json from NFT Storage or OCI registry
func FetchFromLTS(key string) (string, map[string]string) {

	msg := `{"objtype":"Domain","name":"GLOBAL"}`

	cid2json := make(map[string]string, 1)
	cid2json[key] = msg
	return key, cid2json
}
