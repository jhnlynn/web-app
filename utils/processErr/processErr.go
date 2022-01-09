package processErr

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProcessInternalErr general error responding
func ProcessInternalErr(c *gin.Context, err error, state string) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf(state, err),
		})
		panic(err)
		return
	}
}

func ProcessBadRequestErr(c *gin.Context, state string) {

}
