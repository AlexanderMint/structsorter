package structsorter_test

import (
	"testing"

	"github.com/AlexanderMint/structsorter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTest(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, structsorter.Analyzer, "a")
}
