package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var UseGuestbook bool
var UseAttendance bool
var AdminPassword string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Cannot read .env file")
		panic(err.Error())
	}
	UseGuestbook = os.Getenv("USE_GUESTBOOK") == "true"
	UseAttendance = os.Getenv("USE_ATTENDANCE") == "true"
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
}
