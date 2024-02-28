package middlewares

import (
	"bytes"
	"io"
	"time"

	. "github.com/alexfordev/WebAlbums/log"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		start := time.Now()
		data, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewReader(data))
		// body := string(data)
		c.Next()

		responseBody := bodyLogWriter.body.String()

		end := time.Now()

		latency := end.Sub(start)

		path := c.Request.URL.Path + "?" + c.Request.URL.RawQuery

		clientIP := c.ClientIP()
		method := c.Request.Method
		// header := c.Request.Header
		statusCode := c.Writer.Status()

		// Log.Infof("c.Request.URL.Path=%s", c.Request.URL.Path)
		if c.Request.URL.Path == "/touchfan/v2/orders/action/3" {
			Log.Debugf("| %3d | %13v | %15s | %s  %s |request body %s |response body %s |",
				statusCode,
				latency,
				clientIP,
				method, path,
				"",
				"",
			)
			return
		}
		if c.Request.URL.Path == "/touchfan/v2/orders/action/" ||
			c.Request.URL.Path == "/touchfan/v2/orders/action/1" ||
			c.Request.URL.Path == "/touchfan/v2/orders" ||
			c.Request.URL.Path == "/touchfan/v2/metrics" {
			return
		}

		if c.Request.URL.Path == "/touchfan/v2/txs/action/2" {
			if responseBody == "" { // || responseBody == "{\"eCode\":1414,\"eMsg\":\"requests too many\"}" {
				return
			}

		} else {
			// Log.Infof("header:=%#v", header)
		}

		// Log.Infof("c.Request.URL.Path=%s", c.Request.URL.Path)
		// Log.Infof("header:=%#v", header)

		// Log.Debugf("| %3d | %13v | %15s | %s  %s |request body %s |response body %s |",
		// 	statusCode,
		// 	latency,
		// 	clientIP,
		// 	method, path,
		// 	body,
		// 	responseBody,
		// )
	}
}
