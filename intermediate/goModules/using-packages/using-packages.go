package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.default()
    r.get("/ping", func(c *gin.context) {
        c.json(200, gin.h{
            "message": "pong",
        })
    })
    r.run()
}
