package routes

import (
	"invoiceApi/database"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)
func init() {
	database.DBconn()

	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	})
	log.SetReportCaller(true)
	log.Warn("Init Database Connection")

}

func SetupRouter(app *fiber.App) {
		UserRouter(app);
		AuthRouter(app);
}
