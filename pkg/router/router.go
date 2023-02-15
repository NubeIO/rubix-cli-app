package router

import (
	"fmt"
	"github.com/NubeIO/lib-systemctl-go/systemctl"
	"github.com/NubeIO/rubix-edge/controller"
	"github.com/NubeIO/rubix-edge/model"
	"github.com/NubeIO/rubix-edge/pkg/config"
	"github.com/NubeIO/rubix-edge/pkg/logger"
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
		ctx.JSON(http.StatusNotFound, model.Message{Message: message})
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

	api := controller.Controller{
		SystemCtl:     systemctl.New(false, 30),
		RubixRegistry: rubixregistry.New(),
		System:        system.New(&system.System{}),
		FileMode:      0755,
	}
	engine.POST("/api/users/login", api.Login)
	publicSystemApi := engine.Group("/api/system")
	{
		publicSystemApi.GET("/ping", api.Ping)
		publicSystemApi.GET("/device/public", api.GetDeviceInfo)
	}

	handleAuth := func(c *gin.Context) { c.Next() }
	if config.Config.Auth() {
		handleAuth = api.HandleAuth()
	}

	apiRoutes := engine.Group("/api", handleAuth)
	apiProxyRoutes := engine.Group("/ff", handleAuth)
	apiProxyRoutes.Any("/*proxyPath", api.FFProxy) // FLOW-FRAMEWORK PROXY
	apiProxyWiresRoutes := engine.Group("/wires", handleAuth)
	apiProxyWiresRoutes.Any("/*proxyPath", api.WiresProxy) // EDGE-WIRES PROXY
	apiProxyChirpRoutes := engine.Group("/chirp", handleAuth)
	apiProxyChirpRoutes.Any("/*proxyPath", api.ChirpProxy) // CHIRP-STACK PROXY

	appControl := apiRoutes.Group("/systemctl")
	{
		appControl.POST("/enable", api.SystemCtlEnable)
		appControl.POST("/disable", api.SystemCtlDisable)
		appControl.GET("/show", api.SystemCtlShow)
		appControl.POST("/start", api.SystemCtlStart)
		appControl.GET("/status", api.SystemCtlStatus)
		appControl.POST("/stop", api.SystemCtlStop)
		appControl.POST("/reset-failed", api.SystemCtlResetFailed)
		appControl.POST("/daemon-reload", api.SystemCtlDaemonReload)
		appControl.POST("/restart", api.SystemCtlRestart)
		appControl.POST("/mask", api.SystemCtlMask)
		appControl.POST("/unmask", api.SystemCtlUnmask)
		appControl.GET("/state", api.SystemCtlState)
		appControl.GET("/is-enabled", api.SystemCtlIsEnabled)
		appControl.GET("/is-active", api.SystemCtlIsActive)
		appControl.GET("/is-running", api.SystemCtlIsRunning)
		appControl.GET("/is-failed", api.SystemCtlIsFailed)
		appControl.GET("/is-installed", api.SystemCtlIsInstalled)
	}

	syscallControl := apiRoutes.Group("/syscall")
	{
		syscallControl.POST("/unlink", api.SyscallUnlink)
		syscallControl.POST("/link", api.SyscallLink)
	}

	systemTime := apiRoutes.Group("/time")
	{
		systemTime.GET("", api.SystemTime)
		systemTime.POST("", api.SetSystemTime)
		systemTime.POST("ntp/enable", api.NTPEnable)
		systemTime.POST("ntp/disable", api.NTPDisable)
	}

	systemTimeZone := apiRoutes.Group("/timezone")
	{
		systemTimeZone.GET("", api.GetHardwareTZ)
		systemTimeZone.POST("", api.UpdateTimezone)
		systemTimeZone.GET("/list", api.GetTimeZoneList)
		systemTimeZone.POST("/config", api.GenerateTimeSyncConfig)
	}

	systemApi := apiRoutes.Group("/system")
	{
		systemApi.GET("/device", api.GetDeviceInfo)
		systemApi.PATCH("/device", api.UpdateDeviceInfo)
		systemApi.POST("/scanner", api.RunScanner)
		systemApi.GET("/network_interfaces", api.GetNetworkInterfaces)
		systemApi.POST("/reboot", api.RebootHost)
	}

	networking := apiRoutes.Group("/networking")
	{
		networking.GET("", api.Networking)
		networking.GET("/interfaces", api.GetInterfacesNames)
		networking.GET("/internet", api.InternetIP)
	}

	networks := apiRoutes.Group("/networking/networks")
	{
		networks.POST("/restart", api.RestartNetworking)
	}

	networkAddress := apiRoutes.Group("/networking/interfaces")
	{
		networkAddress.POST("/exists", api.DHCPPortExists)
		networkAddress.POST("/auto", api.DHCPSetAsAuto)
		networkAddress.POST("/static", api.DHCPSetStaticIP)
		networkAddress.POST("/reset", api.InterfaceUpDown) //
		networkAddress.POST("/pp", api.InterfaceUp)
		networkAddress.POST("/down", api.InterfaceDown)
	}

	networkFirewall := apiRoutes.Group("/networking/firewall")
	{
		networkFirewall.GET("", api.UWFStatusList)
		networkFirewall.POST("/status", api.UWFStatus)
		networkFirewall.POST("/active", api.UWFActive)
		networkFirewall.POST("/enable", api.UWFEnable)
		networkFirewall.POST("/disable", api.UWFDisable)
		networkFirewall.POST("/port/open", api.UWFOpenPort)
		networkFirewall.POST("/port/close", api.UWFClosePort)
	}

	files := apiRoutes.Group("/files")
	{
		files.GET("/exists", api.FileExists)            // needs to be a file
		files.GET("/walk", api.WalkFile)                // similar as find in linux command
		files.GET("/list", api.ListFiles)               // list all files and folders
		files.POST("/create", api.CreateFile)           // create file only
		files.POST("/copy", api.CopyFile)               // copy either file or folder
		files.POST("/rename", api.RenameFile)           // rename either file or folder
		files.POST("/move", api.MoveFile)               // move files or folders
		files.POST("/upload", api.UploadFile)           // upload single file
		files.POST("/download", api.DownloadFile)       // download single file
		files.GET("/read", api.ReadFile)                // read single file
		files.PUT("/write", api.WriteFile)              // write single file
		files.DELETE("/delete", api.DeleteFile)         // delete single file
		files.DELETE("/delete-all", api.DeleteAllFiles) // deletes file or folder
		files.POST("/write/string", api.WriteStringFile)
		files.POST("/write/json", api.WriteFileJson)
		files.POST("/write/yml", api.WriteFileYml)
	}

	dirs := apiRoutes.Group("/dirs")
	{
		dirs.GET("/exists", api.DirExists)  // needs to be a folder
		dirs.POST("/create", api.CreateDir) // create folder
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
		streamLog.POST("/create", api.CreateLogAndReturn)
		streamLog.DELETE("/:uuid", api.DeleteStreamLog)
		streamLog.DELETE("", api.DeleteStreamLogs)
	}

	device := apiRoutes.Group("/snapshots")
	{
		device.POST("create", api.CreateSnapshot)
		device.POST("restore", api.RestoreSnapshot)
		device.GET("status", api.SnapshotStatus)
	}

	return engine
}
