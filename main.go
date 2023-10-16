// main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"vibank/config"
	"vibank/models"
	"vibank/utils"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Define a secret key for signing JWT tokens
var (
	app *config.Database
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app, err = config.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck)
	r.HandleFunc("/dashboard", DisplayDashboard)
	r.HandleFunc("/register", Register)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/logout", Logout)
	http.ListenAndServe(":"+port, r)
}

func DisplayDashboard(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<!doctype html>
	<html lang="en">
	
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Bootstrap demo</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"
			integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Josefin+Sans&display=swap" rel="stylesheet">
		<style>
			body {
				font-family: 'Josefin Sans', sans-serif;
			}
		</style>
	</head>
	
	<body>
	   <div class="container">
		<div class="row">
			<div class="col">
				<h2 class="h2 text-center mt-5">This is V!Bank Web Application Dashboard</h2>
			</div>
		</div>
	   </div>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"
			integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL"
			crossorigin="anonymous"></script>
	</body>
	
	</html>`))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("application is running . . . "))
}

// Define the RegisterUser handler function
func Register(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to extract user data
	var newUser *models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the user already exists
	existingUser, err := app.GetUserByEmail(newUser.Email)
	if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	newUser.ID, _ = uuid.NewV1()
	// Create the user in the database
	if err := app.CreateUser(newUser); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Define the LoginUser handler function
func Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to extract login credentials
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginInfo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the user by email
	user, err := app.GetUserByEmail(loginInfo.Email)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Check if the user exists and the password matches
	if user == nil || user.Password != loginInfo.Password {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
	}

	// Generate a JWT token
	tokenString, err := utils.GenerateJWTToken(claims)
	if err != nil {
		http.Error(w, "Error creating JWT token", http.StatusInternalServerError)
		return
	}

	// Save the JWT token in a cookie
	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Minute * 60),
	}
	http.SetCookie(w, &cookie)

	// Return a success response with the JWT token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "Login successful",
		Token:   tokenString,
	})
	// Redirect the user to a logout confirmation page or any other desired action
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Define the LogoutUser handler function
func Logout(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie with an empty value and an expiration time in the past
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour), // Set expiration time in the past to delete the cookie
	}
	http.SetCookie(w, &cookie)

	// Redirect the user to a logout confirmation page or any other desired action
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
