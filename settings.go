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
		MutoolPath = SettingFile.ReadString("Mutool", "Active", "mutool")
		MutoolAll = strings.Split(SettingFile.ReadString("Mutool", "All", "mutool"), ":")
	} else {
		MutoolPath = "mutool"
		SettingFile.WriteString("Mutool", "Active", "mutool")
		MutoolAll = []string{"mutool"}
		SettingFile.WriteString("Mutool", "All", "mutool")
	}
}

func UpdateMutoolSetting() {
	SettingFile.WriteString("Mutool", "Active", MutoolPath)
	SettingFile.WriteString("Mutool", "All", strings.Join(MutoolAll, ":"))
}
