package render

import (
	"image"
	"strings"

	"github.com/fogleman/gg"
)

func renderWidgetFrame(ctx *gg.Context, bounds image.Rectangle, title string) image.Rectangle {
	x := float64(bounds.Min.X)
	y := float64(bounds.Min.Y)
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())
	ctx.SetRGB(0, 0, 0)

	ctx.SetFontFace(headerFont)

	ctx.DrawStringAnchored(strings.ToUpper(title), x+(padding/2), y, 0, 1)

	y += float64(headerFont.Metrics().Height.Round())

	ctx.DrawRectangle(x+(padding/2), y+(padding/2), width-padding, height-padding-float64(headerFont.Metrics().Height.Round()))
	ctx.Stroke()

	x += padding / 2
	y += padding / 2

	return image.Rectangle{
		Min: image.Point{
			X: int(x),
			Y: int(y),
		},
		Max: image.Point{
			X: int(bounds.Max.X - int(padding/2)),
			Y: int(bounds.Max.Y - int(padding/2)),
		},
	}
}
