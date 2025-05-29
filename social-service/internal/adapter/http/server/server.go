package server

// const serverIPAddress = "0.0.0.0:%d"

// type API struct {
// 	server *gin.Engine
// 	cfg    config.HTTPServer
// 	addr   string
// }

// func New(cfg config.Config) *API {
// 	gin.SetMode(cfg.Server.HTTPServer.Mode)
// 	server := gin.New()
// 	server.Use(gin.Recovery())

// 	api := &API{
// 		server: server,
// 		cfg:    cfg.Server.HTTPServer,
// 		addr:   fmt.Sprintf(serverIPAddress, cfg.Server.HTTPServer.Port),
// 	}

// 	api.setupRoutes()

// 	return api
// }

// // func (a *API) setupRoutes() {
// // 	v1 := a.server.Group("/api/v1")
// // 	{
// // 		clients := v1.Group("/file")
// // 		{
// // 			clients.PUT("/", a.fileHandler.Create)
// // 			clients.GET("/:key", a.fileHandler.Get)
// // 			clients.DELETE("/:key", a.fileHandler.Delete)
// // 		}
// // 	}
// // }

// func (a *API) Run() error {
// 	return a.server.Run(a.addr)
// }
