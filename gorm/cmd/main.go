package main

import (
	"casbin/gorm"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	Enforcer *casbin.Enforcer
)

func main() {
	initCasbin()
	g := gin.Default()
	api := g.Group("/api/v1", CasbinMiddleware)
	{
		api.GET("/user", test)
	}
	g.Run(":8001")
}

func CasbinMiddleware(c *gin.Context) {
	// userName just for test
	userName := c.DefaultQuery("u", "guest")
	path := c.Request.URL.Path
	method := c.Request.Method
	res, err := Enforcer.Enforce(userName, path, method)
	if err != nil || !res {
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}

	c.Next()
}

func initCasbin() {
	adapter, err := gormadapter.NewAdapterByDB(gorm.DB)
	if err != nil {
		log.Fatal(err)
	}

	Enforcer, err = casbin.NewEnforcer("../conf/model.conf", adapter)
	if err != nil {
		log.Fatalf("NewEnforcer error: %v", err)
	}
	Enforcer.LoadPolicy()
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
	return
}
