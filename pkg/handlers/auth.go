package handlers

import (
	"go-chat-app/pkg/database"
	"go-chat-app/pkg/models"
	"go-chat-app/pkg/utils"
	"log"
	"net/http"
	"encoding/json"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return 
	}

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return 
	}

	if req.Username == "" || req.Email == "" || req.Password == ""{
		http.Error(w, `{"error":"username, email and password are required"}`, http.StatusConflict)
		return 
	}

	existingUser, _ := database.GetUserByUserName(req.Username)
	if existingUser != nil{
		http.Error(w, `{"error":"user already exists"}`, http.StatusConflict)
		return 
	}

	existingEmail, _ := database.GetUserByEmail(req.Email)
	if existingEmail != nil{
		http.Error(w, `{"error":"user already exists"}`, http.StatusConflict)
		return 
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil{
		log.Printf("Error hashing password: %v\n", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return 
	}

	user, err := database.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		log.Printf("error creating user: %v\n", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return 
	}

	token, err := utils.GenerateToken(user.ID, user.UserName, user.Email)
	if err != nil {
		log.Printf("Error generating token: %v\n", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := models.AuthResponse{
		Token : token,
		User : &models.User{
			ID : user.ID,
			UserName: user.UserName,
			Email : user.Email,
			Password : hashedPassword,
			CreatedAt : time.Now(),
		},
	}
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, `{}"error":"method not allowed"}`, http.StatusMethodNotAloowed)
		return 
	}

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return 
	}

	if req.Username == "" || req.Password == ""{
		http.Error(w, `{"error":"username and password are required}`, http.StatusBadRequest)
		return 
	}

	user, err := database.GetUserByUserName(req.Username)
	if err != nil{
		http.Error(w, `{"error":"user could not be found"}`, http.StatusUnauthorized)
		return 
	}

	if !utils.VerifyPassword(user.Password, req.Password){
		http.Error(w, `{"error":"invalid password"}`, http.StatusUnauthorized)
		return 
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Password)
	if err != nil {
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	response := models.AuthResponse{
		Token : token,
		User : &models.User{
			ID : user.ID,
			UserName: user.UserName,
			Email : user.Email,
		},
	}
	json.NewEncoder(w).Encode(response)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
