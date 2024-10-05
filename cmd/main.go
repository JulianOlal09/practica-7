package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()
	users := []User{}
	indexUser := 1
	fmt.Println("Running App")
	//Tomar archivos de la carpeta templates
	router.LoadHTMLGlob("templates/*")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	//API URLs
	router.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	router.POST("/api/users", func(c *gin.Context) {
		var user User
		if c.BindJSON(&user) == nil {
			user.Id = indexUser
			users = append(users, user)
			indexUser++
			c.JSON(200, user)
		} else {
			c.JSON(400, gin.H{
				"error": "invalid payload",
			})
		}
	})
	//Eliminacion de usuarios
	router.DELETE("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}
		fmt.Println("Id a borrar: ", id)
		for i, user := range users {
			if user.Id == idParsed {
				users = append(users[:1], users[i+1:]...)
				c.JSON(200, gin.H{
					"message": "User Deleted",
				})
				return
			}
		}
		c.JSON(201, gin.H{})
	})
	//Actualizar usuarios
	router.PUT("/api/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idParsed, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid id",
			})
			return
		}
		var user User
		err = c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid Payload",
			})
			return
		}
		fmt.Println("Id a actualizar", id)
		for i, u := range users {
			if u.Id == idParsed {
				users[i] = user
				users[i].Id = idParsed
				c.JSON(200, users[i])
				return
			}
		}
		c.JSON(201, gin.H{})
	})
	router.Run(":8001")
}
