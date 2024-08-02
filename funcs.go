package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"os/exec"
	"runtime"
	"strconv"
)

func Command(page *Page) ([]string, bool, string) {
	args := make([]string, 0, 8)
	args = append(args, page.Command)
	var enableMainArg = false
	var mainArg []string
	var dir string
	for _, item := range Items[page.Name] {
		if item.Type == Tip {
			continue
		}
		switch item.VType {
		case String:
			v, ok := item.Value()
			if !ok {
				return nil, false, ""
			}
			s := v.(string)
			if item.IsNecessary && s == "" {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false, ""
			}
			if s == "" {
				continue
			}
			if item.IsMainArg {
				enableMainArg = true
				mainArg = append(mainArg, s)
				continue
			}
			if item.Tag != "" {
				args = append(args, item.Tag)
			}
			args = append(args, s)
		case Int:
			v, ok := item.Value()
			if !ok {
				return nil, false, ""
			}
			s := v.(string)
			if item.IsNecessary && s == "" {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false, ""
			}
			if s == "" {
				continue
			}
			_, err := strconv.Atoi(s)
			if err != nil {
				text := fmt.Sprintf("%s must be a number", item.Name)
				PopupErrorDialog(text)
				return nil, false, ""
			}
			if item.IsMainArg {
				enableMainArg = true
				mainArg = append(mainArg, s)
				continue
			}
			if item.Tag != "" {
				args = append(args, item.Tag)
			}
			args = append(args, s)
		case Bool:
			v, ok := item.Value()
			if !ok {
				return nil, false, ""
			}
			b := v.(bool)
			if b {
				args = append(args, item.Tag)
			}
		case StringList:
			v, ok := item.Value()
			if !ok {
				return nil, false, ""
			}
			s := v.([]string)
			if item.IsNecessary && IsEmpty(s) {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false, ""
			}
			if item.IsMainArg {
				enableMainArg = true
				mainArg = append(mainArg, v.([]string)...)
				continue
			}
			args = append(args, s...)
		case Path:
			v, ok := item.Value()
			if !ok {
				return nil, false, ""
			}
			dir = v.(string)
			if item.IsNecessary && dir == "" {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false, ""
			}
		}
	}
	if enableMainArg {
		args = append(args, mainArg...)
	}
	return args, true, dir
}

func ReverseSlice[T any](s []T) []T {
	l := len(s)
	s2 := make([]T, l)
	for i := 0; i < l; i++ {
		s2[l-1-i] = s[i]
	}
	return s2
}

func StringsToSlice(strings *vcl.TStrings) []string {
	c := strings.Count()
	s := make([]string, c)
	var i int32
	for ; i < c; i++ {
		s[i] = strings.S(i)
	}
	return s
}

func IsIn[T comparable](v T, s []T) bool {
	for _, v2 := range s {
		if v2 == v {
			return true
		}
	}
	return false
}

func IsEmpty(v []string) bool {
	if len(v) == 0 {
		return true
	}
	for _, s := range v {
		if s != "" {
			return false
		}
	}
	return true
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenURI(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}
	
	cmd := exec.Command(run, uri)
	return cmd.Start()
}
