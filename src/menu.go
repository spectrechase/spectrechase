/**********************************************
 *  Copyright 2018 The Spectre Chase Project  *
 **********************************************
 *   This software is made freely available   *
 *   under the MIT license. See LICENSE.txt   *
 *     for the full terms of this license.    *
 **********************************************
 */

package main

import (
	"fmt"
	//"github.com/veandco/go-sdl2/sdl"
	"image/color"
	"textdraw"
	"strings"
)

var textColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type MenuItem struct {
	number   int32
	text     string
	selected bool
}

type Menu []*MenuItem

func initializeMenu() Menu {
	return Menu{
		&MenuItem{number: 0, text: "Start Game", selected: true},
		&MenuItem{number: 1, text: "About", selected: false},
		&MenuItem{number: 2, text: "Exit", selected: false},
	}
}

func drawLogo(w *textdraw.Window, x, y int32) error {
	logoTextLines := []string{"Spectre", "Chase"}
	logoTextColor := color.RGBA{R: 250, G: 206, B: 52, A: 255}
	logoTextOutlineColor := color.RGBA{R: 175, G: 0, B: 0, A: 255}
	//logoTextVerticalPos := []int32{100, 200}

	logoFont, err := textdraw.NewTTFFont(TitleFontAddress, 100, logoTextColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer logoFont.CloseTTFFont()

	width, height, _ := logoFont.Font.SizeUTF8(logoTextLines[0])
	for i, text := range logoTextLines {
		logoFont.Font.SetOutline(0)
		logoFont.Color = logoTextColor
		w.DrawSizedText(text, logoFont, x, y+int32(i)*80, int32(width), int32(height))
		logoFont.Color = logoTextOutlineColor
		logoFont.Font.SetOutline(2)
		w.DrawSizedText(text, logoFont, x-4, y-4+int32(i)*80, int32(width)+4, int32(height)+4)
	}

	return nil
}

func renderMenuScreen(w *textdraw.Window) error {
	menuFont, err := textdraw.NewTTFFont(MenuFontAddress, 20, textColor)
	if err != nil {
		return fmt.Errorf("Error: could not cretae new font - %v", err)
	}
	defer menuFont.CloseTTFFont()
	w.Renderer.Clear()
	drawLogo(w, (Width-287)/2, -10)

	menu := initializeMenu()
	for _, item := range menu {
		if item.selected {
			menuFont.SetColor(color.RGBA{R: 240, G: 255, B: 98, A: 255})
			displayText := strings.Join([]string{">",item.text},"")
			w.DrawText(displayText,
				menuFont,
				(Width-menuFont.GetLineWidth(item.text))/2-20, 230+item.number*40)
		} else {
			menuFont.SetColor(textColor)
			w.DrawText(item.text,
				menuFont,
				(Width-menuFont.GetLineWidth(item.text))/2, 230+item.number*40)
		}
	}
	w.Renderer.Present()
	return nil
}
