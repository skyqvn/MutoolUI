package main

import (
	"github.com/ying32/govcl/vcl"
	"strings"
)

var SettingFile *vcl.TIniFile

var MutoolPath = "mutool"
var MutoolAll []string

func init() {
	SettingFile = vcl.NewIniFile("config.ini")
	if SettingFile.SectionExists("Mutool") {
		MutoolPath = SettingFile.ReadString("Mutool", "Path", "mutool")
		MutoolAll = strings.Split(SettingFile.ReadString("Mutool", "All", "mutool"), ";")
	} else {
		MutoolPath = "mutool"
		SettingFile.WriteString("Mutool", "Path", "mutool")
		MutoolAll = []string{"mutool"}
		SettingFile.WriteString("Mutool", "All", "mutool")
	}
}
