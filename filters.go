package main

import "strings"

var (
	AllFileFilter  = "All File|*"
	PDFFilter      = "PDF Document|*.pdf"
	DocumentFilter = "Document|*.pdf;*.svg;*.xps;*.cbz;*.epub;*.mobi;*.fb2;*.html;*.htm"
	ImageFilter    = "Image|*.png;*.jpg;*.bmp;*.svg"
	ExeFilter      = "Executable File|*.exe"
)

func FilterOr(filter ...string) string {
	return strings.Join(filter, "|")
}
