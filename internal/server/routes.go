package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.healthHandler)

  pollRouter := e.Group("/poll")
  pollRouter.GET("/:poll", s.pollHandler)
  pollRouter.POST("/create-poll", s.createPollHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

func (s *Server) pollHandler(c echo.Context) error {
  poll := c.Param("poll")
  return c.JSON(http.StatusOK, poll)
}

func (s *Server) createPollHandler(c echo.Context) error {
  return c.JSON(http.StatusOK, "create poll")
}
