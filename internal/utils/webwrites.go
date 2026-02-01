package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func WriteJSON(ctx *fasthttp.RequestCtx, data interface{}) {
	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(data)
}

func WriteError(ctx *fasthttp.RequestCtx, err error, code int) {
	ctx.SetStatusCode(code)
	WriteJSON(ctx, map[string]string{"error": err.Error()})
}
