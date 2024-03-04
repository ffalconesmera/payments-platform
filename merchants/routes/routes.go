package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/ffalconesmera/payments-platform/merchants/controller"
	"github.com/ffalconesmera/payments-platform/merchants/middleware"
	"github.com/gin-gonic/gin"
)

type apiRouter struct {
	merchantController controller.MerchantController
}

func InitApiRouter(merchantController controller.MerchantController) *apiRouter {
	return &apiRouter{
		merchantController: merchantController,
	}
}

func (r *apiRouter) SetupPublicRouter(ctx context.Context) *gin.Engine {
	log.Println("initializing public routes..")

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middleware.ContentTypeJsonMiddleware())

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to merchant api..!")
	})

	routesApi := g.Group("/api/v1/merchants")
	{
		routesMerchants := routesApi.Group("/")
		routesMerchants.POST("/login", func(c *gin.Context) { r.merchantController.Login(ctx, c) })
		routesMerchants.POST("/sing-up", func(c *gin.Context) { r.merchantController.SingUp(ctx, c) })
		//routesMerchants.GET("/:merchant_code", func(c *gin.Context) { r.merchantController.FindMerchantByCode(ctx, c) })
	}

	return g
}

func (r *apiRouter) SetupIntranetRouter(ctx context.Context) *gin.Engine {
	log.Println("initializing intranet routes..")

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middleware.ContentTypeJsonMiddleware())

	routesApi := g.Group("/api/v1/merchants")
	{
		routesMerchants := routesApi.Group("/")
		routesMerchants.GET("/:merchant_code", func(c *gin.Context) { r.merchantController.FindMerchantByCode(ctx, c) })
	}

	return g
}

func (r *apiRouter) ListenAndServe(ctx context.Context) {
	s := r.SetupPublicRouter(ctx)
	go s.Run(":8081")

	s2 := r.SetupIntranetRouter(ctx)
	s2.Run(":8085")
}
