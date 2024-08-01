// 由res2go IDE插件自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TEditTextDialog struct {
    *vcl.TForm
    TextEdit     *vcl.TEdit
    Label1       *vcl.TLabel
    Panel1       *vcl.TPanel
    CancelButton *vcl.TButton
    OKButton     *vcl.TButton

    //::private::
    TEditTextDialogFields
}

var EditTextDialog *TEditTextDialog




// vcl.Application.CreateForm(&EditTextDialog)

func NewEditTextDialog(owner vcl.IComponent) (root *TEditTextDialog)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/edittextunit.gfm
var editTextDialogBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(EditTextDialog, &editTextDialogBytes)
