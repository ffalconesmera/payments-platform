package routes

import (
	"log"
	"net/http"

	"github.com/ffalconesmera/payments-platform/merchants/controller"
	"github.com/ffalconesmera/payments-platform/merchants/middleware"
	"github.com/gin-gonic/gin"
)

type apiRouter struct {
	merchantController controller.MerchantController
}

func InitApiRouter(merchantController *controller.MerchantController) *apiRouter {
	return &apiRouter{
		merchantController: *merchantController,
	}
}

func (cr *apiRouter) SetupPublicRouter() *gin.Engine {
	log.Println("initializing public routes..")

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middleware.ContentTypeJsonMiddleware())

	g.GET("/running", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to merchants api..!")
	})

	routesApi := g.Group("/api/v1/merchants")
	{
		routesMerchants := routesApi.Group("/")
		routesMerchants.POST("/login", func(c *gin.Context) { cr.merchantController.Login(c) })
		routesMerchants.POST("/sign-up", func(c *gin.Context) { cr.merchantController.SingUp(c) })
		//routesMerchants.GET("/:merchant_code", func(c *gin.Context) { r.merchantController.FindMerchantByCode(ctx, c) })
	}

	return g
}

func (cr *apiRouter) SetupIntranetRouter() *gin.Engine {
	log.Println("initializing intranet routes..")

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(middleware.ContentTypeJsonMiddleware())

	routesApi := g.Group("/api/v1/merchants")
	{
		routesMerchants := routesApi.Group("/")
		routesMerchants.GET("/:merchant_code", func(c *gin.Context) { cr.merchantController.FindMerchantByCode(c) })
	}

	return g
}

func (cr *apiRouter) ListenAndServe() {
	s := cr.SetupPublicRouter()
	go s.Run(":8081")

	s2 := cr.SetupIntranetRouter()
	s2.Run(":8085")
}
