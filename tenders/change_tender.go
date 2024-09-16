package tenders

import (
	"api-with-go-postgres/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ChangeExistingTender(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) != 5 || pathParts[4] != "edit" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		tenderID := pathParts[3]
		if tenderID == "" {
			http.Error(w, "Missing tender ID", http.StatusBadRequest)
			return
		}
		var tender models.Tender
		err := json.NewDecoder(r.Body).Decode(&tender)
		if err != nil {
			fmt.Println("Error decoding tender into database:", err)
			http.Error(w, "Error decoding the request", http.StatusBadRequest)
			return
		}

		query := `
			UPDATE tender
			SET name=$2, description=$3, service_type=$4, updated_at=NOW()`

		var params []interface{}
		params = append(params, tenderID, tender.Name, tender.Description, tender.ServiceType)

		if tender.OrganizationID != "" {
			query += ", organization_id=$5"
			params = append(params, tender.OrganizationID)
		}
		if tender.CreatorID != "" {
			query += ", creator_id=$6"
			params = append(params, tender.CreatorID)
		}

		query += ", status='PUBLISHED', version=version+1 WHERE id=$1"

		_, err = db.Exec(query, params...)
		if err != nil {
			fmt.Println("Error updating tender in database:", err)
			http.Error(w, "Error updating tender", http.StatusInternalServerError)
			return
		}

		query = `SELECT id, name, description, service_type, status, organization_id, creator_id, version, created_at, updated_at 
            FROM tender WHERE id = $1`
		err = db.QueryRow(query, tenderID).Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.CreatorID, &tender.Version, &tender.CreatedAt, &tender.UpdatedAt)
		if err != nil {
			http.Error(w, "Error fetching newly created tender", http.StatusInternalServerError)
			return
		}
		var creatorUsername string
		err = db.QueryRow("SELECT username FROM employee WHERE id = $1", tender.CreatorID).Scan(&creatorUsername)
		if err != nil {
			fmt.Println("Error fetching creator's username:", err)
			http.Error(w, "Error fetching creator's username", http.StatusInternalServerError)
			return
		}
		tender.CreatorUsername = creatorUsername
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tender)
	}
}
