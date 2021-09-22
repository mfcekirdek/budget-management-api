package middleware

import "github.com/labstack/echo/v4"

func SetCommonHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		context.Response().Header().Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		return next(context)
	}
}
