package database

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// DBConnection is the structure that defined the database engine and collections
type DBConnection struct {
	Collection driver.Collection
	Database   driver.Database
}

var initDone = false
var dbConnection DBConnection

// InitializeDB is the function for connecting to the db engine, creating the database and collections
func InitializeDB(ctx context.Context) DBConnection {

	var db driver.Database
	var col driver.Collection
	var conn driver.Connection
	var client driver.Client
	var err error
	const databaseName = "examples_books"

	if initDone {
		return dbConnection
	}

	if conn, err = http.NewConnection(http.ConnectionConfig{Endpoints: []string{"http://192.168.10.120:8529"}}); err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}

	conn.SetAuthentication(driver.BasicAuthentication("root", "rootpassword"))

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
	}
	return dbConnection
}
