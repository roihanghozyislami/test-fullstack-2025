package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)


type User struct {
	Realname string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"` 
}


func sha1Hash(input string) string {
	hash := sha1.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	app := fiber.New()


	app.Post("/login", func(c *fiber.Ctx) error {
	
		var body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).SendString("Input tidak valid")
		}

		
		key := "login_" + body.Username
		val, err := Rdb.Get(Ctx, key).Result()
		if err != nil {
			return c.Status(401).SendString("Username tidak ditemukan")
		}

	
		var user User
		if err := json.Unmarshal([]byte(val), &user); err != nil {
			return c.Status(500).SendString("Gagal membaca data pengguna")
		}

		
		if user.Password != sha1Hash(body.Password) {
			return c.Status(401).SendString("Password salah")
		}

	
		return c.JSON(fiber.Map{
			"message":  "Login berhasil",
			"realname": user.Realname,
			"email":    user.Email,
		})
	})


	log.Fatal(app.Listen(":3000"))
}
