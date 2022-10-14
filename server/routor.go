package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gokurs/Projects/restaurant/config"
	_ "github.com/gokurs/Projects/restaurant/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net"
)

type Server struct {
	repo Repository
}

// @title           Postgres Crud API
// @version         1.0
// @description     This is a sample server celler server.

// NewRoutor
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func NewRoutor(repo Repository, cfg config.Config) {
	s := Server{
		repo: repo,
	}
	r := gin.Default()
	menu := r.Group("/menu")
	menu.GET("/food1", s.FoodGet1)
	menu.GET("/food2", s.FoodGet2)
	menu.GET("/food3", s.FoodGet3)
	menu.GET("/salad", s.SaladGet)
	menu.GET("/drinks", s.DrinksGet)
	r.POST("/open_chek", s.OpenChek)
	r.POST("/shop", s.Shop)
	r.GET("/chek", s.Chek)
	r.GET("/ingredient", s.Ingredient)
	r.GET("/set", s.Set)
	r.GET("/shop_combo", s.ShopCombo)
	r.DELETE("/delete_basket", s.DeleteBasket)
	r.GET("/tables", s.TableList)
	admin := r.Group("/admin")
	admin.GET("/count_users", s.CountUsers)
	admin.GET("/count_sum", s.CountSum)
	admin.GET("/products", s.ProductList)
	admin.PUT("/update_product", s.UpdateProduct)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(net.JoinHostPort(cfg.Host, cfg.Port))
}
