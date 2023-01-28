package main

import "fmt"

type Colors struct {
	ColorNames  *ColorNames
	ColorSpaces *Colorspaces
}

var colors = Colors{
	ColorNames:  &colornames,
	ColorSpaces: &colorspaces,
}

func (c *Colors) Color() string {

	index := GenerateRandIntBelowMaximum(len(*c.ColorNames))

	return (*c.ColorNames)[index]

}

func (c *Colors) Colorhex() string {

	red := GenerateRandIntBelowMaximum(256)
	green := GenerateRandIntBelowMaximum(256)
	blue := GenerateRandIntBelowMaximum(256)

	return fmt.Sprintf("#%x%x%x", red, green, blue)
}

func (c *Colors) Colorrgb() string {

	red := GenerateRandIntBelowMaximum(255)
	green := GenerateRandIntBelowMaximum(255)
	blue := GenerateRandIntBelowMaximum(255)

	return fmt.Sprintf("%s(%d,%d,%d)", c.ColorSpaces.rgb, red, green, blue)

}

func (c *Colors) Colorrgba() string {

	red := GenerateRandIntBelowMaximum(255)
	green := GenerateRandIntBelowMaximum(255)
	blue := GenerateRandIntBelowMaximum(255)

	alpha := RandomFloatBetweenTwoNumbers(0, 1)

	return fmt.Sprintf("%s(%d,%d,%d,%f)", c.ColorSpaces.rgb, red, green, blue, alpha)

}

func (c *Colors) Colorhsl() string {

	hue := GenerateRandIntBelowMaximum(359)
	saturation := GenerateRandIntBelowMaximum(101)
	lightness := GenerateRandIntBelowMaximum(101)

	return fmt.Sprintf("%s(%d,%d%%,%d%%)", c.ColorSpaces.hsl, hue, saturation, lightness)
}

func (c *Colors) Colorhsla() string {

	hue := GenerateRandIntBelowMaximum(359)
	saturation := GenerateRandIntBelowMaximum(101)
	lightness := GenerateRandIntBelowMaximum(101)
	alpha := RandomFloatBetweenTwoNumbers(0, 1)

	return fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", c.ColorSpaces.hsl, hue, saturation, lightness, alpha)
}
