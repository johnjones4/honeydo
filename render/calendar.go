package render

import (
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/apognu/gocal"
	"github.com/fogleman/gg"
)

func renderCalendar(ctx *gg.Context, bounds image.Rectangle, events []gocal.Event) {
	ctx.SetFontFace(calendarDayFont)
	ctx.SetRGB(0, 0, 0)

	dayWidth := float64(bounds.Dx()) / float64(len(daysOfTheWeek))

	daysHeight := calendarDayFontSize * 2

	y := float64(bounds.Min.Y)

	for i, d := range daysOfTheWeek {
		x := float64(bounds.Min.X) + dayWidth*float64(i)
		ctx.DrawStringAnchored(strings.ToUpper(d.String()), x+(dayWidth/2.0), y+(daysHeight/2), 0.5, 0.4)
		if i != 0 {
			ctx.DrawLine(x, y, x, float64(bounds.Max.Y))
			ctx.Stroke()
		}
	}

	y += daysHeight

	ctx.DrawLine(float64(bounds.Min.X), y, float64(bounds.Max.X), y)
	ctx.Stroke()

	rowOrigin := image.Point{
		X: bounds.Min.X,
		Y: int(y),
	}

	dayHeight := (float64(bounds.Max.Y) - y) / calendarRows

	now := time.Now().Local()
	date := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	startIndex := columnForDate(date)
	nextDate := drawRow(ctx, rowOrigin, dayWidth, dayHeight, startIndex, date, events)

	for {
		y += dayHeight

		if y >= float64(bounds.Max.Y) {
			return
		}

		rowOrigin = image.Point{
			X: bounds.Min.X,
			Y: int(y),
		}
		nextDate = drawRow(ctx, rowOrigin, dayWidth, dayHeight, 0, nextDate, events)

		ctx.SetRGB(0, 0, 0)
		ctx.DrawLine(float64(bounds.Min.X), y, float64(bounds.Max.X), y)
		ctx.Stroke()

	}
}

func columnForDate(d time.Time) int {
	for i, dw := range daysOfTheWeek {
		if dw == d.Weekday() {
			return i
		}
	}
	return -1
}

func drawRow(ctx *gg.Context, origin image.Point, dayWidth, dayHeight float64, startIndex int, startDate time.Time, events []gocal.Event) time.Time {
	date := startDate
	for i := range daysOfTheWeek {
		x := float64(origin.X) + dayWidth*float64(i)
		if i < startIndex {
			ctx.SetRGB(0.95, 0.95, 0.95)
			ctx.DrawRectangle(x+1, float64(origin.Y)+2, dayWidth-1, dayHeight-2)
			ctx.Fill()
		} else {
			tileOrigin := image.Point{
				X: int(x + padding/2),
				Y: origin.Y + int(padding/2),
			}
			tileBounds := image.Rectangle{
				Min: tileOrigin,
				Max: image.Point{
					X: tileOrigin.X + int(dayWidth-padding),
					Y: tileOrigin.Y + int(dayHeight-padding),
				},
			}

			drawDayTile(ctx, tileBounds, date, events)

			date = date.Add(time.Hour * 24)
		}
	}
	return date
}

func drawDayTile(ctx *gg.Context, bounds image.Rectangle, date time.Time, events []gocal.Event) {
	if sameDay(time.Now(), date) {
		ctx.SetFontFace(calendarCurrentDayFont)
	} else {
		ctx.SetFontFace(calendarDayFont)
	}
	ctx.SetRGB(0, 0, 0)
	dateStr := fmt.Sprint(date.Day())
	ctx.DrawStringAnchored(dateStr, float64(bounds.Min.X), float64(bounds.Min.Y), 0, 1)

	ctx.SetFontFace(calendarItemFont)

	y := float64(bounds.Min.Y) + float64(calendarDayFont.Metrics().Height.Round())*1.75
	for _, event := range events {
		spanning := (event.Start.Local().Before(date) && event.End.Local().After(date))
		if (sameDay(event.Start.Local(), date) && sameDay(event.End.Local(), date)) || spanning {
			var thisLineHeight float64
			var summary string
			if spanning {
				summary = event.Summary
				y += padding / 2
				rectHeight := float64(calendarItemFont.Metrics().Height.Round()) + padding
				thisLineHeight = rectHeight
				ctx.SetRGB(0.25, 0.25, 0.25)
				ctx.DrawRectangle(float64(bounds.Min.X)-padding/2, y-padding/2, float64(bounds.Dx())+padding, rectHeight)
				ctx.Fill()
				ctx.SetRGB(1, 1, 1)
			} else {
				summary = fmt.Sprintf("%s %s", event.Summary, event.Start.Local().Format("3:04pm"))
				thisLineHeight = float64(calendarItemFont.Metrics().Height.Round()) * lineHeight
				ctx.SetRGB(0, 0, 0)
			}
			strs := ctx.WordWrap(summary, float64(bounds.Dx()))
			ctx.DrawStringAnchored(strs[0], float64(bounds.Min.X), y, 0, 1)
			y += thisLineHeight
			if y >= float64(bounds.Max.Y) {
				return
			}
		}
	}
}

func sameDay(d1, d2 time.Time) bool {
	return d1.Year() == d2.Year() && d1.Month() == d2.Month() && d1.Day() == d2.Day()
}
