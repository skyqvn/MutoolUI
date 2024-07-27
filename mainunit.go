// 由res2go IDE插件自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainForm struct {
    *vcl.TForm
    MainMenu *vcl.TMainMenu

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// vcl.Application.CreateForm(&MainForm)

func NewMainForm(owner vcl.IComponent) (root *TMainForm)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/mainunit.gfm
var mainFormBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
