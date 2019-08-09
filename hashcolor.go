package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/dim13/crc24"
)

func toColor(c uint32) color.RGBA {
	return color.RGBA{
		R: uint8(c >> 0x10),
		G: uint8(c >> 0x08),
		B: uint8(c),
		A: 0xff,
	}
}

func hex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func tint(c color.Color) color.RGBA {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	rgba.R += (0xff - rgba.R) >> 1
	rgba.G += (0xff - rgba.G) >> 1
	rgba.B += (0xff - rgba.B) >> 1
	return rgba
}

func shade(c color.Color) color.RGBA {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	rgba.R >>= 2
	rgba.G >>= 2
	rgba.B >>= 2
	return rgba
}

func main() {
	arg := strings.Join(os.Args[1:], " ")
	c := toColor(crc24.Sum([]byte(arg)))
	fmt.Println("-fg", hex(tint(c)), "-bg", hex(shade(c)))
}
