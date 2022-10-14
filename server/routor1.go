package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
	"net/http"
	"strconv"
)

// CountUsers
// @Summary      Count User
// @Description
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        number query  string false  "Count User"
// @Success      200 {object} map[string]int
// @Failure      400
// @Failure      500
// @Router       /admin/count_users [GET]
func (s Server) CountUsers(c *gin.Context) {
	number := c.Query("number")
	num, _ := strconv.Atoi(number)

	countUser, err := s.repo.CountUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	if num == 0 {
		c.JSON(http.StatusOK, gin.H{
			"jami": countUser[num],
		})
	} else if num > 0 && num < 11 {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("table-%d", num): countUser[num],
		})
	}
}

// CountSum
// @Summary      Count Sum
// @Description
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        number query  string false  "Count sum"
// @Success      200 {object} map[string]int
// @Failure      400
// @Failure      500
// @Router       /admin/count_sum [GET]
func (s Server) CountSum(c *gin.Context) {
	number := c.Query("number")
	num, _ := strconv.Atoi(number)

	countsum, err := s.repo.CountSum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	if num == 0 {
		c.JSON(http.StatusOK, gin.H{
			"jami": countsum[num],
		})
	} else if num > 0 && num < 11 {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("table-%d", num): countsum[num],
		})
	}
}

// ProductList
// @Summary      Product  list
// @Description
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200 {object} []Product
// @Failure      400
// @Failure      500
// @Router       /admin/products [GET]
func (s Server) ProductList(c *gin.Context) {
	product, err := s.repo.ProductList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct
// @Summary      Update Product
// @Description
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id query  string true  "Update Product"
// @Success      200 {object} map[string]bool
// @Failure      400
// @Failure      500
// @Router       /admin/update_product [PUT]
func (s Server) UpdateProduct(c *gin.Context) {
	id := c.Query("id")

	if err := s.repo.UpdateProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Updated": true,
	})
}

// Ingredient
// @Summary      Ingredient
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        id query  string true  "Ingredient"
// @Success      200 {object} []Product
// @Failure      400
// @Failure      500
// @Router       /ingredient [GET]
func (s Server) Ingredient(c *gin.Context) {
	ids := c.Query("id")

	ingre, err := s.repo.Ingredient(ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ingre)
}

// Set
// @Summary      Combo
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        sum query  string true  "Set"
// @Success      200 {object} [][]MenyuJson
// @Failure      400
// @Failure      500
// @Router       /set [GET]
func (s Server) Set(c *gin.Context) {
	sums := c.Query("sum")
	sum, _ := strconv.Atoi(sums)

	set, err := s.repo.Set(sum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, set)
}

// ShopCombo
// @Summary      Shop Combo
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        request body ShopStruct true  "ShopCombo"
// @Success      200 {object} map[string]bool
// @Failure      400
// @Failure      500
// @Router       /shop_combo [POST]
func (s Server) ShopCombo(c *gin.Context) {
	shops := structs.ShopStruct{}

	if err := c.ShouldBindJSON(&shops); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := s.repo.ShopCombo(shops.TableId, shops.FoodId, shops.SaladId, shops.DrinkId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Buy": true,
	})
}

// DeleteBasket
// @Summary      Delete Basket
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Param        request body ShopStruct true  " Delete Basket"
// @Success      200 {object} map[string]bool
// @Failure      400
// @Failure      500
// @Router       /delete_basket [DELETE]
func (s Server) DeleteBasket(c *gin.Context) {
	del := structs.ShopStruct{}

	if err := c.ShouldBindJSON(&del); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := s.repo.DeleteBasket(del.TableId, del.FoodId, del.SaladId, del.DrinkId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"successfully": true,
	})
}

// TableList
// @Summary      Table List
// @Description
// @Tags         others
// @Accept       json
// @Produce      json
// @Success      200 {object} []Table
// @Failure      400
// @Failure      500
// @Router       /tables [GET]
func (s Server) TableList(c *gin.Context) {
	tableList, err := s.repo.TableList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tableList)
}
