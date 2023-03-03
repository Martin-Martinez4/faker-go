package english

import (
	"fmt"
	"strconv"
)

type Colors struct {
	colornames  *[]colorname
	ColorSpaces *Colorspaces
}

var colors = Colors{
	colornames:  &colornames,
	ColorSpaces: &colorspaces,
}

func (e *EnDict) Color(seed_optional ...int64) string {

	index := GenerateRandIntBelowMaximum(len(*e.colornames), seed_optional...)

	return (*e.colornames)[index]

}

func (e *EnDict) Colorhex(seed_optional ...int64) string {

	var red string
	var blue string
	var green string

	redint := GenerateRandIntBelowMaximum(256, seed_optional...)
	greenint := GenerateRandIntBelowMaximum(256, seed_optional...)
	blueint := GenerateRandIntBelowMaximum(256, seed_optional...)

	red = strconv.FormatInt(int64(redint), 16)
	green = strconv.FormatInt(int64(greenint), 16)
	blue = strconv.FormatInt(int64(blueint), 16)

	if len(red) == 1 {

		red = fmt.Sprintf(red + "0")

	}

	if len(blue) == 1 {

		blue = fmt.Sprintf(blue + "0")

	}

	if len(green) == 1 {

		green = fmt.Sprintf(green + "0")

	}

	return fmt.Sprintf("#%s%s%s", red, green, blue)
}

func (e *EnDict) Colorrgb(seed_optional ...int64) string {

	red := GenerateRandIntBelowMaximum(255, seed_optional...)
	green := GenerateRandIntBelowMaximum(255, seed_optional...)
	blue := GenerateRandIntBelowMaximum(255, seed_optional...)

	return fmt.Sprintf("%s(%d,%d,%d)", e.ColorSpaces.rgb, red, green, blue)

}

func (e *EnDict) Colorrgba(seed_optional ...int64) string {

	red := GenerateRandIntBelowMaximum(255, seed_optional...)
	green := GenerateRandIntBelowMaximum(255, seed_optional...)
	blue := GenerateRandIntBelowMaximum(255, seed_optional...)

	alpha := RandomFloatBetweenTwoNumbers(0, 1, seed_optional...)

	return fmt.Sprintf("%s(%d,%d,%d,%f)", e.ColorSpaces.rgba, red, green, blue, alpha)

}

func (e *EnDict) Colorhsl(seed_optional ...int64) string {

	hue := GenerateRandIntBelowMaximum(359, seed_optional...)
	saturation := GenerateRandIntBelowMaximum(101, seed_optional...)
	lightness := GenerateRandIntBelowMaximum(101, seed_optional...)

	return fmt.Sprintf("%s(%d,%d%%,%d%%)", e.ColorSpaces.hsl, hue, saturation, lightness)
}

func (e *EnDict) Colorhsla(seed_optional ...int64) string {

	hue := GenerateRandIntBelowMaximum(359, seed_optional...)
	saturation := GenerateRandIntBelowMaximum(101, seed_optional...)
	lightness := GenerateRandIntBelowMaximum(101, seed_optional...)
	alpha := RandomFloatBetweenTwoNumbers(0, 1, seed_optional...)

	return fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", e.ColorSpaces.hsla, hue, saturation, lightness, alpha)
}
