package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type FileEdit struct {
	*vcl.TPanel
	Button *vcl.TButton
	Edit   *vcl.TEdit
	
	Type FileEditType
	
	Filter string
}

type FileEditType int

const (
	Open FileEditType = iota
	Save
)

func NewFileEdit(owner vcl.IComponent, feType FileEditType) *FileEdit {
	fe := new(FileEdit)
	fe.Type = feType
	fe.TPanel = vcl.NewPanel(owner)
	fe.TPanel.SetBevelOuter(types.BvNone)
	fe.TPanel.SetHeight(ValueItemHeight)
	fe.Button = vcl.NewButton(owner)
	fe.Button.SetAlign(types.AlRight)
	fe.Button.SetCaption("Browse")
	fe.Button.SetOnClick(func(sender vcl.IObject) {
		fe.Brose()
	})
	fe.Button.SetParent(fe.TPanel)
	fe.Edit = vcl.NewEdit(owner)
	fe.Edit.SetParent(fe.TPanel)
	fe.Edit.SetAlign(types.AlClient)
	return fe
}

func (fe *FileEdit) SetParent(value vcl.IWinControl) {
	fe.TPanel.SetParent(value)
}

func (fe *FileEdit) Text() string {
	return fe.Edit.Text()
}

func (fe *FileEdit) Brose() {
	switch fe.Type {
	case Open:
		dlg := vcl.NewOpenDialog(MainForm)
		dlg.SetFilter(fe.Filter)
		dlg.Execute()
		fe.Edit.SetText(dlg.FileName())
	case Save:
		dlg := vcl.NewSaveDialog(MainForm)
		dlg.SetFilter(fe.Filter)
		dlg.Execute()
		fe.Edit.SetText(dlg.FileName())
	}
}
