package main

import "github.com/ying32/govcl/vcl"

var SettingFile *vcl.TIniFile

var MutoolPath = "mutool"

// var TempDir = "./tmp"

func init() {
	SettingFile = vcl.NewIniFile("config.ini")
	if SettingFile.SectionExists("Mutool") {
		MutoolPath = SettingFile.ReadString("Mutool", "Path", "mutool")
		// TempDir = SettingFile.ReadString("Mutool", "TempDir", "./tmp")
	} else {
		MutoolPath = "mutool"
		SettingFile.WriteString("Mutool", "Path", "mutool")
		// TempDir = "./tmp"
		// SettingFile.WriteString("Mutool", "TempDir", "./tmp")
	}
}
