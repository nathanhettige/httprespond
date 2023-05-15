package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Any("/:input", httpResponse)
	router.Run(":3000")
}

func httpResponse(c *gin.Context) {
	input := c.Param("input")

	status, err := strconv.Atoi(input)
	if err != nil {
		c.String(http.StatusNotAcceptable, fmt.Sprintf("Invalid http response status code - %s.", input))
		return
	}

	if http.StatusText(status) == "" {
		c.String(http.StatusNotAcceptable, fmt.Sprintf("HTTP Status code %d doesn't exist.", status))
		return
	}

	c.String(status, fmt.Sprintf("%d %s", status, http.StatusText(status)))
}
