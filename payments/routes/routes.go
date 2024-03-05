package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/ffalconesmera/payments-platform/payments/controller"
	"github.com/ffalconesmera/payments-platform/payments/middleware"
	"github.com/gin-gonic/gin"
)

type apiRouter struct {
	paymentController controller.PaymentController
}

func InitApiRouter(ctx context.Context, paymentController controller.PaymentController) *apiRouter {
	return &apiRouter{
		paymentController: paymentController,
	}
}

func (r *apiRouter) SetupApiRouter(ctx context.Context) *gin.Engine {
	log.Println("initializing routes..")

	g := gin.New()
	g.Use(gin.Recovery())

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to payments api..!",
		})
	})

	routesApi := g.Group("/api/v1/payments")
	{
		routesPayments := routesApi.Group("/")
		routesPayments.POST("/checkouts/:merchant_code", func(c *gin.Context) { r.paymentController.CheckoutPayment(ctx, c) }).Use(middleware.ContentTypeJsonMiddleware())
		routesPayments.POST("/process/:payment_code", func(c *gin.Context) { r.paymentController.ProcessPayment(ctx, c) }).Use(middleware.JWTokenMiddleware())
		routesPayments.GET("/:payment_code", func(c *gin.Context) { r.paymentController.CheckPayment(ctx, c) }).Use(middleware.JWTokenMiddleware())
		routesPayments.POST("/refunds/:payment_code", func(c *gin.Context) { r.paymentController.RefundPayment(ctx, c) }).Use(middleware.JWTokenMiddleware())
	}

	return g
}

func (r *apiRouter) ListenAndServe(ctx context.Context) {
	s := r.SetupApiRouter(ctx)
	s.Run(":8082")
}
