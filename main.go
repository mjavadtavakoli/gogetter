package main

import (
	"gogetters/routes" // برای main.go

	"github.com/gin-gonic/gin"
	// در فایل‌های مربوطه
)

func main() {
	r := gin.Default()

	routes.UserRoutes(r)

	r.Run(":8080")
}
