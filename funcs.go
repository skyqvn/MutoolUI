package main

import (
	"fmt"
	"strconv"
)

func Command(page *Page) ([]string, bool) {
	args := make([]string, 0, 8)
	args = append(args, page.Command)
	var enableMainArg = false
	var mainArg string
	for _, item := range Items[page.Name] {
		if item.Type == Tip {
			continue
		}
		switch item.VType {
		case String:
			v, ok := item.Value()
			if !ok {
				return nil, false
			}
			s := v.(string)
			if item.IsNecessary && s == "" {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false
			}
			if s == "" {
				continue
			}
			if item.IsMainArg {
				enableMainArg = true
				mainArg = s
				continue
			}
			if item.Tag != "" {
				args = append(args, item.Tag)
			}
			args = append(args, s)
		case Int:
			v, ok := item.Value()
			if !ok {
				return nil, false
			}
			s := v.(string)
			if item.IsNecessary && s == "" {
				text := fmt.Sprintf("The %s fields is necessary", item.Name)
				PopupErrorDialog(text)
				return nil, false
			}
			if s == "" {
				continue
			}
			_, err := strconv.Atoi(s)
			if err != nil {
				text := fmt.Sprintf("%s must be a number", item.Name)
				PopupErrorDialog(text)
				return nil, false
			}
			if item.IsMainArg {
				enableMainArg = true
				mainArg = s
				continue
			}
			if item.Tag != "" {
				args = append(args, item.Tag)
			}
			args = append(args, s)
		case Bool:
			v, ok := item.Value()
			if !ok {
				return nil, false
			}
			b := v.(bool)
			if b {
				args = append(args, item.Tag)
			}
		case StringList:
			v, ok := item.Value()
			if !ok {
				return nil, false
			}
			args = append(args, v.([]string)...)
		}
	}
	if enableMainArg {
		args = append(args, mainArg)
	}
	return args, true
}

func ReverseSlice[T any](s []T) []T {
	l := len(s)
	s2 := make([]T, l)
	for i := 0; i < l; i++ {
		s2[l-1-i] = s[i]
	}
	return s2
}
