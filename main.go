package main

import (
	"flag"
	"net/http"

	"github.com/rs/cors"

	"github.com/almanalfaruq/alfarpos-backend/routes"
	"github.com/almanalfaruq/alfarpos-backend/util"

	_ "github.com/almanalfaruq/alfarpos-backend/docs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/golog"
)

var config util.Config
var databaseConnection util.DBConn

// @title AlfarPOS BackEnd
// @version 1.0
// @description This is a backend server for alfarpos.

// @contact.name Almantera Tiantana Al Faruqi
// @contact.url https://twitter.com/almanalfaruq
// @contact.email alman.alfaruq@gmail.com

// @host localhost:8080
// @BasePath /api/
// @query.collection.format multi

func main() {
	var shouldDropDB bool
	flag.BoolVar(&shouldDropDB, "drop", false, "flag to drop db")
	flag.Parse()

	initMigration(shouldDropDB)
	initRouter()

	defer databaseConnection.Close()
}

func initMigration(shouldDropDB bool) {
	err := config.Read("./config.yaml", &config)
	if err != nil {
		panic(err)
	}
	golog.Info("Connecting to database...")
	databaseConnection.Open(config)
	golog.Info("Connected!")
	// Database Drop
	if shouldDropDB {
		golog.Warn("Dropping database...")
		databaseConnection.DropDb()
		golog.Info("Dropped!")
	}

	// Database Migration
	golog.Warn("Migrating database...")
	databaseConnection.MigrateDb()
	golog.Info("Migrated!")
}

func initRouter() {
	routes := routes.GetAllRoutes(&databaseConnection, config)
	http.Handle("/", routes)
	handler := cors.Default().Handler(routes)
	golog.Info("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		golog.Fatal(err)
		panic(err)
	}
}
