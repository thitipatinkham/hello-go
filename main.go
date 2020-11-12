package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	ctrl "github.com/thitipat-artisan/hello-go/controller"
)

// LoginInput ...
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const key = "hello"
const temp = "be7afbb85d5ee3b4bff8fcbbc477e31dd295238624edec95790e636b82ae45bf"

func auth(c *fiber.Ctx) error {
	if c.Path() == "/login" && c.Method() == "POST" {
		return c.Next()
	}
	token := c.Request().Header.Peek("Authorization")
	if string(token) == string(temp) {
		return c.Next()
	}
	return c.SendStatus(401)
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
func login(c *fiber.Ctx) error {
	var input LoginInput
	json.Unmarshal(c.Body(), &input)
	fmt.Printf("%+v", input)

	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(fmt.Sprintf("%s", input.Username)))
	token := hex.EncodeToString(mac.Sum(nil))
	fmt.Println(token)
	return c.SendString(fmt.Sprintf("token:%s", token))
}

func getUser(c *fiber.Ctx) error {
	return c.JSON(ctrl.GetUsers())
}

func main() {
	app := fiber.New()

	// app.Use(auth)

	app.Get("/", hello)
	app.Post("/login", login)
	app.Get("/users", getUser)

	app.Listen(":8080")
}
