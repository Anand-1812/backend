package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(client *mongo.Client) http.HandlerFunc {
	repo := NewMongoUserRepository(client)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		if req.Email == "" || req.Password == "" {
			http.Error(w, "email and password are required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		_, err := repo.GetByEmail(ctx, req.Email)
		if err == nil {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		}
		if !errors.Is(err, ErrUserNotFound) {
			http.Error(w, "failed to check user", http.StatusInternalServerError)
			return
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}

		user := &User{
			Email:        req.Email,
			PasswordHash: string(passwordHash),
		}

		if err := repo.Create(ctx, user); err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusCreated, user)
	}
}

func Login(client *mongo.Client) http.HandlerFunc {
	repo := NewMongoUserRepository(client)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req LoginRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		if req.Email == "" || req.Password == "" {
			http.Error(w, "email and password are required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		user, err := repo.GetByEmail(ctx, req.Email)
		if errors.Is(err, ErrUserNotFound) {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "failed to get user", http.StatusInternalServerError)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}

		token, err := createToken(user)
		if err != nil {
			http.Error(w, "failed to create token", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, TokenResponse{Token: token})
	}
}

func createToken(user *User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is required")
	}

	now := time.Now()
	claims := Claims{
		UserID: user.ID.Hex(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.Hex(),
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
