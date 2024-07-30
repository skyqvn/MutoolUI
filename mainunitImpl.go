package main

import (
	"github.com/ying32/govcl/vcl"
)

// ::private::
type TMainFormFields struct {
	MainPageControl                                           *vcl.TPageControl
	DrawPage, MergePage, ConvertPage, PosterPage, ExtractPage *vcl.TTabSheet
	RunButton                                                 *vcl.TButton
	SettingAction                                             *vcl.TMenuItem
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	InitUI()
}
