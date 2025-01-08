package logic

import (
	"fmt"
	"time"
	"warehouse/models"
	"warehouse/utils"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthLogic struct {
	JWTSecret string
	Config    utils.Config
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Register handles user registration by inserting the user into the database after hashing the password.
func (al *AuthLogic) Register(user *models.User) error {
	if err := utils.EnsureDBConnection(al.Config); err != nil {
		return fmt.Errorf("failed to ensure database connection: %w", err)
	}

	// Hash password before saving
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashPass)
	// Insert user into database
	result, err := utils.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}

	// Get the last inserted user ID
	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	user.ID = int(userID)

	return nil
}

// Authenticate checks user credentials and returns a JWT token if valid.
func (al *AuthLogic) Authenticate(username, password string) (string, models.User, error) {
	if err := utils.EnsureDBConnection(al.Config); err != nil {
		return "", models.User{}, fmt.Errorf("failed to ensure database connection: %w", err)
	}

	// Query the user from the database
	user := models.User{}
	err := utils.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username).Scan(
		&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return "", models.User{}, fmt.Errorf("user not found")
	}

	// Compare the provided password with the stored hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", models.User{}, fmt.Errorf("invalid password")
	}

	// Create JWT Token
	expiredTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	// Sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(al.JWTSecret))
	if err != nil {
		return "", models.User{}, fmt.Errorf("invalid token: " + err.Error())
	}

	// Return the token and user info
	return tokenString, user, nil
}
