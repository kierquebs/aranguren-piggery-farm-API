package setup

import (
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/appointment"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/authenticate"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/classification"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/contents"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/qrcode"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/stocks"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/transaction"
	u "github.com/kierquebs/aranguren-piggery-farm-API/controller/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	//API Group
	api := app.Group("/API")

	//Stocks Group
	stock := api.Group("/stock", logger.New())
	stock.Post("/FindByQR", stocks.FindByQR)
	stock.Get("/FindByID/:id", stocks.FindByID)
	stock.Get("/ListAll", stocks.ListAll)
	stock.Post("/Create", stocks.Create)
	stock.Post("/GeneralExpectedWeight", stocks.GeneralExpectedWeight)

	//Classification Group
	class := api.Group("/classification", logger.New())
	class.Post("/Create", classification.Create)
	class.Get("/ListAll", classification.ListAll)

	//QR Group
	qr := api.Group("/qr", logger.New())
	qr.Post("/Generate", qrcode.Generate)

	//Transaction Group
	trn := api.Group("/transaction", logger.New())
	trn.Post("/Sell", transaction.Create)
	trn.Get("/IsSold/:id", transaction.IsSold)
	trn.Get("/Find", transaction.Find)
	trn.Get("/Find/:refID", transaction.FindByRefID)

	//User Group
	user := api.Group("/user", logger.New())
	user.Post("/Login", authenticate.Login)
	user.Post("/Token/:token", authenticate.TokenValidate)
	user.Post("/Create", u.Create)

	//Web Group
	web := api.Group("/web", logger.New())
	web.Get("/Contents", contents.ListAll)
	web.Post("/Update", contents.UpdateWebContent)
	web.Get("/Contacts", contents.ListContact)

	//Appointment Group
	apt := api.Group("/appointment", logger.New())
	apt.Post("/Create", appointment.Create)
	apt.Get("/List", appointment.ListAll)
	apt.Post("/Update", appointment.Update)
	apt.Post("/Delete/:id", appointment.Delete)
}
