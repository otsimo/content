package content

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

const (
	WikiEndpoint    = "/wiki"
	WebhookEndpoint = "/webhook"
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

func httpErrorHandler(err error, c *echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code()
		msg = he.Error()
	}
	r := c.Response()
	if !r.Committed() {
		b, _ := json.Marshal(NewErrorResponse(code, msg))
		r.Header().Set(echo.ContentType, echo.ApplicationJSONCharsetUTF8)
		r.WriteHeader(code)
		r.Write(b)
	}
	logrus.Errorln(err)
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
	e.Use(logger())
	e.Use(mw.Recover())

	e.Post(WebhookEndpoint, s.webhookHandler)
	for _, v := range s.Content.GitPublicDirs {
		e.Static(v.Path, v.Dir)
	}
	e.Static(WikiEndpoint, s.Config.PublicDir)

	return e
}
