package sortslice_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/johanbrandhorst/sortslicevet"
) 

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, sortslice.Analyzer, "a")
}
