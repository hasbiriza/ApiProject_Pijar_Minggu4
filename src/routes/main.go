package routes

import (
	"api_tugas_minggu4/src/controllers/category_controllers"
	members_controller "api_tugas_minggu4/src/controllers/members_controllers"
	"api_tugas_minggu4/src/controllers/order_controllers"
	"api_tugas_minggu4/src/controllers/products_controllers"
	"api_tugas_minggu4/src/controllers/ratings_controllers"
	"api_tugas_minggu4/src/middleware"
	"fmt"
	"net/http"

	"github.com/goddtriffin/helmet"
)

func Router() {
	helmet := helmet.Default()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World ,Ini Adalah Halaman paling awal ya")
	})

	// http.Handle("/costumers", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(coscontrol.Costumers))))

	http.Handle("/ratings", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ratings_controllers.Data_ratings))))
	http.Handle("/rating/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(ratings_controllers.Data_rating))))
	http.Handle("/orders/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(order_controllers.Data_orders))))
	http.Handle("/order/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(order_controllers.Data_order))))
	http.Handle("/members", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.Data_members))))
	http.Handle("/member/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(members_controller.Data_member))))
	http.Handle("/products/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(products_controllers.Data_products))))
	http.Handle("/product/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(products_controllers.Data_product))))
	http.Handle("/categories", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(category_controllers.Data_categories))))
	http.Handle("/category/", helmet.Secure(middleware.XssMiddleware(http.HandlerFunc(category_controllers.Data_category))))

}
