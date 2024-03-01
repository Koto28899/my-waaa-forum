package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func InvalidRequestParamErr(param string) error {
	return fmt.Errorf("invalid request parameter key: `%s`", param)
}
