package v1

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	models "github.com/PetarPeychev/test-signer-service/api/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type SignAnswersRequest struct {
	QuestionsAnswers []models.QuestionAnswer `json:"answers"`
}

type SignAnswersResponse struct {
	SignatureID int `json:"signatureId"`
}

func SignAnswers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "invalid userID", http.StatusBadRequest)
			return
		}

		_, claims, _ := jwtauth.FromContext(r.Context())
		jwtUserID, ok := claims["userID"]
		if !ok {
			http.Error(w, "Unauthorized: JWT doesn't contain a userID", http.StatusUnauthorized)
			return
		}

		if int(jwtUserID.(float64)) != id {
			http.Error(w, "Unauthorized: JWT doesn't match userID", http.StatusUnauthorized)
			return
		}

		var request SignAnswersRequest
		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "invalid request format", http.StatusBadRequest)
			return
		}

		signatureId, err := insertSignature(db, id, request.QuestionsAnswers)
		if err != nil {
			log.Println(err)
			http.Error(w, "error inserting signature", http.StatusInternalServerError)
			return
		}

		response := SignAnswersResponse{
			SignatureID: signatureId,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

func insertSignature(db *sql.DB, userID int, qa []models.QuestionAnswer) (id int, err error) {
	if err := ensureUserExists(db, userID); err != nil {
		return 0, err
	}

	qaJSON, err := json.Marshal(qa)
	if err != nil {
		return 0, err
	}

	var signatureID int
	err = db.QueryRow(
		"INSERT INTO signatures (user_id, questions_answers) VALUES ($1, $2) RETURNING id",
		userID,
		qaJSON,
	).Scan(&signatureID)
	if err != nil {
		return 0, err
	}

	return signatureID, nil
}

func ensureUserExists(db *sql.DB, userID int) error {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE id = $1", userID).Scan(&id)

	if err == sql.ErrNoRows {
		_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", userID)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
