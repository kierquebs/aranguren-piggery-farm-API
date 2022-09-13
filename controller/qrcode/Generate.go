package qrcode

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/controller/stocks"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"
)

func Generate(c *fiber.Ctx) error {

	genQR := new(model.GenerateQR)
	utils.BodyParser(c, genQR)

	data := string(genQR.ID) + "." + genQR.Added_Date
	hasher := md5.New()
	hasher.Write([]byte(data))

	encoded := hex.EncodeToString(hasher.Sum(nil))

	err := utils.GenerateQR(encoded, encoded)
	if err != nil {
		return err
	}

	result, err := utils.UploadQR(encoded + ".png")
	if err != nil {
		return err
	}

	err = stocks.UpdateQR(c, encoded)
	if err != nil {
		return err
	}

	return c.JSON(result)

}
