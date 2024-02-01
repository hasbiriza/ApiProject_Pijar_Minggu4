// package main

// import (
// 	"api_tugas_minggu4/src/config"
// 	"api_tugas_minggu4/src/helper"
// 	"api_tugas_minggu4/src/routes"
// 	"fmt"
// 	"net/http"

// 	"github.com/subosito/gotenv"
// )

// func main() {
// if err := gotenv.Load(); err != nil {
// 	fmt.Println("Failed to load .env file")
// 	return
// }

// 	config.InitDB()
// 	helper.Migration()
// 	defer config.DB.Close()
// 	routes.Router()

// if err := http.ListenAndServe(":8080", nil); err != nil {
// 	fmt.Printf("Failed to start the server: %v\n", err)
// }

// 	fmt.Print("Server running at http://localhost:8080\n")
// }

package main

import (
	"api_tugas_minggu4/src/config"
	"api_tugas_minggu4/src/helper"
	"api_tugas_minggu4/src/routes"
	"fmt"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	// gotenv.Load()
	if err := gotenv.Load(); err != nil {
		fmt.Println("Failed to load .env file")
		return
	}
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Router()
	fmt.Print("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start the server: %v\n", err)
	}
}
