package authenticate

import (
	"database/sql"
	"log"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils/password"
)

type userDetails struct {
	ID          int
	First_Name  string
	Last_Name   string
	Middle_Name string
	Password    string
}

func getUserByUsername(u string) (userDetails, error) {
	un := userDetails{}

	sqlStatement := `SELECT id, first_name, middle_name, last_name, password FROM public.t_user WHERE username = $1;`
	row := database.CCDB.QueryRow(sqlStatement, u)
	switch err := row.Scan(&un.ID, &un.First_Name, &un.Middle_Name, &un.Last_Name, &un.Password); err {
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

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	lm := model.LoginUserModel{}
	if err := c.BodyParser(&lm); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	user, err := getUserByUsername(lm.Username)
	if err != nil {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Username not exist!", "data": nil})
	}

	if validatePassword(user.Password, lm.Password) {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Invalid password!", "data": nil})
	}

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte(config.Config("SECRET")))
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	user.Password = ""

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Succesfully LogIn ", "data": user})
}

// validatePassword returns True if valid else returns false
func validatePassword(hashedPass string, pass string) bool {
	err1 := password.Compare(hashedPass, pass)
	if err1 != nil {
		return false
	}
	return true
}
