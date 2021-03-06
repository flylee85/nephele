package handler

import (
	"strings"

	"github.com/ctripcorp/nephele/codec"
	"github.com/ctripcorp/nephele/context"
)

func getImageHandler() HandlerFunc {
	return func(ctx *context.Context) {
		decoder := codec.GetDecoder(ctx)
		if err := decoder.Decode(strings.TrimPrefix(ctx.HTTP().Request.RequestURI, "/image/")); err != nil {
			ctx.HTTP().String(404, err.Error())
			return
		}
		image, err := decoder.CreateIndex().FindOriginalImage()
		if err != nil {
			ctx.HTTP().String(404, err.Error())
			return
		}
		if err := image.Use(decoder.Transformer()).Transform(ctx); err != nil {
			ctx.HTTP().String(404, err.Error())
			return
		}
		ctx.HTTP().Writer.Write(image.Blob())
	}
}
