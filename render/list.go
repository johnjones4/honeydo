package render

import (
	"fmt"
	"image"
	"main/service"

	"github.com/fogleman/gg"
)

func renderCardList(ctx *gg.Context, bounds image.Rectangle, items []service.Card) {
	x := float64(bounds.Min.X) + (padding / 2)
	y := float64(bounds.Min.Y) + (padding / 2)
	width := float64(bounds.Dx()) - padding
	ctx.SetFontFace(listItemFont)
	ctx.SetRGB(0, 0, 0)

	for i, item := range items {
		truncated := ctx.WordWrap(fmt.Sprintf("%d. %s", i+1, item.Name), width)
		str := truncated[0]
		if len(truncated) > 1 {
			str += "..."
		}
		ctx.DrawStringAnchored(str, x, y, 0, 1)

		y += float64(listItemFont.Metrics().Height.Round()) * lineHeight
		if y > float64(bounds.Max.Y) {
			return
		}
	}
}
