package routes

import (
	"api_tugas_minggu4/src/controllers/category_controllers"
	members_controller "api_tugas_minggu4/src/controllers/members_controllers"
	"api_tugas_minggu4/src/controllers/order_controllers"
	"api_tugas_minggu4/src/controllers/products_controllers"
	"api_tugas_minggu4/src/controllers/ratings_controllers"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World ,Ini Adalah Halaman paling awal ya")
	})

	http.HandleFunc("/ratings", ratings_controllers.Data_ratings)
	http.HandleFunc("/rating/", ratings_controllers.Data_rating)
	http.HandleFunc("/orders", order_controllers.Data_orders)
	http.HandleFunc("/order/", order_controllers.Data_order)
	http.HandleFunc("/members", members_controller.Data_members)
	http.HandleFunc("/member/", members_controller.Data_member)
	http.HandleFunc("/products", products_controllers.Data_products)
	http.HandleFunc("/product/", products_controllers.Data_product)
	http.HandleFunc("/categories", category_controllers.Data_categories)
	http.HandleFunc("/category/", category_controllers.Data_category)
}
