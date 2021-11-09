package render

import (
	"os"
	"path"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

const (
	width                      float64 = 1200.0
	height                     float64 = 825.0
	padding                    float64 = width * 0.01
	calendarWidth              float64 = width * 0.61803398875
	listItemFontSize           float64 = height * 0.02
	listItemFontFace                   = "OpenSans-Regular.ttf"
	headerFontSize             float64 = listItemFontSize * 1.1
	headerFontFace                     = "OpenSans-Bold.ttf"
	lineHeight                 float64 = 1.2
	calendarDayFontFace                = "OpenSans-Regular.ttf"
	calendarCurrentDayFontFace         = "OpenSans-Bold.ttf"
	calendarDayFontSize        float64 = listItemFontSize * 0.75
	calendarRows                       = 6.0
	calendarItemFontFace               = "OpenSans-Regular.ttf"
	calendarItemFontSize       float64 = listItemFontSize * 0.7
)

var (
	listItemFont           font.Face
	headerFont             font.Face
	calendarDayFont        font.Face
	calendarCurrentDayFont font.Face
	calendarItemFont       font.Face
	updateTimeFont         font.Face
)

var (
	daysOfTheWeek = []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}
)

func LoadFonts() error {
	var err error

	listItemFont, err = gg.LoadFontFace(path.Join(os.Getenv("FONTS_PATH"), listItemFontFace), listItemFontSize)
	if err != nil {
		return err
	}

	headerFont, err = gg.LoadFontFace(path.Join(os.Getenv("FONTS_PATH"), headerFontFace), headerFontSize)
	if err != nil {
		return err
	}

	calendarDayFont, err = gg.LoadFontFace(path.Join(os.Getenv("FONTS_PATH"), calendarDayFontFace), calendarDayFontSize)
	if err != nil {
		return err
	}

	calendarCurrentDayFont, err = gg.LoadFontFace(path.Join(os.Getenv("FONTS_PATH"), calendarCurrentDayFontFace), calendarDayFontSize)
	if err != nil {
		return err
	}

	calendarItemFont, err = gg.LoadFontFace(path.Join(os.Getenv("FONTS_PATH"), calendarItemFontFace), calendarItemFontSize)
	if err != nil {
		return err
	}

	updateTimeFont = calendarItemFont

	return nil
}
