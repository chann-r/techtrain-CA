package controllers

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
  router := gin.Default()

  router.GET("/user/get", func(c *gin.Context) { c.JSON(201, "a") })

  Router = router
}
