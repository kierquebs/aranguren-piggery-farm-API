package authenticate

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
)

type tokenDetails struct {
	Username string
	Token    string
}

func TokenValidate(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	un := tokenDetails{}
	token := c.Params("token")
	sqlStatement := `SELECT username,token FROM public.t_user WHERE token = $1;`
	row := database.CCDB.QueryRow(sqlStatement, token)
	switch err := row.Scan(&un.Username, &un.Token); err {
	case sql.ErrNoRows:
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Invalid token", "data": nil})
	case nil:
		return nil
	default:
		panic(err)
	}
}
