package main

import (
	"flag"
	"fmt"
	"os"
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

var showVersion = flag.Bool("version", false, "show tool version")

func init() {
	if info, available := debug.ReadBuildInfo(); available && commit == unknown {
		version = info.Main.Version
		if info.Main.Sum != "" {
			commit = fmt.Sprintf("%s, mod sum: %s", commit, info.Main.Sum)
		}
	}
}

func hack() {
	if *showVersion {
		_, _ = fmt.Printf(message, version, date, commit)
		os.Exit(0)
	}
}
