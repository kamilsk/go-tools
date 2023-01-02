package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

const (
	unknown = "unknown"
	message = "%s, build date: %s, git commit: %s\n"
)

var (
	commit  = unknown
	date    = unknown
	version = "dev"
)

func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		if info.Main.Sum != "" {
			commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
		}
	}
	flag.Func("version", "show tool version", func(string) error {
		_, _ = fmt.Printf(message, version, date, commit)
		return flag.ErrHelp
	})
}
