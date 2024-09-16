package tenders

import (
	"api-with-go-postgres/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOrganizationsByUsername(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		if username == "" {
			http.Error(w, "Username parameter is missing", http.StatusBadRequest)
			return
		}

		query := `
	SELECT o.id, o.name, o.description, o.type, o.created_at, o.updated_at
	FROM organization o
	JOIN organization_responsible org_res ON o.id = org_res.organization_id
	JOIN employee e ON org_res.user_id = e.id
	WHERE e.username = $1
`
		rows, err := db.Query(query, username)
		if err != nil {
			fmt.Println("Error selecting organizations from database:", err)
			http.Error(w, "Error retrieving organizations from database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var organizations []models.Organization
		for rows.Next() {
			var org models.Organization
			if err := rows.Scan(&org.ID, &org.Name, &org.Description, &org.Type, &org.CreatedAt, &org.UpdatedAt); err != nil {
				fmt.Println("Error scanning organization from database:", err)
				http.Error(w, "Error scanning organization", http.StatusInternalServerError)
				return
			}
			organizations = append(organizations, org)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error reading rows in database", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(organizations); err != nil {
			http.Error(w, "Error encoding organizations", http.StatusInternalServerError)
			return
		}
	}
}
