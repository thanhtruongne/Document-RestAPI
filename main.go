package main

import (
	"fmt"
	"net/http"

	"github.com/user/Practice_api/components/tokenProviders/jwt"
	"github.com/user/Practice_api/middleware"
	"github.com/user/Practice_api/modules/items/transport/gindata"
	"github.com/user/Practice_api/modules/upload"
	"github.com/user/Practice_api/modules/users/storage"
	"github.com/user/Practice_api/modules/users/transport/ginUser"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123@tcp(127.0.0.1:3307)/data_Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect failed", err)
		return
	}
	fmt.Println("Connect success", db)
	r := gin.Default()

	authStore := storage.SqlInstance(db) // tạo5 instance kết nối db
	tokenProviders := jwt.NewTokenJWTProvider("jwt", "22123123")
	authMiddleWare := middleware.AuthenCationRequried(authStore, tokenProviders)
	r.Use(middleware.Recovery())

	v1 := r.Group("/v1")
	{
		//user
		v1.POST("/register", ginUser.RegisterUserData(db))
		v1.POST("/login", ginUser.LoginUserData(db, tokenProviders))
		v1.PUT("/upload", upload.UploadImage(db))
		items := v1.Group("/items")
		{
			items.GET("", gindata.ListItemData(db))
			items.POST("/create", authMiddleWare, gindata.CreateDataItemcompiler(db))
			items.GET("/get/:id", gindata.GetDataItem(db))
			items.GET("/profile", ginUser.Profile())
			// items.PUT("/:id", updateItem(db))
			// items.DELETE("/remove/:id", deleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "oke",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func createItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data models.ItemCreate

// 		if err := c.ShouldBind(data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Create(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": data.Id,
// 		})
// 	}
// }

// func getItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data models.ItemCreate
// 		id, err := strconv.Atoi(c.Param("id")) // convert string to int
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": data,
// 		})

// 	}
// }

// func updateItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var data models.ItemUpdate

// 		id, err := strconv.Atoi(c.Param("id")) // convert string to int
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if err := c.ShouldBind(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  true,
// 			"message": "Cập nhật thành công",
// 		})
// 	}
// }

// func deleteItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {

// 		id, err := strconv.Atoi(c.Param("id")) // convert string to int
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Table(models.Test{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  true,
// 			"message": "Xóa thành công",
// 		})
// 	}
// }

// func readItem(db *gorm.DB) func(*gin.Context) {
// 	return func(c *gin.Context) {
// 		var paging common.Paging
// 		var data []models.Test
// 		if err := c.ShouldBind(&paging); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"message": err.Error(),
// 			})
// 			return
// 		}

// 		paging.Process()

// 		if err := db.Table(models.Test{}.TableName()).Count(&paging.Total).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))

// 	}
// }
