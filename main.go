package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"seleksi-compfest-backend/controller"
	"seleksi-compfest-backend/database"
	"seleksi-compfest-backend/entity"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	initDB()
	// initFileSystem()

	router := mux.NewRouter().StrictSlash(true)
	initializeHandlers(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000", "https://kantin-kejujuran-dimasfm.herokuapp.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	entity.StoreBalance = entity.Balance{Amount: 0}

	handler := c.Handler(router)
	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}).Methods("GET")
	router.HandleFunc("/product", controller.ReadProduct).Methods("GET")
	router.HandleFunc("/product", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", controller.ReadProductById).Methods("GET")
	router.HandleFunc("/product/{id}", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", controller.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product", controller.DeleteAllProduct).Methods("DELETE")
	router.HandleFunc("/pid/", controller.ReadProductId).Methods("GET")
	router.HandleFunc("/img/product/{id}", controller.GetProductImg).Methods("GET")

	router.HandleFunc("/balance", controller.GetBalance).Methods("GET")
	router.HandleFunc("/balance", controller.UpdateBalance).Methods("POST")
	router.HandleFunc("/balance/add", controller.AddBalance).Methods("POST")
	router.HandleFunc("/balance/substract", controller.SubstractBalance).Methods("POST")
}

func initFileSystem() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory: ", err)
	}
	assetsImagePath := filepath.Join(dir, "assets/images")
	err = os.Mkdir(assetsImagePath, 0700)
	if err != nil {
		panic("Error creating img directory: " + err.Error())
	}
}

func initDB() {
	var dbUrl string
	if _, present := os.LookupEnv("DATABASE_URL"); !present {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}
	dbUrl = os.Getenv("DATABASE_URL")
	log.Println("Database url: ", dbUrl)

	err := database.Connect(dbUrl)
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}
	database.MigrateProduct(&entity.Product{})
}
