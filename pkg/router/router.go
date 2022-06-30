package router

import (
	"fmt"
	"github.com/NubeIO/edge/controller"
	dbase "github.com/NubeIO/edge/database"
	"github.com/NubeIO/edge/pkg/config"
	dbhandler "github.com/NubeIO/edge/pkg/handler"
	"github.com/NubeIO/edge/pkg/logger"
	"github.com/spf13/viper"
	"io"

	"github.com/NubeIO/edge/service/apps/installer"
	"github.com/NubeIO/edge/service/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"time"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	// Set gin access logs
	fileLocation := fmt.Sprintf("%s/edge.access.log", config.Config.GetAbsDataDir())
	f, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		logger.Logger.Errorf("Failed to create access log file: %v", err)
	} else {
		gin.SetMode(viper.GetString("gin.loglevel"))
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders: []string{
			"X-FLOW-Key", "Authorization", "Content-Type", "Upgrade", "Origin",
			"Connection", "Accept-Encoding", "Accept-Language", "Host",
		},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		AllowAllOrigins:        true,
		AllowBrowserExtensions: true,
		MaxAge:                 12 * time.Hour,
	}))

	appDB := &dbase.DB{
		DB: db,
	}
	dbHandler := &dbhandler.Handler{
		DB: appDB,
	}
	dbhandler.Init(dbHandler)

	install := installer.New(&installer.Installer{
		DB: appDB,
	})

	api := controller.Controller{DB: appDB, Installer: install}
	identityKey := "uuid"

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "go-proxy-service",
		Key:           []byte(os.Getenv("JWTSECRET")),
		Timeout:       time.Hour * 1000,
		MaxRefresh:    time.Hour,
		IdentityKey:   identityKey,
		PayloadFunc:   auth.MapClaims,
		Authenticator: api.Login,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization",
		TimeFunc:    time.Now,
	})

	admin := r.Group("/api")

	r.POST("/api/users", api.AddUser)
	r.POST("/api/users/login", authMiddleware.LoginHandler)

	users := admin.Group("/users")
	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("/", api.GetUsers)
		users.GET("/:uuid", api.GetUser)
		users.PATCH("/:uuid", api.UpdateUser)
		users.DELETE("/:uuid", api.DeleteUser)
		users.DELETE("/drop", api.DropUsers)
	}

	store := admin.Group("/stores")
	{
		store.GET("/", api.GetAppStores)
		store.POST("/", api.CreateAppStore)
		store.GET("/:uuid", api.GetAppStore)
		store.PATCH("/:uuid", api.UpdateAppStore)
		store.DELETE("/:uuid", api.DeleteAppStore)
		store.DELETE("/drop", api.DropAppStores)
	}

	app := admin.Group("/apps")
	{
		app.GET("/", api.GetApps)
		app.POST("/", api.InstallApp)
		app.GET("/:uuid", api.GetApp)
		app.PATCH("/:uuid", api.UpdateApp)
		app.DELETE("/", api.UnInstallApp)
		app.DELETE("/drop", api.DropApps)
		// stats
		app.POST("/progress/install", api.GetInstallProgress)
		app.POST("/progress/uninstall", api.GetUnInstallProgress)
		app.POST("/stats", api.AppStats)
	}
	appControl := admin.Group("/apps/control")
	{
		appControl.POST("/", api.AppService)
		appControl.POST("/bulk", api.AppService)
	}

	device := admin.Group("/device")
	{
		device.GET("/", api.GetDeviceInfo)
		device.POST("/", api.AddDeviceInfo)
		device.PATCH("/", api.UpdateDeviceInfo)

	}

	system := admin.Group("/system")
	{
		system.GET("/ping", api.Ping)
		system.GET("/time", api.HostTime)
		system.GET("/product", api.GetProduct)
		system.POST("/scanner", api.RunScanner)
	}

	networking := admin.Group("/networking")
	{
		networking.GET("/networks", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)
		networking.GET("/update/schema", api.GetIpSchema)
		networking.POST("/update/dhcp", api.SetDHCP)
		networking.POST("/update/static", api.SetStaticIP)
	}

	files := admin.Group("/files")
	{
		files.GET("/read/*filePath", api.ReadDirs)
		files.POST("/download/*filePath", api.DownloadFile)
		files.DELETE("/delete/*filePath", api.DeleteFile)
		files.POST("/rename", api.RenameFile)
		files.POST("/move", api.MoveFile)
		files.POST("/upload", api.UploadFile)
	}

	dirs := admin.Group("/dirs")
	{
		dirs.DELETE("/delete/*filePath", api.DeleteDir)
		dirs.DELETE("/force/*filePath", api.DeleteDirForce)
		dirs.POST("/move", api.CopyDir)

	}
	zip := admin.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	return r
}
