package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	"github.com/dim13/hashcolor"
)

func hex(c color.Color) string {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}

func main() {
	c := hashcolor.New(strings.Join(os.Args[1:], " "))
	t, s := hashcolor.Tint(c), hashcolor.Shade(c)
	fmt.Println("-fg", hex(t), "-bg", hex(s))
}
