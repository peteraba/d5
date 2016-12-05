package admin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
	})
}

func OkWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		responseStatus: statusOk,
		responseData:   data,
	})
}

func Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		responseStatus: statusError,
	})
}

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		responseStatus: statusError,
		responseError:  errorMissingData,
	})
}

func InternalServerError(c *gin.Context, err error, message string) {
	if message != "" {
		log.Printf("%s: %v", message, err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		responseError: fmt.Sprint(err),
	})
}
