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

func PopupEditTextDialog(owner vcl.IComponent, text string, callback func(string)) {
	dlg := NewEditTextDialog(owner)
	dlg.TextEdit.SetText(text)
	dlg.OKButton.SetOnClick(func(sender vcl.IObject) {
		callback(dlg.TextEdit.Text())
		dlg.Close()
	})
	dlg.CancelButton.SetOnClick(func(sender vcl.IObject) {
		dlg.Close()
	})
	dlg.ShowModal()
}
