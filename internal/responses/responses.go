package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TwoHundred(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func FourHundred(c *gin.Context) {
	c.String(http.StatusBadRequest, "Four hundred")
}
