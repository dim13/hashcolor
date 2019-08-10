package hashcolor

import (
	"image/color"

	"github.com/dim13/crc24"
)

// New color from hashed string
func New(s string) color.RGBA {
	c := crc24.Sum([]byte(s))
	return color.RGBA{
		R: uint8(c >> 0x10),
		G: uint8(c >> 0x08),
		B: uint8(c),
		A: 0xff,
	}
}

// Tint color
func Tint(c color.Color) color.RGBA {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	rgba.R += (0xff - rgba.R) >> 1
	rgba.G += (0xff - rgba.G) >> 1
	rgba.B += (0xff - rgba.B) >> 1
	return rgba
}

// Shade color
func Shade(c color.Color) color.RGBA {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	rgba.R >>= 2
	rgba.G >>= 2
	rgba.B >>= 2
	return rgba
}
