package tenders

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mymodule/src/models"
	"net/http"
)

func CreateNewTender(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		var tender models.Tender

		err := json.NewDecoder(r.Body).Decode(&tender)
		if err != nil {
			fmt.Println("Error decoding tender into database:", err)
			http.Error(w, "Error decoding the request", http.StatusBadRequest)
			return
		}

		var creatorID string
		err = db.QueryRow("SELECT id FROM employee WHERE username = $1", tender.CreatorUsername).Scan(&creatorID)
		if err != nil {
			http.Error(w, "Creator not found", http.StatusBadRequest)
			return
		}

		var newTenderID string
		query := `
            INSERT INTO tender (name, description, service_type, status, organization_id, creator_id, version) 
            VALUES ($1, $2, $3, 'CREATED', $4, $5, 1) 
            RETURNING id`
		err = db.QueryRow(query, tender.Name, tender.Description, tender.ServiceType, tender.OrganizationID, creatorID).Scan(&newTenderID)
		if err != nil {
			fmt.Println("Error inserting tender into database:", err)
			http.Error(w, "Error inserting new tender into database", http.StatusInternalServerError)
			return
		}

		var newTender models.Tender
		fetchQuery := `
            SELECT id, name, description, service_type, status, organization_id, creator_id, version, created_at, updated_at 
            FROM tender WHERE id = $1`
		err = db.QueryRow(fetchQuery, newTenderID).Scan(&newTender.ID, &newTender.Name, &newTender.Description, &newTender.ServiceType, &newTender.Status, &newTender.OrganizationID, &newTender.CreatorID, &newTender.Version, &newTender.CreatedAt, &newTender.UpdatedAt)
		if err != nil {
			http.Error(w, "Error fetching newly created tender", http.StatusInternalServerError)
			return
		}
		var creatorUsername string
		err = db.QueryRow("SELECT username FROM employee WHERE id = $1", newTender.CreatorID).Scan(&creatorUsername)
		if err != nil {
			fmt.Println("Error fetching creator's username:", err)
			http.Error(w, "Error fetching creator's username", http.StatusInternalServerError)
			return
		}
		newTender.CreatorUsername = creatorUsername
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newTender)
	}
}
