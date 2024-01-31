package ratings_controllers

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	models "api_tugas_minggu4/src/models/ratings_models" //Alias tulisan models disamping import
	"encoding/json"
	"fmt"
	"net/http"
)

func Data_ratings(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll_rating().Value)
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
		var product models.Ratings
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		models.Post_rating(&product)
		w.WriteHeader(http.StatusCreated)
		msg := map[string]string{
			"Message": "Product Created",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_rating(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	id := r.URL.Path[len("/rating/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select_rating(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Ratings
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		newProduct := models.Ratings{
			Member_id:  updateProduct.Member_id,
			Product_id: updateProduct.Product_id,
			Value:      updateProduct.Value,
		}

		models.Updates_rating(id, &newProduct)
		msg := map[string]string{
			"Message": "Product Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else if r.Method == "DELETE" {
		models.Delete_rating(id)
		msg := map[string]string{
			"Message": "Product Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
