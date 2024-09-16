package main

import (
	"api-with-go-postgres/database"
	"api-with-go-postgres/tenders"
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/api/ping", pingHandler)                                  //Убедиться, что сервер готов обрабатывать запросы.
	http.HandleFunc("/api/tenders", tenders.GetAllTenders(db))                 //Возвращает список тендеров с возможностью фильтрации по типу услуг.
	http.HandleFunc("/api/tenders/new", tenders.CreateNewTender(db))           //Создает новый тендер с заданными параметрами.
	http.HandleFunc("/api/tenders/my", tenders.GetOrganizationsByUsername(db)) //Возвращает список тендеров текущего пользователя.
	http.HandleFunc("/api/tenders/", tenders.ChangeExistingTender(db))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error to start HTTP SERVER:", err)
	}
}
