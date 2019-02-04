package api

import (
	"github.com/ClubCedille/pixicoreAPI/pkg/config"
	"github.com/gin-gonic/gin"
)

//// *************************** Controllers ****************************
// InitController Generate new controller for PXE boot
func InitController(confFile *config.ConfigFile) Controller {
	ctrl := Controller{currentConfig: confFile}
	return ctrl
}

// Cors cors for the api
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetRouter(controller Controller) *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("v1")

	{
		v1.GET("/", controller.Getlocal)
		v1.GET("/boot/:macAddress", controller.BootServer)
		v1.GET("/single/:macAddress", controller.InstallServer)
		v1.GET("/all/", controller.InstallAll)
		v1.GET("/servers", controller.GetServers)
		v1.GET("/update/:macAddress", controller.UpdateTest)
	}
	return r
}
