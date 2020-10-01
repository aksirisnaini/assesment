package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "berhasil",
		})
	})

	router.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "user " + id,
		})
	})

	//query
	router.GET("/user-detail", func(c *gin.Context) {
		id := c.Query("id")
		role := c.DefaultQuery("role", "0")
		c.JSON(200, gin.H{
			"message": "user " + id + " | role " + role,
		})
	})

	//raw
	router.POST("/user", func(c *gin.Context) {
		name := c.PostForm("name")
		address := c.PostForm("address")
		fmt.Println(name)
		c.JSON(200, gin.H{
			"name":    name,
			"address": address,
		})
	})

	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	router.POST("/user-create", func(c *gin.Context) {
		var userVar user
		err := c.Bind(&userVar)
		if err != nil {
		}
		c.JSON(200, gin.H{
			"username": userVar.Username,
			"password": userVar.Password,
		})
	})

	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	{
		v1.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1 get user",
			})
		})
		v1.GET("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v1 get user",
			})
		})
		v2.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2 get user",
			})
		})
		v2.GET("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "v2 get user",
			})
		})

		router.MaxMultipartMemory = 8 << 20 //: 32Mb

		router.POST("/upload", func(c *gin.Context) {
			//Mulitipart form
			form, err := c.MultipartForm()
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
				return
			}
			files := form.File["upload"]
			fmt.Println(len(files))
			var namaUpload, namaTidakUpload string
			for _, file := range files {
				fmt.Println("Size :", file.Size)
				fmt.Println("Nama File ", file.Filename)
				// if file.Size < 200000 {
				// 	namaUpload = namaUpload + " | " + file.Filename
				// 	path := "images/" + file.Filename
				// 	c.SaveUploadedFile(file, path)
				// }
				if strings.HasSuffix(file.Filename, ".jpg") || strings.HasSuffix(file.Filename, ".png") || strings.HasSuffix(file.Filename, ".pdf") {
					namaUpload = namaUpload + " | " + file.Filename
					path := "images/" + file.Filename
					c.SaveUploadedFile(file, path)
				}
				dataFile := file.Header
				for key, value := range dataFile {
					fmt.Println(key, " ", value)
				}

				namaTidakUpload = namaTidakUpload + " | " + file.Filename
				//log.Println(file.Filename)
			}
			c.JSON(200, gin.H{
				"file upload":       namaUpload,
				"file tidak upload": namaTidakUpload,
			})
			//c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
		})
	}
	router.Run()
}
