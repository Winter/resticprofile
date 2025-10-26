package hook

import (
	"github.com/creativeprojects/resticprofile/monitor"
	"github.com/creativeprojects/resticprofile/util/templates"
)

type Context struct {
	templates.DefaultData

	ProfileName    string
	ProfileCommand string
	Error          ErrorContext
	Stdout         string
	Summary        *monitor.Summary
}

type ErrorContext struct {
	Message     string
	CommandLine string
	ExitCode    string
	Stderr      string
}
