package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

// create_mod <modpath> <version> <srcdir> <outzip>
// Packages srcdir as a Go module zip (proxy.golang.org format) for the given
// module path + version, using golang.org/x/mod/zip.CreateFromDir.
func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: create_mod <modpath> <version> <srcdir> <outzip>\n")
		os.Exit(2)
	}
	modpath, version, srcdir, outzip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(outzip)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create out zip: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modpath, Version: version}
	if err := zip.CreateFromDir(f, m, srcdir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s for %s@%s\n", outzip, modpath, version)
}
