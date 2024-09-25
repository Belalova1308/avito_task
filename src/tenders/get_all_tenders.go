package tenders

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mymodule/src/models"
	"net/http"
)

func GetAllTenders(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintln(w, "Method Not Allowed")
			return
		}
		tenderType := r.URL.Query().Get("serviceType")

		query := "SELECT id, name, description, service_type, status, organization_id, creator_id, version, created_at, updated_at FROM tender"
		var args []interface{}

		if tenderType != "" {
			query += " WHERE service_type = $1"
			args = append(args, tenderType)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			fmt.Println("Error retrieving tenders from database:", err)
			http.Error(w, "Error retrieving tenders", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tenders []models.Tender
		for rows.Next() {
			var tender models.Tender
			if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.CreatorID, &tender.Version, &tender.CreatedAt, &tender.UpdatedAt); err != nil {
				fmt.Println("Error scanning tender:", err)
				http.Error(w, "Error scanning tender", http.StatusInternalServerError)
				return
			}

			var creatorUsername string
			err = db.QueryRow("SELECT username FROM employee WHERE id = $1", tender.CreatorID).Scan(&creatorUsername)
			if err != nil {
				fmt.Println("Error fetching creator username:", err)
				http.Error(w, "Error fetching creator's username", http.StatusInternalServerError)
				return
			}

			tender.CreatorUsername = creatorUsername
			tenders = append(tenders, tender)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error reading rows", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(tenders); err != nil {
			http.Error(w, "Error encoding tenders", http.StatusInternalServerError)
			return
		}
	}
}
