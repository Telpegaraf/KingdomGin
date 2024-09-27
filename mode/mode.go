package mode

import "github.com/gin-gonic/gin"

const (
	Dev     = "dev"
	Prod    = "prod"
	TestDev = "testdev"
)

var mode = Dev

func Set(newMode string) {
	mode = newMode
	updateGinMode()
}

func Get() string {
	return mode
}

func IsDev() bool {
	return Get() == Dev || Get() == TestDev
}

func updateGinMode() {
	switch Get() {
	case Dev:
		gin.SetMode(gin.DebugMode)
	case TestDev:
		gin.SetMode(gin.TestMode)
	case Prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		panic("Unknown Mode")
	}
}
