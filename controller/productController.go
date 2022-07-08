package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"seleksi-compfest-backend/database"
	"seleksi-compfest-backend/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Println("Error parsing multipart form: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Mengambil file dari form
	uploadedFile, handler, err := r.FormFile("productImage")
	if err != nil {
		log.Println("Error getting File: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Mengambil nama file
	filename := handler.Filename

	//Membuat id
	id := uuid.New().String()

	// Membuat file baru dengan nama file yang baru
	newFile, err := os.Create(dir + "/assets/images/" + id + "-" + filename)
	if err != nil {
		log.Println("Error creating new file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	if _, err := io.Copy(newFile, uploadedFile); err != nil {
		log.Println("Error copying file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	price, _ := strconv.ParseFloat(r.FormValue("productPrice"), 64)

	product := entity.Product{
		ID:          id,
		Name:        r.FormValue("productName"),
		Price:       price,
		Image:       id + "-" + filename,
		Description: r.FormValue("productDescription"),
		Date:        time.Now(),
	}

	if result := database.Connector.Create(&product); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProductImg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var product entity.Product

	if result := database.Connector.Where("id = ?", key).Find(&product); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	img, err := os.Open(dir + "/assets/images/" + product.Image)
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	w.WriteHeader(http.StatusOK)
	io.Copy(w, img)

}

func ReadProduct(w http.ResponseWriter, r *http.Request) {
	var products []entity.Product

	database.Connector.Find(&products)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

type userId struct {
	Id string `json:"id"`
}

func ReadProductId(w http.ResponseWriter, r *http.Request) {
	var ids []userId

	database.Connector.Model(&entity.Product{}).Find(&ids)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ids)
}

func ReadProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var product entity.Product

	if result := database.Connector.Where("id = ?", key).Find(&product); result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)

	if result := database.Connector.Save(&product); result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product entity.Product
	var getProduct entity.Product
	database.Connector.Where("id = ?", id).Find(&getProduct)
	database.Connector.Where("id = ?", id).Delete(&product)
	dir, _ := os.Getwd()
	image := filepath.Join(dir, "assets/images/"+getProduct.Image)
	os.Remove(image)
	w.WriteHeader(http.StatusNoContent)
}

func DeleteAllProduct(w http.ResponseWriter, r *http.Request) {
	database.Connector.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.Product{})
	dir, _ := os.Getwd()
	imagesPath := filepath.Join(dir, "assets/images")
	os.RemoveAll(imagesPath)
	os.MkdirAll(imagesPath, 0700)
	w.WriteHeader(http.StatusNoContent)
}
