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
	
	Filter   string
	FileName string
	Title    string
}

type FileEditType int

const (
	Open FileEditType = iota
	Save
	SaveDir
)

func NewFileEdit(owner vcl.IComponent, feType FileEditType) *FileEdit {
	fe := new(FileEdit)
	fe.Type = feType
	fe.TPanel = vcl.NewPanel(owner)
	fe.TPanel.SetBevelOuter(types.BvNone)
	fe.TPanel.SetHeight(ValueItemHeight)
	fe.SetColor(ControlColor)
	fe.Button = vcl.NewButton(owner)
	fe.Button.SetAlign(types.AlRight)
	fe.Button.SetCaption("Browse")
	fe.Button.SetOnClick(func(sender vcl.IObject) {
		fe.Brose()
	})
	// fe.Button.SetColor(BackgroundColor)
	fe.Button.SetParent(fe.TPanel)
	fe.Edit = vcl.NewEdit(owner)
	fe.Edit.SetParent(fe.TPanel)
	fe.Edit.SetAlign(types.AlClient)
	fe.Edit.SetColor(ControlColor)
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
		dlg.SetFileName(fe.FileName)
		if fe.Title == "" {
			dlg.SetTitle("Open existing file")
		} else {
			dlg.SetTitle(fe.Title)
		}
		dlg.Execute()
		fn := dlg.FileName()
		if fn != "" {
			fe.Edit.SetText(fn)
		}
	case Save:
		dlg := vcl.NewSaveDialog(MainForm)
		dlg.SetFilter(fe.Filter)
		dlg.SetFileName(fe.FileName)
		if fe.Title == "" {
			dlg.SetTitle("Save file as")
		} else {
			dlg.SetTitle(fe.Title)
		}
		dlg.Execute()
		fn := dlg.FileName()
		if fn != "" {
			fe.Edit.SetText(fn)
		}
	case SaveDir:
		dlg := vcl.NewSelectDirectoryDialog(MainForm)
		dlg.SetFilter(fe.Filter)
		dlg.SetFileName(fe.FileName)
		if fe.Title == "" {
			dlg.SetTitle("Select directory")
		} else {
			dlg.SetTitle(fe.Title)
		}
		dlg.Execute()
		fn := dlg.FileName()
		if fn != "" {
			fe.Edit.SetText(fn)
		}
	}
}
