package main

import (
	"log"
	"os"

	"github.com/andreas-jonsson/ssa-test/packages"
	"github.com/andreas-jonsson/ssa-test/ssa"
	"github.com/andreas-jonsson/ssa-test/ssa/ssautil"
)

func main() {
	// Load, parse, and type-check the whole program.
	cfg := packages.Config{Mode: packages.LoadAllSyntax}
	initial, err := packages.Load(&cfg, "github.com/andreas-jonsson/ssa-test")
	if err != nil {
		log.Fatal(err)
	}

	// Create SSA packages for well-typed packages and their dependencies.
	prog, _ := ssautil.AllPackages(initial, ssa.SanityCheckFunctions)

	// Build SSA code for the whole program.
	prog.Build()

	for _, pkg := range prog.AllPackages() {
		for _, m := range pkg.Members {
			if f, ok := m.(*ssa.Function); ok {
				f.WriteTo(os.Stdout)
			}
		}
	}
}
