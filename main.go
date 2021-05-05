package main

import (
	"flag"
	"net/http"

	"github.com/rs/cors"

	"github.com/almanalfaruq/alfarpos-backend/model"
	profileentity "github.com/almanalfaruq/alfarpos-backend/model/profile"
	"github.com/almanalfaruq/alfarpos-backend/repository"
	profilerepo "github.com/almanalfaruq/alfarpos-backend/repository/profile"
	"github.com/almanalfaruq/alfarpos-backend/routes"
	"github.com/almanalfaruq/alfarpos-backend/util"

	_ "github.com/almanalfaruq/alfarpos-backend/docs"
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

	// Populate data for payment and customer
	if shouldDropDB {
		golog.Info("Populating first data for payment and customer...")
		populateFirstData()
		golog.Info("Done populating data")
	}
}

func initRouter() {
	routes := routes.GetAllRoutes(&databaseConnection, config)
	http.Handle("/", routes)
	handler := cors.New(cors.Options{
		AllowedMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: true,
	}).Handler(routes)
	golog.Info("Server listening at http://localhost:8000")
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		golog.Fatal(err)
		panic(err)
	}
}

func populateFirstData() {
	payment := model.Payment{
		Name: "Cash",
	}
	paymentRepo := repository.NewPaymentRepo(&databaseConnection)
	_, err := paymentRepo.New(payment)
	if err != nil {
		panic(err)
	}

	customer := model.Customer{
		Name: "Customer",
	}
	customerRepo := repository.NewCustomerRepo(&databaseConnection)
	_, err = customerRepo.New(customer)
	if err != nil {
		panic(err)
	}

	profile := profileentity.Profile{
		Name:            "Toko Alfar",
		Address:         "Depan Pasar Ketaon, Banyudono, Boyolali",
		Phone:           "(0276) 3283720",
		ThankyouMessage: "Terima kasih atas kunjungan anda",
		FootNote:        "Barang yang sudah dibeli tidak dapat dikembalikan",
	}
	profileRepo := profilerepo.NewProfile(&databaseConnection)
	_, err = profileRepo.New(profile)
	if err != nil {
		panic(err)
	}
}
