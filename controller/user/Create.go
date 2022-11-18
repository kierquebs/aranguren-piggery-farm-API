package user

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils/password"
)

func Create(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	usn := model.CreateUserModel{}
	if err := c.BodyParser(&usn); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	if !checkExisitingUser(usn.Username) {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Username already exist"})
	}

	hash, err1 := password.Encrypt(usn.Password)
	if err1 != nil {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Unable to Encrypt password"})
	}

	usn.Password = string(hash)

	_, err := database.CCDB.Exec(`INSERT INTO public.t_user(
		username,
		password,
		first_name,
		middle_name,
		last_name)
		VALUES ($1,$2,$3,$4,$5);`,
		usn.Username,
		usn.Password,
		usn.First_Name,
		usn.Middle_Name,
		usn.Last_Name,
	)

	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "New user successfully added."})
}

// checkExistingUser return True if no rows and False if have rows
func checkExisitingUser(username string) bool {

	var id int

	sqlStatement := `SELECT id FROM public.t_user WHERE username = $1;`
	row := database.CCDB.QueryRow(sqlStatement, username)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return true
	case nil:
		return false
	default:
		panic(err)
	}

}
