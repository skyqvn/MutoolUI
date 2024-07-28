package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"os/exec"
)

// ::private::
type TMainFormFields struct {
	MainPageControl                  *vcl.TPageControl
	DrawPage, MergePage, ConvertPage *vcl.TTabSheet
	RunButton                        *vcl.TButton
}

type Page struct {
	Name    string
	Command string
	Page    *vcl.TTabSheet
}

// var MutoolPath string // TODO
var MutoolPath = "mutool"

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
)

var Items map[string][]*Item

var (
	DefaultFont *vcl.TFont
	BoldFont    *vcl.TFont
)

func InitUI() {
	DefaultFont = vcl.NewFont()
	DefaultFont.SetColor(FontColor)
	BoldFont = vcl.NewFont()
	BoldFont.SetStyle(types.NewSet(types.FsBold))
	BoldFont.SetColor(FontColor)
	MainForm.SetFont(DefaultFont)
	MainForm.SetColor(BackgroundColor)
	MainForm.MainPageControl = vcl.NewPageControl(MainForm)
	MainForm.MainPageControl.SetAlign(types.AlClient)
	MainForm.DrawPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.MergePage = MainForm.MainPageControl.AddTabSheet()
	MainForm.ConvertPage = MainForm.MainPageControl.AddTabSheet()
	MainForm.MainPageControl.SetParent(MainForm)
	MainForm.RunButton = vcl.NewButton(MainForm)
	MainForm.RunButton.SetCaption("Run")
	MainForm.RunButton.SetAlign(types.AlBottom)
	MainForm.RunButton.SetHeight(30)
	MainForm.RunButton.SetParent(MainForm)
	MainForm.RunButton.SetOnClick(func(sender vcl.IObject) {
		page := Pages[MainForm.MainPageControl.ActivePageIndex()]
		a, ok := Command(page)
		if !ok {
			return
		}
		cmd := exec.Command(MutoolPath, a...)
		err := cmd.Start()
		if err != nil {
			PopupErrorDialog(err.Error())
		}
		go func() {
			err = cmd.Wait()
			if err == nil {
				vcl.ThreadSync(func() {
					PopupInfoDialog("Done!")
				})
			} else {
				vcl.ThreadSync(func() {
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
		// {
		// 	Name:    "Merge",
		// 	Command: "merge",
		// 	Page:    MainForm.MergePage,
		// },
		// {
		// 	Name: "Convert",
		// 	Page: MainForm.ConvertPage,
		// },
	}
	draw1 := NewFileEdit(MainForm, Open)
	draw1.Filter = DocumentFilter
	draw2 := NewFileEdit(MainForm, Save)
	draw2.Filter = ImageFilter
	draw3 := vcl.NewEdit(MainForm)
	draw3.SetColor(ControlColor)
	draw3.SetNumbersOnly(true)
	draw4 := NewMultipleItems(MainForm, func(owner vcl.IComponent) vcl.IWinControl {
		return NewFileEdit(MainForm, Open)
	}, func(item *MultipleItem) string {
		return item.Item.(*FileEdit).Text()
	})
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
				Type:  Tip,
				Label: "Tip: Replace page number with %d",
			},
			{
				Type:    Value,
				Name:    "Resolution",
				Label:   "Resolution:",
				Control: draw4,
				Value: func() (any, bool) {
					return draw4.Value(), true
				},
				IsNecessary: false,
				Tag:         "-r",
				VType:       StringList,
				IsMainArg:   false,
			},
			// {
			// 	Name:    "Vads",
			// 	Label:   "Targeddssft:",
			// 	Control: merge3,
			// 	Value: func() (any, bool) {
			// 		i, err := strconv.Atoi(merge3.Text())
			// 		if err != nil {
			// 			PopupErrorDialog("XXX must be a number")
			// 		}
			// 		return i, err == nil
			// 	},
			// 	IsNecessary: false,
			// 	Tag:         "-f",
			// 	VType:        Int,
			// 	IsMainArg:   false,
			// },
		},
		// "Merge": {
		// 	{
		// 		Name:    "Source",
		// 		Label:   "Source:",
		// 		Control: merge1,
		// 		Value: func() (any, bool) {
		// 			return merge1.Text(), true
		// 		},
		// 		IsNecessary: true,
		// 		Tag:         "",
		// 		VType:        String,
		// 		IsMainArg:   true,
		// 	},
		// 	{
		// 		Name:    "Target",
		// 		Label:   "Target:",
		// 		Control: merge2,
		// 		Value: func() (any, bool) {
		// 			return merge2.Text(), true
		// 		},
		// 		IsNecessary: true,
		// 		Tag:         "-o",
		// 		VType:        String,
		// 		IsMainArg:   false,
		// 	},
		// "Convert": {
		// 	{
		// 		Label: "File:",
		// 		Value: vcl.NewEdit(MainForm),
		// 	},
		// 	{
		// 		Label: "Output:",
		// 		Value: vcl.NewEdit(MainForm),
		// 	},
		// },
	}
	for _, page := range Pages {
		page.Page.SetCaption(page.Name)
		NewPage(page.Page, ReverseSlice(Items[page.Name]))
	}
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	InitUI()
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
			if i != 0 {
				NewSpace(p)
			}
			panel := vcl.NewPanel(MainForm)
			panel.SetHeight(TipItemHeight)
			panel.SetAlign(types.AlTop)
			panel.SetBevelOuter(types.BvNone)
			panel.SetCaption(item.Label)
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
