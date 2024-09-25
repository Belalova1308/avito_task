package main

import (
	"fmt"
	"log"
	"mymodule/src/database"
	"mymodule/src/tenders"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Not Allowed")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func main() {

	db := database.Connect()
	defer db.Close()

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		log.Fatal("SERVER_ADDRESS environment variable is required")
	}

	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/tenders", tenders.GetAllTenders(db))                 //Возвращает список тендеров с возможностью фильтрации по типу услуг.
	http.HandleFunc("/api/tenders/new", tenders.CreateNewTender(db))           //Создает новый тендер с заданными параметрами.
	http.HandleFunc("/api/tenders/my", tenders.GetOrganizationsByUsername(db)) //Возвращает список тендеров текущего пользователя.
	http.HandleFunc("/api/tenders/", tenders.ChangeExistingTender(db))

	fmt.Printf("Starting server at %s\n", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
