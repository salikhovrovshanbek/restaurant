package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	structs "github.com/gokurs/Projects/restaurant/repository/struct"
)

// FoodGet1
// @Summary      National Food
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Success      200 {object} []MenyuJson
// @Failure      400
// @Failure      500
// @Router       /menu/food1 [GET]
func (h Server) FoodGet1(c *gin.Context) {
	r1, err := h.repo.Food1()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r1)
}

// FoodGet2
// @Summary      European Food
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Success      200 {object} []MenyuJson
// @Failure      400
// @Failure      500
// @Router       /menu/food2 [GET]
func (h Server) FoodGet2(c *gin.Context) {
	r2, err := h.repo.Food2()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r2)
}

// FoodGet3
// @Summary      Turkish Food
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Success      200 {object} []MenyuJson
// @Failure      400
// @Failure      500
// @Router       /menu/food3 [GET]
func (h Server) FoodGet3(c *gin.Context) {
	r3, err := h.repo.Food3()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, r3)

}

// SaladGet
// @Summary      Salad Menu
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Success      200 {object} []MenyuJson
// @Failure      400
// @Failure      500
// @Router       /menu/salad [GET]
func (h Server) SaladGet(c *gin.Context) {
	r4, err := h.repo.Salad()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, r4)

}

// DrinksGet
// @Summary      Drinks Menu
// @Description
// @Tags         menu
// @Accept       json
// @Produce      json
// @Success      200 {object} []MenyuJson
// @Failure      400
// @Failure      500
// @Router       /menu/drinks [GET]
func (h Server) DrinksGet(c *gin.Context) {
	r5, err := h.repo.Drinks()

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	}
	c.JSON(200, r5)
}

// OpenChek
// @Summary      Open Chek
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        table_id query string true "table id"
// @Success      200 {object} map[string]bool
// @Failure      400
// @Failure      500
// @Router       /open_chek [POST]
func (h Server) OpenChek(c *gin.Context) {
	t_id := c.Query("table_id")
	err := h.repo.OpenChek(t_id)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}

// Shop
// @Summary      Shop
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        request body ShopStruct true  "Shop"
// @Success      200 {object} map[string]bool
// @Failure      400
// @Failure      500
// @Router       /shop [POST]
func (h Server) Shop(c *gin.Context) {
	var req structs.ShopStruct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.repo.Shop(req.TableId, req.FoodId, req.SaladId, req.DrinkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

// Chek
// @Summary      Chek
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        table_id query string true "table id"
// @Success      200 {object} map[string]uint32
// @Failure      400
// @Failure      500
// @Router       /chek [GET]
func (h Server) Chek(c *gin.Context) {
	a := c.Query("table_id")
	r8, err := h.repo.Chek(a)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, r8)
}
