package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/heat1q/boardsite/api/log"
)

func RequestLogger(logger echo.Logger) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// measure response time
			start := time.Now()

			// set logger to request context
			ctx := context.WithValue(c.Request().Context(), log.ContextKey, logger)
			c.SetRequest(c.Request().WithContext(ctx))

			var reqBody, respBody []byte
			cfg := middleware.BodyDumpConfig{
				Handler: func(c echo.Context, req []byte, resp []byte) {
					reqBody = req
					// only log JSON responses, no png etc.
					if len(resp) < 2<<10 && json.Valid(resp) {
						respBody = resp
					}
				},
			}
			bodyDump := middleware.BodyDumpWithConfig(cfg)

			err := bodyDump(next)(c)

			sb := strings.Builder{}
			sb.WriteString(fmt.Sprintf("incoming request (%d) %s %s -- %s",
				c.Response().Status,
				c.Request().Method,
				c.Request().RequestURI,
				c.Request().Header.Get("User-Agent"),
			))

			if len(reqBody) > 0 {
				var body bytes.Buffer
				if err := json.Compact(&body, reqBody); err == nil {
					sb.WriteString(fmt.Sprintf(" -- req body: %s", body.String()))
				} else {
					sb.WriteString(fmt.Sprintf(" -- req body content length: %d", len(reqBody)))
				}
			}

			if len(respBody) > 0 {
				// dont spam the logs with huge responses
				if len(respBody) < 2<<10 && json.Valid(respBody) {
					sb.WriteString(fmt.Sprintf(" -- resp body: %s", string(respBody)))
				} else {
					sb.WriteString(fmt.Sprintf(" -- resp body content length: %d", len(respBody)))
				}
			}

			if err != nil {
				sb.WriteString(fmt.Sprintf(" -- error: %v", err))
			}

			elapsed := time.Since(start)
			sb.WriteString(fmt.Sprintf(" -- took %s", elapsed))

			if err != nil {
				log.Ctx(c.Request().Context()).Error(sb.String())
			} else {
				log.Ctx(c.Request().Context()).Info(sb.String())
			}

			return err
		}
	}
}
