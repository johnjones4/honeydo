package render

import (
	"fmt"
	"image"
	"main/service"
	"time"

	"github.com/apognu/gocal"
	"github.com/fogleman/gg"
)

type List struct {
	ID    string
	Title string
	Cards []service.Card
}

func Render(events []gocal.Event, lists []List) *gg.Context {
	ctx := gg.NewContext(int(width), int(height))

	ctx.SetRGB(1, 1, 1)
	ctx.DrawRectangle(0, 0, width, height)
	ctx.Fill()

	ctx.SetRGB(0, 0, 0)
	ctx.SetFontFace(updateTimeFont)
	updateStr := fmt.Sprintf("Last Updated: %s", time.Now().Local().Format("Jan 2 @ 3:04:05 PM"))
	ctx.DrawStringAnchored(updateStr, width/2, height-float64(updateTimeFont.Metrics().Height.Round())/3, 0.5, 0)
	realHeight := height - float64(updateTimeFont.Metrics().Height.Round())*1.1

	cw := calendarWidth
	calendarWidgetBounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: int(padding / 2),
		},
		Max: image.Point{
			X: int(cw),
			Y: int(realHeight),
		},
	}
	calendarBounds := renderWidgetFrame(ctx, calendarWidgetBounds, "Calendar")
	renderCalendar(ctx, calendarBounds, events)

	listHeight := (realHeight - (padding / 2)) / float64(len(lists))

	x := calendarWidth
	listWidth := width - x
	lh := listHeight
	for i, list := range lists {
		origin := image.Point{
			X: int(x),
			Y: int(padding/2) + (i * int(listHeight)),
		}
		listWidgetBounds := image.Rectangle{
			Min: origin,
			Max: image.Point{
				X: origin.X + int(listWidth),
				Y: origin.Y + int(lh),
			},
		}
		listBounds := renderWidgetFrame(ctx, listWidgetBounds, list.Title)
		renderCardList(ctx, listBounds, list.Cards)
	}

	return ctx
}
