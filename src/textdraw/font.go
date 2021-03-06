/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package textdraw

import (
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/sdl"
	"image/color"
)

type Font struct {
	Font  *ttf.Font
	Color color.RGBA
}

func NewTTFFont(fontFile string, fontSize int, color color.RGBA) (*Font, error) {
	f, err := ttf.OpenFont(fontFile, fontSize)
	return &Font{Font: f, Color: color}, err
}

func (f *Font) CloseTTFFont() {
	f.Font.Close()
	*f = Font{}
}

func (f *Font) SetColor(color color.RGBA) {
	f.Color = color
}

func (f *Font) GetColor() color.RGBA {
	return f.Color
}

func (f *Font) GetSDLColor() sdl.Color {
	return sdl.Color{R: f.Color.R, G: f.Color.G, B: f.Color.B, A: f.Color.A}
}

func (f *Font) GetLineWidth(text string) int32 {
	width,_,_ := f.Font.SizeUTF8(text)
	return int32(width)
}
