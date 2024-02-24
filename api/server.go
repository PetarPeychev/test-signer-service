package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	v1 "github.com/PetarPeychev/test-signer-service/api/handlers/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

type Server struct {
	router  *chi.Mux
	db      *sql.DB
	address string
}

func NewServer(config Config) (*Server, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBHost,
		config.DBPort,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for {
		err = db.Ping()
		if err != nil {
			log.Println(err)
			log.Println("Waiting for database...")
			time.Sleep(3 * time.Second)
		} else {
			log.Println("Connected to database")
			break
		}
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Route("/signatures", func(r chi.Router) {
					// tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)
					// log.Println("JWT secret:", config.JWTSecret)
					// r.With(
					// 	jwtauth.Verifier(tokenAuth),
					// 	jwtauth.Authenticator(tokenAuth),
					// ).Post("/", v1.SignAnswers(db))
					r.Post("/", v1.SignAnswers(db))
					r.Get("/{signatureID}", v1.VerifySignature(db))
				})
			})
		})
	})

	return &Server{router: r, db: db, address: config.Address}, nil
}

func (server *Server) ListenAndServe() error {
	log.Printf("Listening on %s...\n", server.address)
	err := http.ListenAndServe(server.address, server.router)
	if err != nil {
		log.Println(err)
	}
	return err
}
