package main

import "github.com/ying32/govcl/vcl/types"

var (
	BackgroundColor = RGB(0xFF, 0xFB, 0xF0)
	ControlColor    = RGB(0xFF, 0xEF, 0xC6)
	FontColor       = RGB(0x00, 0x00, 0x00)
	TipColor        = RGB(0xBE, 0xFF, 0x7D)
)

func RGB(r, g, b byte) types.TColor {
	return types.TColor(uint32(r) | (uint32(g) << 8) | (uint32(b) << 16))
}
