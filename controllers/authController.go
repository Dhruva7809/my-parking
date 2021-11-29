package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/0xPiyush/my-parking/db"
	"github.com/0xPiyush/my-parking/db/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

const secretKey = "2k is not enough for this app bro"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body"})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: data["password"],
	}

	// check if user with same email exists
	db.DB.Where(&models.User{Email: user.Email}).First(&user)

	if user.Id != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User already exists"})
	}
	// create user
	db.DB.Create(&user)
	return c.JSON(fiber.Map{"message": "success"})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body"})
	}

	user := models.User{
		Email:    data["email"],
		Password: data["password"],
	}

	db.DB.Where(&user).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: &jwt.Time{time.Now().Add(time.Hour * 24)},
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error signing token"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "success"})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"message": "success"})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// return the user from the database
	user := models.User{}
	db.DB.First(&user, claims.Issuer)

	return c.JSON(user)
}

func Locations(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	// return all locations from the database
	locations := []models.Location{}
	db.DB.Find(&locations)

	return c.JSON(locations)
}

func Bookings(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	// return all Bookings from the database
	bookings := []models.Booking{}
	db.DB.Find(&bookings)

	return c.JSON(bookings)
}

func AddBooking(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error parsing body"})
	}

	booking := models.Booking{
		Id:       data["id"],
		Location: data["location"],
		ParkName: data["ParkingName"],
		Slot:     data["slot"],
		UserId:   data["userId"],
		Rate:     data["rate"],
		From:     data["from"],
		To:       data["to"],
	}

	// create booking
	db.DB.Create(&booking)
	return c.JSON(fiber.Map{"message": "success"})
}
