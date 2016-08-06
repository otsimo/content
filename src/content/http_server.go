package content

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/otsimo/health"
)

const (
	WikiEndpoint    = "/wiki"
	WebhookEndpoint = "/webhook"
	HealthEndpoint  = "/health"
)

type Response struct {
	Type int         `json:"type"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse returns a *Response object for errors
func NewErrorResponse(resType int, message string) *Response {
	return &Response{
		Type: resType,
		Data: &ErrorResponse{
			Message: message,
		},
	}
}

func httpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Error()
	}
	r := c.Response()
	if !r.Committed() {
		b, _ := json.Marshal(NewErrorResponse(code, msg))
		r.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		r.WriteHeader(code)
		r.Write(b)
	}
	logrus.Errorln(err)
}

func (s *Server) healthHandler(ctx echo.Context) error {
	if err := health.Check(s.checks); err != nil {
		logrus.Errorf("health check failed, %v", err)
		return ctx.JSON(http.StatusInternalServerError, health.StatusResponse{
			Status: "error",
			Details: &health.StatusResponseDetails{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		})
	} else {
		return ctx.JSON(http.StatusOK, health.StatusResponse{
			Status: "ok",
		})
	}
}

func (s *Server) HttpServer() *echo.Echo {
	// Echo instance
	e := echo.New()
	e.SetHTTPErrorHandler(httpErrorHandler)

	// Debug mode
	if s.Config.Debug {
		e.Debug()
	}
	// Logger
	cnf := LoggerConfig{
		Format: "time=\"${time_rfc3339}\" remote_ip=${remote_ip} method=${method} " +
			"uri=${uri} status=${status} took=${response_time}, sent=${response_size} bytes\n",
		Output: os.Stdout,
	}

	e.Use(LoggerWithConfig(cnf))
	e.Use(mw.Recover())

	e.Get(HealthEndpoint, s.healthHandler)
	e.Post(WebhookEndpoint, s.webhookHandler)
	for _, v := range s.Content.GitPublicDirs {
		e.Static(v.Path, v.Dir)
	}
	e.Static(WikiEndpoint, s.Config.PublicDir)

	return e
}
