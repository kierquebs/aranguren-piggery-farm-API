package setup

import (
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
	stock.Get("/FindByQR", stocks.FindByQR)
	stock.Post("/Create", stocks.Create)

	//QR Group
	qr := api.Group("/qr", logger.New())
	qr.Post("/Generate", qrcode.Generate)
}
