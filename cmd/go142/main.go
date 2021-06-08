package main

import (
	"go142"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(go142.Analyzer) }
