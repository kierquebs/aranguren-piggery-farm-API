package authenticate

import (
	"database/sql"
	"fmt"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
)

type userDetails struct {
	ID          int
	First_Name  string
	Last_Name   string
	Middle_Name string
}

func getUserByUsername(u string) (userDetails, error) {
	un := userDetails{}

	sqlStatement := `SELECT id, first_name, middle_name, last_name FROM public.t_user WHERE username = $1;`
	row := database.CCDB.QueryRow(sqlStatement, u)
	switch err := row.Scan(&un.ID, &un.First_Name, &un.Middle_Name, &un.Last_Name); err {
	case sql.ErrNoRows:
		return un, err
	case nil:
		return un, nil
	default:
		panic(err)
	}
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Login(c *fiber.Ctx) error {

	uName := c.Params("username")
	fmt.Println(uName)
	user, err := getUserByUsername(uName)
	if err != nil {
		return c.JSON(fiber.Map{"responseCode": 200, "message": err.Error(), "data": nil})
	}

	// if valid(identity) {
	// 	email, err = getUserByEmail(identity)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	// 	}
	// } else {
	// 	user, err = getUserByUsername(identity)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	// 	}
	// }

	// if email == nil && user == nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	// }

	// if !CheckPasswordHash(pass, ud.Password) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	// }

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte(config.Config("SECRET")))
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": user})
}
