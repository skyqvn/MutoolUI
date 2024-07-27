package main

import "strings"

var (
	PDFFilter      = "PDF Document|*.pdf"
	DocumentFilter = "Document|*.pdf;*.svg;*.xps;*.cbz;*.epub;*.mobi;*.fb2"
	ImageFilter    = "Image|*.png;*.jpg;*.bmp;*.svg"
)

func FilterOr(filter ...string) string {
	return strings.Join(filter, "|")
}
