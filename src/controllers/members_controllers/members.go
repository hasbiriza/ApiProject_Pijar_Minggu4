package members_controller

import (
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/middleware"
	models "api_tugas_minggu4/src/models/members_models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Data_members(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)
	if r.Method == "GET" {
		res, err := json.Marshal(models.SelectAll_member().Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "POST" {
		var member models.Member
		err := json.NewDecoder(r.Body).Decode(&member)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode")
			return
		}

		models.Post_member(&member)
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
			return
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func Data_member(w http.ResponseWriter, r *http.Request) {
	middleware.GetCleanedInput(r)
	helper.EnableCors(w)

	id := r.URL.Path[len("/member/"):]

	if r.Method == "GET" {
		res, err := json.Marshal(models.Select_member(id).Value)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "PUT" {
		var updateProduct models.Member
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, "Gagal Decode boss")
			return
		}

		newProduct := models.Member{
			Member_name: updateProduct.Member_name,
			Email:       updateProduct.Email,
			Password:    updateProduct.Password,
			Role:        updateProduct.Role,
			Address:     updateProduct.Address,
			Phone:       updateProduct.Phone,
		}

		models.Updates_member(id, &newProduct)
		msg := map[string]string{
			"Message": "Member Updated",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if r.Method == "DELETE" {
		models.Deletes_members(id)
		msg := map[string]string{
			"Message": "Member Deleted",
		}
		res, err := json.Marshal(msg)
		if err != nil {
			http.Error(w, "Gagal Konversi Json", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(res) // Perubahan disini
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else {
		http.Error(w, "method tidak diizinkan", http.StatusMethodNotAllowed)
	}
}
