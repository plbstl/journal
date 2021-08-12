package cmd

import "github.com/fatih/color"

const (
	Author = "Author"
	Note   = "Note"
	Tag    = "Tag"
)

var (
	cyan = color.New(color.FgCyan).SprintFunc()
	red  = color.New(color.FgRed).SprintFunc()
)

var authorsDir, cfgFile, diaryConfigDir, notesDir, tagsDir string
