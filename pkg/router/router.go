package router

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
	"gthub.com/NubeIO/rubix-cli-app/controller"
	dbase "gthub.com/NubeIO/rubix-cli-app/database"
	dbhandler "gthub.com/NubeIO/rubix-cli-app/pkg/handler"
	"gthub.com/NubeIO/rubix-cli-app/pkg/logger"
	"gthub.com/NubeIO/rubix-cli-app/service/auth"
	"io"
	"os"
	"time"
)

func initWs() *melody.Melody {
	return melody.New()
}

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()
	var ws = initWs()
	// Write gin access log to file
	f, err := os.OpenFile("rubix.access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("Failed to create access log file: %v", err)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Set default middlewares
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

	api := controller.Controller{DB: appDB, WS: ws}
	identityKey := "uuid"

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
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

	//web socket route
	r.GET("/ws", func(c *gin.Context) {
		err := ws.HandleRequest(c.Writer, c.Request)
		fmt.Println(err)
		//if err != nil {
		//	return
		//}
	})

	ws.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		ws.Broadcast(msg)
	})

	admin := r.Group("/api")

	r.POST("/api/users", api.AddUser)
	r.POST("/api/users/login", authMiddleware.LoginHandler)

	users := admin.Group("/users")
	users.Use(authMiddleware.MiddlewareFunc())
	{
		//users.GET("/schema", api.UsersSchema)
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

	system := admin.Group("/system")
	{
		system.GET("/product", api.GetProduct)
	}
	return r
}
