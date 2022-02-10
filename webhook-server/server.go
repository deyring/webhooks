package webhookserver

import (
	"io/ioutil"
	"net/http"

	"github.com/deyring/webhooks/handler"
	"github.com/gin-gonic/gin"
)

func New(handler *handler.WebhookHandler) *Server {
	return &Server{
		webhookHandler: handler,
	}
}

type Server struct {
	webhookHandler *handler.WebhookHandler
}

func (s *Server) ListenAndServe() error {
	r := gin.Default()
	r.POST("/webhooks/:hook", s.handleAll)
	return r.Run()
}

func (s *Server) handleAll(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}
	err = s.webhookHandler.Handle(c.Param("hook"), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
