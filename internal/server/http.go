package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/docs"
	"github.com/go-nunu/nunu-layout-advanced/internal/handler"
	"github.com/go-nunu/nunu-layout-advanced/internal/middleware"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/go-nunu/nunu-layout-advanced/pkg/server/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *config.Config,
	userHandler *handler.UserHandler,
	am *middleware.AuthMiddleware,
) *http.Server {
	if conf.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.Http.Host),
		http.WithServerPort(conf.Http.Port),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.GET("/api/docs/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	v1 := s.Group("/api/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(am.NoStrictAuth())
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(am.StrictAuth())
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
