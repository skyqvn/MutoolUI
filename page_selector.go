package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type PageSelector struct {
	*vcl.TPanel
	Item      vcl.IWinControl
	Pages     *vcl.TEdit
	ValueFunc func(item vcl.IWinControl) string
}

func NewPageSelector(owner vcl.IComponent, item vcl.IWinControl, value func(item vcl.IWinControl) string) *PageSelector {
	ps := new(PageSelector)
	ps.TPanel = vcl.NewPanel(MainForm)
	ps.TPanel.SetBevelOuter(types.BvNone)
	ps.TPanel.SetHeight(ValueItemHeight)
	ps.Pages = vcl.NewEdit(owner)
	ps.Pages.SetWidth(80)
	ps.Pages.SetAlign(types.AlRight)
	ps.Pages.SetParent(ps.TPanel)
	ps.Item = item
	ps.Item.SetAlign(types.AlClient)
	ps.Item.SetParent(ps.TPanel)
	ps.ValueFunc = value
	return ps
}

func (ps *PageSelector) SetParent(value vcl.IWinControl) {
	ps.TPanel.SetParent(value)
}

func (ps *PageSelector) Value() []string {
	return []string{ps.ValueFunc(ps.Item), ps.Pages.Text()}
}
