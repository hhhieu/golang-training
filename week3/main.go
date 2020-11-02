package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	ErrDatabase = 1
)

func (Blog) TableName() string {
	return "blogs"
}

type Blog struct {
	ID      int
	Title   string `form:"title" json:"title" binding:"min=4,max=255"`
	Content string `form:"content"`
}

func main() {
	p := bluemonday.UGCPolicy()
	dns := "hieu:password@tcp(127.0.0.1:3306)/dogfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/health/live", func(c *gin.Context) {
		c.String(200, "ok")
	})
	blogRouter := r.Group("/blog")
	blogRouter.Use(func(c *gin.Context) {
		lang, _ := c.Cookie("lang")
		if lang == "" {
			lang = "en"
		}
		c.Set("lang", lang)
	})
	blogRouter.Use(func(c *gin.Context) {
		lang, _ := c.Cookie("lang")
		if lang == "" {
			lang = "en"
		}
		c.Set("lang", lang)
	})
	blogRouter.POST("/", func(c *gin.Context) {
		lang, _ := c.Get("lang")
		c.Header("Debug", lang.(string))
		blog := Blog{}
		// db.Migrator().AutoMigrate(&Blog{})

		if err := c.ShouldBindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    ErrDatabase,
				"message": " Invalid format data",
				"detail":  err.Error(),
			})
			return
		}
		blog.Title = p.Sanitize(blog.Title)
		blog.Content = p.Sanitize(blog.Content)
		if err := db.Create(&blog).Error; err != nil {
			c.JSON(500, gin.H{
				"code":    1,
				"message": "Can not create blog",
				"detail":  err.Error(),
			})
			return
		}
		c.JSON(200, blog)
	})

	// blogRouter.GET("/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var blog Blog
	// 	result := db.Find(&blog, id)
	// 	c.Header("DEBUG", result.Error())
	// 	c.JSON(200, &blog)
	// })
	blogRouter.GET("/", func(c *gin.Context) {
		blogs := []Blog{}
		strSize := c.DefaultQuery("size", "2")
		strPage := c.DefaultQuery("page", "1")
		q := c.DefaultQuery("q", "")
		size, _ := strconv.ParseInt(strSize, 10, 32)
		page, _ := strconv.ParseInt(strPage, 10, 32)
		builder := db.Limit(int(size)).Offset(int(page)).Order("id desc")
		if q != "" {
			/**
			tao _search (index field), fulltext search
			Analyzer => filter nhung ky tu
			immutable, mutator
			*/
			builder.Where(`title Like ? OR content  LIKE ?`, "%"+q+"%", "%"+q+"%")
		}
		if err := builder.Find(&blogs).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    ErrDatabase,
				"message": "Can't list",
				"detail":  err.Error(),
			})
		}
		c.JSON(200, &blogs)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
