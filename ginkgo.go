package goldga

import (
	"path/filepath"
	"strings"

	"github.com/onsi/ginkgo/v2"
)

func getCurrentGinkgoTestDescription() ginkgo.SpecReport {
	return ginkgo.CurrentSpecReport()
}

func getGinkgoPath() string {
	desc := getCurrentGinkgoTestDescription()
	path := desc.FileName()

	if path == "" {
		panic("current file name is empty")
	}

	name := filepath.Base(path)

	if ext := filepath.Ext(name); ext != "" {
		name = strings.TrimSuffix(name, ext)
		name = strings.TrimSuffix(name, "_test")
	}

	return filepath.Join("testdata", name+".golden")
}

func getGinkgoTestName() string {
	testName := getCurrentGinkgoTestDescription().FullText()

	if testName == "" {
		panic("current test name is empty")
	}

	return testName
}
