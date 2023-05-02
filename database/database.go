// Package database - Handles all interaction with ArangoDB
package database

import (
	"context"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = InitLogger() // setup the logger

// DBConnection is the structure that defined the database engine and collections
type DBConnection struct {
	Collection driver.Collection
	Database   driver.Database
}

var initDone = false          // has the data been initialized
var dbConnection DBConnection // database connection definition

// GetEnvDefault is a convenience function for handling env vars
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key) // get the env var
	if !ex {                     // not found return default
		return defVal
	}
	return val // return value for env var
}

// InitLogger sets up the Zap Logger to log to the console in a human readable format
func InitLogger() *zap.Logger {
	prodConfig := zap.NewProductionConfig()
	prodConfig.Encoding = "console"
	prodConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	prodConfig.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	logger, _ := prodConfig.Build()
	return logger
}

// InitializeDB is the function for connecting to the db engine, creating the database and collections
func InitializeDB() DBConnection {

	var db driver.Database
	var col driver.Collection
	var conn driver.Connection
	var client driver.Client
	var err error
	const databaseName = "ortelius"

	ctx := context.Background()

	if initDone {
		return dbConnection
	}

	dbhost := GetEnvDefault("ARANGO_HOST", "localhost")
	dbport := GetEnvDefault("ARANGO_PORT", "8529")
	dbuser := GetEnvDefault("ARANGO_USER", "root")
	dbpass := GetEnvDefault("ARANGO_PASS", "")
	dburl := GetEnvDefault("ARANGO_URL", "http://"+dbhost+":"+dbport)

	if conn, err = http.NewConnection(http.ConnectionConfig{Endpoints: []string{dburl}}); err != nil {
		logger.Sugar().Fatalf("Failed to create HTTP connection: %v", err)
	}

	_, err = conn.SetAuthentication(driver.BasicAuthentication(dbuser, dbpass))

	if err == nil {
		if client, err = driver.NewClient(driver.ClientConfig{Connection: conn}); err != nil {
			logger.Sugar().Fatalf("Failed to create Client: %v", err)
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
				logger.Sugar().Fatalf("Failed to create Database: %v", err)
			}
		} else {
			if db, err = client.CreateDatabase(ctx, databaseName, nil); err != nil {
				logger.Sugar().Fatalf("Failed to create Database: %v", err)
			}
		}

		exists, _ = db.CollectionExists(ctx, "evidence")
		if exists {
			if col, err = db.Collection(ctx, "evidence"); err != nil {
				logger.Sugar().Fatalf("Failed to use collection: %v", err)
			}
		} else {
			if col, err = db.CreateCollection(ctx, "evidence", nil); err != nil {
				logger.Sugar().Fatalf("Failed to create collection: %v", err)
			}
		}

		initDone = true

		dbConnection = DBConnection{
			Database:   db,
			Collection: col,
		}
	}
	return dbConnection
}

// PersistOnLTS interacts with the db abstraction microservice to
// store the cid/json data on NFT Storage or the OCI registry
func PersistOnLTS(cid2json map[string]string) {

	logger.Sugar().Infof("%+v\n", cid2json)
}

// FetchFromLTS interacts with the db abstraction microservice to
// fetch the json from NFT Storage or OCI registry
func FetchFromLTS(key string) (string, map[string]string) {

	msg := `{"objtype":"Domain","name":"GLOBAL"}`

	cid2json := make(map[string]string, 1)
	cid2json[key] = msg
	return key, cid2json
}
