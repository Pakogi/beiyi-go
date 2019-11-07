package conf

import (
	"os"

	"beiyi/model"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()

	model.Database(os.Getenv("DATABASE_DSN"))
}
