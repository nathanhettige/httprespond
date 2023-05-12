package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TwoHundred(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func ThreeHundred(c *gin.Context) {
	c.String(http.StatusMultipleChoices, "Multiple Choices")

}

func FourHundred(c *gin.Context) {
	c.String(http.StatusBadRequest, "Bad Request")
}
