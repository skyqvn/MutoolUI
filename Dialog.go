package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func PopupInfoDialog(text string) {
	vcl.MessageDlg(text, types.MtInformation, types.MbOK)
}

func PopupErrorDialog(text string) {
	vcl.MessageDlg(text, types.MtError, types.MbOK)
}
