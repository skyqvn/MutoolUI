package main

import "strings"

var (
	AllFile        = "All File|*"
	PDFFilter      = "PDF Document|*.pdf"
	DocumentFilter = "Document|*.pdf;*.svg;*.xps;*.cbz;*.epub;*.mobi;*.fb2"
	ImageFilter    = "Image|*.png;*.jpg;*.bmp;*.svg"
	ExeFilter      = "Executable File|*.exe"
)

func FilterOr(filter ...string) string {
	return strings.Join(filter, "|")
}
