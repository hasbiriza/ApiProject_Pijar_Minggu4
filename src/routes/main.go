package routes

import (
	"api_tugas_minggu4/src/controllers/category_controllers"
	customers_controller "api_tugas_minggu4/src/controllers/customers_controllers"
	"api_tugas_minggu4/src/controllers/products_controllers"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World ,Ini Adalah Halaman paling awal ya")
	})

	http.HandleFunc("/customers", customers_controller.Data_customers)
	http.HandleFunc("/customer", customers_controller.Data_customer)
	http.HandleFunc("/products", products_controllers.Data_products)
	http.HandleFunc("/product/", products_controllers.Data_product)
	http.HandleFunc("/categories", category_controllers.Data_categories)
	http.HandleFunc("/category", category_controllers.Data_category)
}
