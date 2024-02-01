package category_controllers

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	models "api_tugas_minggu4/src/models/category_models" //Alias tulisan models disamping import
	"encoding/json"
	"fmt"
	"net/http"
)

func Data_categories(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll_category().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "POST" {
		var product models.Category
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		models.Post_category(&product)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_category(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/category/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select_category(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Category
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		newProduct := models.Category{
			Category_name: updateProduct.Category_name,
		}

		models.Updates_category(id, &newProduct)
		msg := map[string]string{
			"Message": "Product Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if r.Method == "DELETE" {
		models.Deletes_category(id)
		msg := map[string]string{
			"Message": "Product Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
