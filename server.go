package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Skus        []skus `json:"skus"`
}

type skus struct {
	Sku            string  `json:"sku"`
	Price          float64 `json:"price"`
	Stock          int     `json:"stock"`
	VariationType  string  `json:"variation_type"`
	VariationValue string  `json:"variation_value"`
}

var products = []product{
	{
		ID: 1, Title: "Ar condicionado", Description: "Ar condicionado",
		Skus: []skus{
			{Sku: "CODIGO001", Price: 179.99, Stock: 55, VariationType: "voltage", VariationValue: "110V"},
			{Sku: "CODIGO002", Price: 239.99, Stock: 45, VariationType: "voltage", VariationValue: "220V"},
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/products", getAllProducts)
	router.GET("/products/:id", getProductById)
	router.POST("/products", postProduct)
	router.PUT("/products/:id", putProduct)

	router.Run("localhost:8080")
}

func getAllProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func postProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func putProduct(c *gin.Context) {
	var newProduct product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product ID not found!"})
		return
	}

	for k, a := range products {
		if a.ID == id {
			products[k] = newProduct
			c.IndentedJSON(http.StatusOK, products[k])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product ID not found!"})
}

func getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Loop through the list of products, looking for
	// an album whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found!"})
}
