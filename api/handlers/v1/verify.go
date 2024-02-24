package v1

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type VerifySignatureResponse struct {
	UserID      int    `json:"userId"`
	SignatureID int    `json:"signatureId"`
	Timestamp   string `json:"timestamp"`
}

func VerifySignature(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		signatureID := chi.URLParam(r, "signatureID")

		uid, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "invalid userID", http.StatusBadRequest)
			return
		}

		sid, err := strconv.Atoi(signatureID)
		if err != nil {
			http.Error(w, "invalid signatureID", http.StatusBadRequest)
			return
		}

		response, err := verifySignature(db, uid, sid)
		if err != nil {
			http.Error(w, "signature not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "error encoding response", http.StatusInternalServerError)
			return
		}
	}
}

func verifySignature(db *sql.DB, userID, signatureID int) (*VerifySignatureResponse, error) {
	var response VerifySignatureResponse
	var timestamp time.Time

	err := db.QueryRow(
		"SELECT user_id, id, timestamp FROM signatures WHERE user_id = $1 AND id = $2",
		userID,
		signatureID,
	).Scan(&response.UserID, &response.SignatureID, &timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("signature not found")
		}
		return nil, err
	}

	response.UserID = userID
	response.SignatureID = signatureID
	response.Timestamp = timestamp.Format(time.RFC3339)

	return &response, nil
}
