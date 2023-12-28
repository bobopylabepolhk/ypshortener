package echoeasyjson

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type EasyJSONSerializer struct{}

// Serialize and Deserialize fallback to encoding/json
// echo uses the supplied JSONSerizalier for any under the hood operation within context
// some JSON marshalling/unmarshalling is done on plain maps e.x error responses
// https://github.com/labstack/echo/blob/master/echo.go#L453

func (s EasyJSONSerializer) Serialize(ctx echo.Context, i interface{}, _ string) error {
	o, ok := i.(easyjson.Marshaler)
	if ok {
		_, err := easyjson.MarshalToWriter(o, ctx.Response())
		return err
	}

	enc := json.NewEncoder(ctx.Response())
	return enc.Encode(i)
}

func (s EasyJSONSerializer) Deserialize(ctx echo.Context, i interface{}) error {
	o, ok := i.(easyjson.Unmarshaler)

	if ok {
		err := easyjson.UnmarshalFromReader(ctx.Request().Body, o)
		return err
	}

	err := json.NewDecoder(ctx.Request().Body).Decode(i)
	return err
}
