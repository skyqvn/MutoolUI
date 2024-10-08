package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"os/exec"
)

type Page struct {
	Name    string
	Command string
	Page    *vcl.TTabSheet
}

var Pages []*Page

const SpaceHeight = 15
const ValueItemHeight = 30
const TipItemHeight = 20

type Item struct {
	Type        ItemType
	Name        string
	Label       string
	Control     vcl.IControl
	Value       func() (any, bool)
	IsNecessary bool
	Tag         string
	VType       ValueType
	IsMainArg   bool
}

type ItemType int

const (
	Value ItemType = iota
	Tip
)

type ValueType int

const (
	String ValueType = iota
	Int
	Bool
	StringList
	Path
)

var Items map[string][]*Item

var (
	DefaultFont *vcl.TFont
	BoldFont    *vcl.TFont
)

func InitUI() {
	MainForm.SetLeft((vcl.Screen.WorkAreaWidth() - MainForm.Width()) / 2)
	// MainForm.SetTop((vcl.Screen.WorkAreaHeight() - MainForm.Height()) / 2)
	MainForm.SetTop(30)
	DefaultFont = vcl.NewFont()
	DefaultFont.SetColor(FontColor)
	BoldFont = vcl.NewFont()
	BoldFont.SetStyle(types.NewSet(types.FsBold))
	BoldFont.SetColor(FontColor)
	MainForm.SetFont(DefaultFont)
	MainForm.SetColor(BackgroundColor)
	InitMenu()
	MainForm.MainPageControl = vcl.NewPageControl(MainForm)
	MainForm.MainPageControl.SetAlign(types.AlClient)
	MainForm.DrawPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.MergePage = MainForm.MainPageControl.AddTabSheet()
	MainForm.ConvertPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.PosterPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.ExtractPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.MainPageControl.SetParent(MainForm)
	MainForm.RunButton = vcl.NewButton(MainForm)
	MainForm.RunButton.SetCaption("Run")
	MainForm.RunButton.SetAlign(types.AlBottom)
	MainForm.RunButton.SetHeight(30)
	MainForm.RunButton.SetParent(MainForm)
	MainForm.RunButton.SetOnClick(func(sender vcl.IObject) {
		page := Pages[MainForm.MainPageControl.ActivePageIndex()]
		a, ok, dir := Command(page)
		if !ok {
			return
		}
		cmd := exec.Command(MutoolPath, a...)
		cmd.Dir = dir
		err := cmd.Start()
		if err != nil {
			PopupErrorDialog(err.Error())
			return
		}
		MainForm.RunButton.SetEnabled(false)
		MainForm.RunButton.SetCaption("Doing···")
		go func() {
			err = cmd.Wait()
			if err == nil {
				vcl.ThreadSync(func() {
					MainForm.RunButton.SetEnabled(true)
					MainForm.RunButton.SetCaption("Run")
					PopupInfoDialog("Done!")
				})
			} else {
				vcl.ThreadSync(func() {
					MainForm.RunButton.SetEnabled(true)
					MainForm.RunButton.SetCaption("Run")
					PopupErrorDialog(err.Error())
				})
			}
		}()
	})
	Pages = []*Page{
		{
			Name:    "Draw",
			Command: "draw",
			Page:    MainForm.DrawPage,
		},
		{
			Name:    "Merge",
			Command: "merge",
			Page:    MainForm.MergePage,
		},
		{
			Name:    "Convert",
			Command: "convert",
			Page:    MainForm.ConvertPage,
		},
		{
			Name:    "Poster",
			Command: "poster",
			Page:    MainForm.PosterPage,
		},
		{
			Name:    "Extract",
			Command: "extract",
			Page:    MainForm.ExtractPage,
		},
	}
	draw1 := NewFileEdit(MainForm, Open)
	draw1.Filter = DocumentFilter
	draw2 := NewFileEdit(MainForm, Save)
	draw2.Filter = ImageFilter
	draw2.FileName = "out%d.png"
	draw3 := vcl.NewEdit(MainForm)
	draw3.SetColor(ControlColor)
	draw3.SetNumbersOnly(true)
	draw4 := vcl.NewEdit(MainForm)
	draw4.SetPasswordChar('*')
	draw4.SetColor(ControlColor)
	merge1 := NewMultipleItems(MainForm, func(owner vcl.IComponent) vcl.IWinControl {
		fe := NewFileEdit(MainForm, Open)
		fe.Filter = DocumentFilter
		return NewPageSelector(MainForm, fe, func(item vcl.IWinControl) string {
			return item.(*FileEdit).Text()
		})
	}, func(item *MultipleItem) []string {
		return item.Item.(*PageSelector).Value()
	})
	merge1.SetColor(ControlColor)
	merge2 := NewFileEdit(MainForm, Save)
	merge2.Filter = DocumentFilter
	merge2.FileName = "out.pdf"
	merge3 := vcl.NewEdit(MainForm)
	merge3.SetPasswordChar('*')
	merge3.SetColor(ControlColor)
	convert1 := NewFileEdit(MainForm, Open)
	convert1.Filter = FilterOr(DocumentFilter, ImageFilter)
	convert2 := NewFileEdit(MainForm, Save)
	convert2.Filter = DocumentFilter
	convert2.FileName = "out.pdf"
	convert3 := vcl.NewEdit(MainForm)
	convert3.SetPasswordChar('*')
	convert3.SetColor(ControlColor)
	poster1 := NewFileEdit(MainForm, Open)
	poster1.Filter = DocumentFilter
	poster2 := NewFileEdit(MainForm, Save)
	poster2.Filter = DocumentFilter
	poster2.FileName = "out.pdf"
	poster3 := vcl.NewEdit(MainForm)
	poster3.SetColor(ControlColor)
	poster3.SetNumbersOnly(true)
	poster4 := vcl.NewEdit(MainForm)
	poster4.SetColor(ControlColor)
	poster4.SetNumbersOnly(true)
	poster5 := vcl.NewEdit(MainForm)
	poster5.SetPasswordChar('*')
	poster5.SetColor(ControlColor)
	extract1 := NewFileEdit(MainForm, Open)
	extract1.Filter = DocumentFilter
	extract2 := NewFileEdit(MainForm, SaveDir)
	extract3 := vcl.NewEdit(MainForm)
	extract3.SetPasswordChar('*')
	extract3.SetColor(ControlColor)
	Items = map[string][]*Item{
		"Draw": {
			{
				Type:    Value,
				Name:    "Source",
				Label:   "Source:",
				Control: draw1,
				Value: func() (any, bool) {
					return draw1.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       String,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Target",
				Label:   "Target:",
				Control: draw2,
				Value: func() (any, bool) {
					return draw2.Text(), true
				},
				IsNecessary: true,
				Tag:         "-o",
				VType:       String,
				IsMainArg:   false,
			},
			{
				Type:  Tip,
				Label: "Tip: Replace page number with %d",
			},
			{
				Type:    Value,
				Name:    "Resolution",
				Label:   "Resolution:",
				Control: draw3,
				Value: func() (any, bool) {
					return draw3.Text(), true
				},
				IsNecessary: false,
				Tag:         "-r",
				VType:       Int,
				IsMainArg:   false,
			},
			{
				Type:    Value,
				Name:    "Password",
				Label:   "Password:",
				Control: draw4,
				Value: func() (any, bool) {
					return draw4.Text(), true
				},
				IsNecessary: false,
				Tag:         "-p",
				VType:       String,
				IsMainArg:   false,
			},
		},
		"Merge": {
			{
				Type:    Value,
				Name:    "Source",
				Label:   "Source:",
				Control: merge1,
				Value: func() (any, bool) {
					return merge1.Value(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       StringList,
				IsMainArg:   true,
			},
			{
				Type:  Tip,
				Label: "Tip: Comma separated list of page ranges. The first page is “1”, and the last page is “N”. The default is “1-N”.",
			},
			{
				Type:  Tip,
				Label: "For example: \"1-3,5,7-N\"",
			},
			{
				Type:    Value,
				Name:    "Target",
				Label:   "Target:",
				Control: merge2,
				Value: func() (any, bool) {
					return merge2.Text(), true
				},
				IsNecessary: true,
				Tag:         "-o",
				VType:       String,
				IsMainArg:   false,
			},
			{
				Type:    Value,
				Name:    "Password",
				Label:   "Password:",
				Control: merge3,
				Value: func() (any, bool) {
					return merge3.Text(), true
				},
				IsNecessary: false,
				Tag:         "-p",
				VType:       String,
				IsMainArg:   false,
			},
		},
		"Convert": {
			{
				Type:    Value,
				Name:    "Source",
				Label:   "Source:",
				Control: convert1,
				Value: func() (any, bool) {
					return convert1.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       String,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Target",
				Label:   "Target:",
				Control: convert2,
				Value: func() (any, bool) {
					return convert2.Text(), true
				},
				IsNecessary: true,
				Tag:         "-o",
				VType:       String,
				IsMainArg:   false,
			},
			{
				Type:    Value,
				Name:    "Password",
				Label:   "Password:",
				Control: convert3,
				Value: func() (any, bool) {
					return convert3.Text(), true
				},
				IsNecessary: false,
				Tag:         "-p",
				VType:       String,
				IsMainArg:   false,
			},
		},
		"Poster": {
			{
				Type:    Value,
				Name:    "Source",
				Label:   "Source:",
				Control: poster1,
				Value: func() (any, bool) {
					return poster1.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       String,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Target",
				Label:   "Target:",
				Control: poster2,
				Value: func() (any, bool) {
					return poster2.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       String,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Divide X",
				Label:   "Divide X:",
				Control: poster3,
				Value: func() (any, bool) {
					return poster3.Text(), true
				},
				IsNecessary: true,
				Tag:         "-x",
				VType:       Int,
				IsMainArg:   false,
			},
			{
				Type:    Value,
				Name:    "Divide Y",
				Label:   "Divide Y:",
				Control: poster4,
				Value: func() (any, bool) {
					return poster4.Text(), true
				},
				IsNecessary: true,
				Tag:         "-y",
				VType:       Int,
				IsMainArg:   false,
			},
			{
				Type:    Value,
				Name:    "Password",
				Label:   "Password:",
				Control: poster5,
				Value: func() (any, bool) {
					return poster5.Text(), true
				},
				IsNecessary: false,
				Tag:         "-p",
				VType:       String,
				IsMainArg:   false,
			},
		},
		"Extract": {
			{
				Type:    Value,
				Name:    "Source",
				Label:   "Source:",
				Control: extract1,
				Value: func() (any, bool) {
					return extract1.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       String,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Target",
				Label:   "Target:",
				Control: extract2,
				Value: func() (any, bool) {
					return extract2.Text(), true
				},
				IsNecessary: true,
				Tag:         "",
				VType:       Path,
				IsMainArg:   true,
			},
			{
				Type:    Value,
				Name:    "Password",
				Label:   "Password:",
				Control: extract3,
				Value: func() (any, bool) {
					return extract3.Text(), true
				},
				IsNecessary: false,
				Tag:         "-p",
				VType:       String,
				IsMainArg:   false,
			},
		},
	}
	for _, page := range Pages {
		page.Page.SetCaption(page.Name)
		NewPage(page.Page, ReverseSlice(Items[page.Name]))
	}
}

func NewPage(parent vcl.IWinControl, table []*Item) {
	sp := vcl.NewScrollBox(MainForm)
	sp.SetAlign(types.AlClient)
	sp.SetColor(BackgroundColor)
	sp.SetParent(parent)
	p := vcl.NewPanel(MainForm)
	p.SetBevelOuter(types.BvNone)
	p.SetAlign(types.AlClient)
	p.SetColor(BackgroundColor)
	p.SetParent(sp)
	for i, item := range table {
		switch item.Type {
		case Value:
			if i != 0 && table[i-1].Type != Tip {
				NewSpace(p)
			}
			itemPanel := vcl.NewPanel(MainForm)
			itemPanel.SetHeight(ValueItemHeight)
			itemPanel.SetAlign(types.AlTop)
			itemPanel.SetBevelOuter(types.BvNone)
			itemPanel.SetColor(BackgroundColor)
			label := vcl.NewLabel(MainForm)
			label.SetParent(itemPanel)
			label.SetCaption(item.Label)
			if item.IsNecessary {
				label.SetFont(BoldFont)
			} else {
				label.SetFont(DefaultFont)
			}
			label.SetColor(BackgroundColor)
			label.SetAlign(types.AlLeft)
			item.Control.SetAlign(types.AlClient)
			item.Control.SetParent(itemPanel)
			m, ok := item.Control.(*MultipleItems)
			if ok {
				m.Append()
			}
			itemPanel.SetParent(p)
		case Tip:
			if i != 0 && table[i-1].Type != Tip {
				NewSpace(p)
			}
			panel := vcl.NewPanel(MainForm)
			panel.SetHeight(TipItemHeight)
			panel.SetAlign(types.AlTop)
			panel.SetBevelOuter(types.BvNone)
			panel.SetCaption(item.Label)
			panel.SetAlignment(types.TaLeftJustify)
			panel.SetColor(TipColor)
			panel.SetParent(p)
		}
	}
}

func NewSpace(parent vcl.IWinControl) {
	space := vcl.NewPanel(MainForm)
	space.SetBevelOuter(types.BvNone)
	space.SetAlign(types.AlTop)
	space.SetHeight(SpaceHeight)
	space.SetParent(parent)
}

func InitMenu() {
	MainForm.SettingAction = vcl.NewMenuItem(MainForm)
	MainForm.SettingAction.SetCaption("Setting")
	MainForm.SettingAction.SetOnClick(func(sender vcl.IObject) {
		sf := NewSettingForm(MainForm)
		sf.ActiveEdit = NewFileEdit(sf, Open)
		sf.ActiveEdit.SetAlign(types.AlClient)
		sf.ActiveEdit.Filter = ExeFilter
		sf.ActiveEdit.FileName = "mutool.exe"
		sf.ActiveEdit.Edit.SetColor(colors.ClDefault)
		sf.ActiveEdit.Edit.SetText(MutoolPath)
		sf.ActiveEdit.SetParent(sf.ActiveEditPanel)
		sf.AllListBox.Items().AddStrings2(MutoolAll)
		if sf.AllListBox.Items().Count() != 0 {
			sf.AllListBox.SetItemIndex(0)
		}
		sf.SetAsActiveButton.SetOnClick(func(sender vcl.IObject) {
			if i := sf.AllListBox.ItemIndex(); i != -1 {
				sf.ActiveEdit.SetText(sf.AllListBox.Items().S(i))
			}
		})
		sf.AddButton.SetOnClick(func(sender vcl.IObject) {
			PopupEditTextDialog(sf, "", func(s string) {
				sf.AllListBox.Items().Add(s)
				sf.AllListBox.SetItemIndex(sf.AllListBox.Items().Count() - 1)
			})
		})
		sf.EditButton.SetOnClick(func(sender vcl.IObject) {
			PopupEditTextDialog(sf, sf.AllListBox.Items().S(sf.AllListBox.ItemIndex()), func(s string) {
				sf.AllListBox.Items().SetS(sf.AllListBox.ItemIndex(), s)
			})
		})
		sf.DeleteButton.SetOnClick(func(sender vcl.IObject) {
			sf.AllListBox.Items().Delete(sf.AllListBox.ItemIndex())
			if sf.AllListBox.Items().Count() != 0 {
				sf.AllListBox.SetItemIndex(0)
			}
		})
		sf.OKButton.SetOnClick(func(sender vcl.IObject) {
			MutoolPath = sf.ActiveEdit.Text()
			its := sf.AllListBox.Items()
			if its.IndexOf(MutoolPath) == -1 {
				its.Insert(0, MutoolPath)
			}
			MutoolAll = StringsToSlice(its)
			UpdateMutoolSetting()
			sf.Close()
		})
		sf.CancelButton.SetOnClick(func(sender vcl.IObject) {
			sf.Close()
		})
		sf.ShowModal()
	})
	MainForm.MainMenu.Items().Add(MainForm.SettingAction)
	
	MainForm.HelpAction = vcl.NewMenuItem(MainForm)
	MainForm.HelpAction.SetCaption("Help")
	MainForm.HelpAction.SetOnClick(func(sender vcl.IObject) {
		if err := OpenURI("./docs/help/contents.html"); err != nil {
			PopupErrorDialog(err.Error() + "\nUnable to open the document, please open ./docs/help/contents.html manually.")
		}
	})
	MainForm.MainMenu.Items().Add(MainForm.HelpAction)
}
