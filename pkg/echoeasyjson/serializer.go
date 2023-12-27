package echoeasyjson

import (
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type EasyJSONSerializer struct{}

func (s EasyJSONSerializer) Serialize(ctx echo.Context, i interface{}, indent string) error {
	_, err := easyjson.MarshalToWriter(i.(easyjson.Marshaler), ctx.Response())
	// TODO try to implement prettified json

	return err
}

func (s EasyJSONSerializer) Deserialize(ctx echo.Context, i interface{}) error {
	err := easyjson.UnmarshalFromReader(ctx.Request().Body, i.(easyjson.Unmarshaler))

	return err
}
