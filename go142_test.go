package go142_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"go142"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, go142.Analyzer, "a")
}
