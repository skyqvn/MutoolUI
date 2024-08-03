package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type MultipleItems struct {
	*vcl.TPanel
	Items []*MultipleItem
	
	NewItem   func(owner vcl.IComponent) vcl.IWinControl
	ValueFunc func(item *MultipleItem) []string
}

type MultipleItem struct {
	*vcl.TPanel
	Item                    vcl.IWinControl
	AddButton, RemoveButton *vcl.TButton
	Index                   int
}

func NewMultipleItems(owner vcl.IComponent, newItem func(owner vcl.IComponent) vcl.IWinControl, value func(item *MultipleItem) []string) *MultipleItems {
	mi := new(MultipleItems)
	mi.TPanel = vcl.NewPanel(owner)
	mi.SetBevelOuter(types.BvNone)
	mi.NewItem = newItem
	mi.ValueFunc = value
	return mi
}

func (mi *MultipleItems) SetParent(value vcl.IWinControl) {
	mi.TPanel.SetParent(value)
}

func (mi *MultipleItems) Update() {
	if len(mi.Items) == 1 {
		p := mi.TPanel.Parent()
		if p != nil {
			p.SetHeight(ValueItemHeight)
		}
		mi.SetHeight(ValueItemHeight)
		mi.Items[0].Index = 0
		mi.Items[0].RemoveButton.SetEnabled(false)
		mi.Items[0].SetTop(0)
	} else {
		p := mi.TPanel.Parent()
		if p != nil {
			p.SetHeight(int32(ValueItemHeight * len(mi.Items)))
		}
		for i, item := range mi.Items {
			item.Index = i
			item.RemoveButton.SetEnabled(true)
			item.SetTop(int32(ValueItemHeight * i))
		}
	}
}

func (mi *MultipleItems) Append() {
	mi.Items = append(mi.Items, NewMultipleItem(mi))
	mi.Update()
}

func (mi *MultipleItems) Insert(at int) {
	mi.Items = Insert(mi.Items, at, NewMultipleItem(mi))
	mi.Update()
}

func (mi *MultipleItems) Delete(at int) {
	mi.Items[at].TPanel.Free()
	mi.Items = append(mi.Items[:at], mi.Items[at+1:]...)
	mi.Update()
}

func (mi *MultipleItems) Value() []string {
	arr := make([]string, 0, len(mi.Items))
	for _, item := range mi.Items {
		arr = append(arr, mi.ValueFunc(item)...)
	}
	return arr
}

func NewMultipleItem(mi *MultipleItems) *MultipleItem {
	m := new(MultipleItem)
	owner := mi.TPanel.Owner()
	m.TPanel = vcl.NewPanel(owner)
	m.TPanel.SetBevelOuter(types.BvNone)
	m.TPanel.SetAnchors(types.NewSet(types.AkLeft, types.AkRight, types.AkTop))
	m.TPanel.SetHeight(ValueItemHeight)
	m.TPanel.SetLeft(0)
	m.TPanel.SetWidth(mi.TPanel.Width())
	m.TPanel.SetParent(mi)
	m.AddButton = vcl.NewButton(owner)
	m.AddButton.SetCaption("+")
	m.AddButton.SetAnchors(types.NewSet(types.AkRight, types.AkTop))
	m.AddButton.SetWidth(ValueItemHeight)
	m.AddButton.SetHeight(ValueItemHeight)
	m.AddButton.SetTop(0)
	m.AddButton.SetLeft(m.TPanel.Width() - 2*ValueItemHeight)
	m.AddButton.SetOnClick(func(sender vcl.IObject) {
		mi.Insert(m.Index + 1)
	})
	m.RemoveButton = vcl.NewButton(owner)
	m.RemoveButton.SetCaption("-")
	m.RemoveButton.SetAnchors(types.NewSet(types.AkRight, types.AkTop))
	m.RemoveButton.SetWidth(ValueItemHeight)
	m.RemoveButton.SetHeight(ValueItemHeight)
	m.RemoveButton.SetTop(0)
	m.RemoveButton.SetLeft(m.TPanel.Width() - ValueItemHeight)
	m.RemoveButton.SetOnClick(func(sender vcl.IObject) {
		mi.Delete(m.Index)
	})
	m.AddButton.SetParent(m.TPanel)
	m.RemoveButton.SetParent(m.TPanel)
	m.Item = mi.NewItem(owner)
	m.Item.SetAnchors(types.NewSet(types.AkLeft, types.AkRight, types.AkTop))
	m.Item.SetLeft(0)
	m.Item.SetTop(0)
	m.Item.SetWidth(m.TPanel.Width() - 2*ValueItemHeight)
	m.Item.SetHeight(ValueItemHeight)
	m.Item.SetParent(m.TPanel)
	return m
}
