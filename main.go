package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanhettige/httprespond/internal/responses"
)

func main() {
	router := gin.Default()

	router.Any("/200", responses.TwoHundred)
	router.Any("/300", responses.ThreeHundred)
	router.Any("/400", responses.FourHundred)

	router.Run(":3000")
}
