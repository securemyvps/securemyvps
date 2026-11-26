package main

import (
	"flag"
	"fmt"
	"os"
)

// version is the release version. The release workflow replaces it at build
// time with -ldflags "-X main.version=<version>".
var version = "0.0.0-dev"

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	fmt.Fprintf(os.Stderr, "securemyvps %s (Phase 0 stub; remediation engine not wired yet)\n", version)
	os.Exit(0)
}
