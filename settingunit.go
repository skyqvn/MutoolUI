// 由res2go IDE插件自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TSettingForm struct {
    *vcl.TForm
    MainPageControl   *vcl.TPageControl
    MutoolPage        *vcl.TTabSheet
    Panel2            *vcl.TPanel
    Label1            *vcl.TLabel
    ActiveEditPanel   *vcl.TPanel
    Panel3            *vcl.TPanel
    Label2            *vcl.TLabel
    AllListBox        *vcl.TListBox
    Panel4            *vcl.TPanel
    SetAsActiveButton *vcl.TButton
    AddButton         *vcl.TButton
    EditButton        *vcl.TButton
    DeleteButton      *vcl.TButton
    Panel1            *vcl.TPanel
    CancelButton      *vcl.TButton
    OKButton          *vcl.TButton

    //::private::
    TSettingFormFields
}

var SettingForm *TSettingForm




// vcl.Application.CreateForm(&SettingForm)

func NewSettingForm(owner vcl.IComponent) (root *TSettingForm)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/settingunit.gfm
var settingFormBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(SettingForm, &settingFormBytes)
