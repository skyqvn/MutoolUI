// 由res2go IDE插件自动生成，不要编辑。
package main

import (
	_ "embed"
	"github.com/ying32/govcl/vcl"
)

type TSettingForm struct {
	*vcl.TForm
	MainPageControl *vcl.TPageControl
	MutoolPage      *vcl.TTabSheet
	Panel1          *vcl.TPanel
	CancelButton    *vcl.TButton
	OKButton        *vcl.TButton
	
	// ::private::
	TSettingFormFields
}

var SettingForm *TSettingForm

// vcl.Application.CreateForm(&SettingForm)

func NewSettingForm(owner vcl.IComponent) (root *TSettingForm) {
	vcl.CreateResForm(owner, &root)
	return
}

//go:embed resources/settingunit.gfm
var settingFormBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(SettingForm, &settingFormBytes)
