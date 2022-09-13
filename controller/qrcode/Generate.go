package qrcode

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
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

	d := utils.UploadQR(encoded + ".png")

	return c.JSON(string(d))

}
