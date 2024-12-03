package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhongxic/gin-template/pkg/result"
)

type Controller struct {
}

func (p *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, result.SuccessWithData("pong"))
}
