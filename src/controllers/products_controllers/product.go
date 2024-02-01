package products_controllers

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	models "api_tugas_minggu4/src/models/products_models" //Alias tulisan models disamping import
	"encoding/json"
	"fmt"
	"net/http"
)

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
