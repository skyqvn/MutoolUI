// 由res2go IDE插件自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("MuToolUI")
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&MainForm)
    vcl.Application.Run()
}
