// auth.go
package handler

import (
	"Project2/app/configs"
	"Project2/app/models"
	"Project2/app/utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var authUserCollection *mongo.Collection = configs.GetCollection(configs.DB, "authUserCollection")
var RoleCollection *mongo.Collection = configs.GetCollection(configs.DB, "roleCollection")
var jwtKey = []byte("my_secret_key")

// Login
func Login(c *fiber.Ctx) error {
	var authUser models.AuthUser
	if err := c.BodyParser(&authUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingAuthUser models.AuthUser
	err := authUserCollection.FindOne(ctx, bson.M{"email": authUser.Email}).Decode(&existingAuthUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "user not found"})
	}

	if !utils.CompareHashPassword(authUser.Password, existingAuthUser.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid password"})
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		// Assuming a user role here, you might need to adjust this
		StandardClaims: jwt.StandardClaims{
			Subject:   existingAuthUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
		SameSite: "Lax",
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "user logged in"})
}

// SignUp
func SignUp(c *fiber.Ctx) error {
	var authUser models.AuthUser
	if err := c.BodyParser(&authUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingAuthUser models.AuthUser
	err := authUserCollection.FindOne(ctx, bson.M{"email": authUser.Email}).Decode(&existingAuthUser)
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "user is already registered"})
	}

	var errHash error
	authUser.Password, errHash = utils.GenerateHashPassword(authUser.Password)
	if errHash != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": errHash.Error()})
	}

	// Check if role exists
	var role models.Role
	err = RoleCollection.FindOne(ctx, bson.M{"_id": authUser.RoleID}).Decode(&role)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "role not found"})
	}

	authUser.Role = role

	_, err = authUserCollection.InsertOne(ctx, authUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "user created"})
}

// Home
func Home(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if claims.Role != "user" && claims.Role != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "home page", "role": claims.Role})
}

// Premium
func Premium(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	if claims.Role != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "premium page", "role": "admin"})
}

// Logout
func Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Status(http.StatusOK).JSON(fiber.Map{"success": "logout"})
}
