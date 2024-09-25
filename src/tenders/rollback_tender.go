package tenders

import (
	"database/sql"
	"net/http"
	"strings"
)

func RollbackTender(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) != 6 || pathParts[4] != "rollback" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		tenderID := pathParts[3]
		if tenderID == "" {
			http.Error(w, "Missing tender ID", http.StatusBadRequest)
			return
		}
		tenderVersion := pathParts[5]
		if tenderVersion == "" {
			http.Error(w, "Missing tender version", http.StatusBadRequest)
			return
		}

	}
}
