package products_controllers

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	models "api_tugas_minggu4/src/models/products_models" //Alias tulisan models disamping import
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// /////////////////////////////Upload_File////////////////////////////////////////////////
func Handle_upload(w http.ResponseWriter, r *http.Request) {
	const (
		AllowedExtensions = ".jpg,.jpeg,.pdf,.png"
		MaxFileSize       = 2 << 20 //2MB
	)
	//Memerikan method request, harus POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Mendapatkan File dari form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(handler.Filename)
	ext = strings.ToLower(ext)
	allowedExts := strings.Split(AllowedExtensions, ",")
	validExtension := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			validExtension = true
			break
		}
	}
	if !validExtension {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
		return
	}

	//Mengecek ukuran file
	fileSize := handler.Size
	if fileSize > MaxFileSize {
		http.Error(w, "File size exceeds the allowed limit", http.StatusBadRequest)
		return
	}

	timeStamp := time.Now().Format("20060102_150405")

	//Membuat nama unik untuk file
	filename := fmt.Sprintf("src/uploads/%s_%s", timeStamp, handler.Filename)

	//Membuat file untuk menyimpan gambar

	out, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	//Menyalin isi file yang diuopload ke file yang baru dibuat
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Menyampaikan respons berhasil
	msg := map[string]string{
		"Message": "File uploaded succesfully",
	}
	res, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}

	w.Write(res)

}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")
	res, err := json.Marshal(models.FindData(keyword).Value)
	if err != nil {
		http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		return
	}
	w.Write(res)

}

// ////////////////////CRUD PRODUCTS//////////////////////////////////////////////
func Data_products(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll_product().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "POST" {
		var product models.Products
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		models.Post_product(&product)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_product(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/product/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select_product(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Products
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		newProduct := models.Products{
			Product_name: updateProduct.Product_name,
			Price:        updateProduct.Price,
			Color:        updateProduct.Color,
			Size:         updateProduct.Size,
			Stock:        updateProduct.Stock,
			Description:  updateProduct.Description,
			Condition:    updateProduct.Condition,
		}

		models.Updates_products(id, &newProduct)
		msg := map[string]string{
			"Message": "Product Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "DELETE" {
		models.Delete_products(id)
		msg := map[string]string{
			"Message": "Product Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
