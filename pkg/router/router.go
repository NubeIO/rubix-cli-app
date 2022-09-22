package router

import (
	"fmt"
	"github.com/NubeIO/lib-rubix-installer/installer"
	"github.com/NubeIO/rubix-edge/controller"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/pkg/logger"
	"github.com/NubeIO/rubix-edge/service/apps"
	"github.com/NubeIO/rubix-edge/service/system"
	"github.com/NubeIO/rubix-registry-go/rubixregistry"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"time"
)

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		message := fmt.Sprintf("%s %s [%d]: %s", ctx.Request.Method, ctx.Request.URL, 404, "rubix-edge: api not found")
		ctx.JSON(http.StatusNotFound, controller.Message{Message: message})
	}
}

func Setup() *gin.Engine {
	engine := gin.New()
	// Set gin access logs
	if viper.GetBool("gin.log.store") {
		fileLocation := fmt.Sprintf("%s/edge.access.log", config.Config.GetAbsDataDir())
		f, err := os.OpenFile(fileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			logger.Logger.Errorf("Failed to create access log file: %v", err)
		} else {
			gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		}
	}
	gin.SetMode(viper.GetString("gin.log.level"))
	engine.NoRoute(NotFound())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.New(cors.Config{
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

	edgeSystem := system.New(&system.System{})
	edgeApp := apps.EdgeApp{App: installer.New(&installer.App{})}
	api := controller.Controller{EdgeApp: &edgeApp, RubixRegistry: rubixregistry.New(), System: edgeSystem, FileMode: 0755}
	engine.POST("/api/users/login", api.Login)
	publicSystemApi := engine.Group("/api/system")
	{
		publicSystemApi.GET("/ping", api.Ping)
		publicSystemApi.GET("/device", api.GetDeviceInfo)
		publicSystemApi.GET("/product", api.GetProduct)
		publicSystemApi.GET("/network_interfaces", api.GetNetworkInterfaces)
	}

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		// handleAuth = api.HandleAuth() // TODO add back in auth
	}

	apiRoutes := engine.Group("/api", handleAuth)
	apiProxyRoutes := engine.Group("/ff", handleAuth)
	apiProxyRoutes.Any("/*proxyPath", api.FFProxy) // FLOW-FRAMEWORK PROXY

	edgeApps := apiRoutes.Group("/apps")
	{
		edgeApps.GET("/", api.ListApps)
		edgeApps.GET("/status", api.ListAppsStatus)
		edgeApps.POST("/upload", api.UploadApp)
		edgeApps.POST("/service/upload", api.UploadServiceFile)
		edgeApps.POST("/service/install", api.InstallService)
		edgeApps.DELETE("/", api.UninstallApp)
	}

	appControl := apiRoutes.Group("/apps/control")
	{
		appControl.POST("/action", api.SystemCtlAction)
		appControl.POST("/status", api.SystemCtlStatus)
		appControl.POST("/action/mass", api.ServiceMassAction)
		appControl.POST("/status/mass", api.ServiceMassStatus)
	}

	appBackups := apiRoutes.Group("/backup")
	{
		appBackups.POST("/restore/full", api.RestoreBackup)
		appBackups.POST("/restore/app", api.RestoreAppBackup)
		appBackups.POST("/run/full", api.FullBackUp)
		appBackups.POST("/run/app", api.BackupApp)
		appBackups.GET("/list/full", api.ListFullBackups)
		appBackups.GET("/list/apps", api.ListAppsBackups)
		appBackups.GET("/list/app", api.ListAppBackups)
	}

	systemTime := apiRoutes.Group("/time")
	{
		systemTime.GET("/", api.SystemTime)
		systemTime.POST("/", api.SetSystemTime)
	}

	systemTimeZone := apiRoutes.Group("/timezone")
	{
		systemTimeZone.GET("/", api.GetHardwareTZ)
		systemTimeZone.POST("/", api.UpdateTimezone)
		systemTimeZone.GET("/list", api.GetTimeZoneList)
		systemTimeZone.POST("/config", api.GenerateTimeSyncConfig)
	}

	systemApi := apiRoutes.Group("/system")
	{
		systemApi.PATCH("/device", api.UpdateDeviceInfo)
		systemApi.POST("/scanner", api.RunScanner)
	}

	networking := apiRoutes.Group("/networking")
	{
		networking.GET("/", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)
	}

	networks := apiRoutes.Group("/networking/networks")
	{
		networks.POST("/restart", api.RestartNetworking)
		networks.POST("/up", api.InterfaceUp)
		networks.POST("/down", api.InterfaceDown)
	}

	networkAddress := apiRoutes.Group("/networking/interfaces")
	{
		networkAddress.POST("/exists", api.DHCPPortExists)
		networkAddress.POST("/auto", api.DHCPSetAsAuto)
		networkAddress.POST("/static", api.DHCPSetStaticIP)
	}

	networkFirewall := apiRoutes.Group("/networking/firewall")
	{
		networkFirewall.GET("/", api.UWFStatusList)
		networkFirewall.POST("/status", api.UWFStatus)
		networkFirewall.POST("/active", api.UWFActive)
		networkFirewall.POST("/enable", api.UWFEnable)
		networkFirewall.POST("/disable", api.UWFDisable)
		networkFirewall.POST("/port/open", api.UWFOpenPort)
		networkFirewall.POST("/port/close", api.UWFClosePort)
	}

	files := apiRoutes.Group("/files")
	{
		files.GET("/exists", api.FileExists)
		files.GET("/list", api.ListFiles) // /api/files/list?file=/data
		files.GET("/walk", api.WalkFile)
		files.GET("/read", api.ReadFile) // path=/data/flow-framework/config/config.yml
		files.POST("/create", api.CreateFile)
		files.POST("/write/string", api.WriteFile)
		files.POST("/write/json", api.WriteFileJson)
		files.POST("/write/yml", api.WriteFileYml)
		files.POST("/rename", api.RenameFile)
		files.POST("/copy", api.CopyFile)
		files.POST("/move", api.MoveFile)
		files.POST("/upload", api.UploadFile)
		files.POST("/download", api.DownloadFile)
		files.DELETE("/delete", api.DeleteFile)
		files.DELETE("/delete/all", api.DeleteAllFiles)
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.GET("/exists", api.DirExists)
		dirs.POST("/create", api.CreateDir)
		dirs.POST("/copy", api.CopyDir)
		dirs.DELETE("/delete", api.DeleteDir)
	}

	zip := apiRoutes.Group("/zip")
	{
		zip.POST("/unzip", api.Unzip)
		zip.POST("/zip", api.ZipDir)
	}

	user := apiRoutes.Group("/users")
	{
		user.PUT("", api.UpdateUser)
		user.GET("", api.GetUser)
	}

	token := apiRoutes.Group("/tokens")
	{
		token.GET("", api.GetTokens)
		token.POST("/generate", api.GenerateToken)
		token.PUT("/:uuid/block", api.BlockToken)
		token.PUT("/:uuid/regenerate", api.RegenerateToken)
		token.DELETE("/:uuid", api.DeleteToken)
	}

	streamLog := apiRoutes.Group("/logs")
	{
		streamLog.GET("", api.GetStreamLogs)
		streamLog.GET("/:uuid", api.GetStreamLog)
		streamLog.POST("", api.CreateStreamLog)
		streamLog.DELETE("/:uuid", api.DeleteStreamLog)
		streamLog.DELETE("", api.DeleteStreamLogs)
	}

	return engine
}
