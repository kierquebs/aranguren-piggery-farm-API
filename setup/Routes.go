package setup

import (
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/classification"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/qrcode"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/stocks"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	//API Group
	api := app.Group("/API")

	//Stocks Group
	stock := api.Group("/stock", logger.New())
	stock.Post("/FindByQR", stocks.FindByQR)
	stock.Get("/ListAll", stocks.ListAll)
	stock.Post("/Create", stocks.Create)

	//Classification Group
	class := api.Group("/classification", logger.New())
	class.Post("/Create", classification.Create)
	class.Get("/ListAll", classification.ListAll)

	//QR Group
	qr := api.Group("/qr", logger.New())
	qr.Post("/Generate", qrcode.Generate)
}
