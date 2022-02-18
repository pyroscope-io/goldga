package goldga

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/reporters"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	var specReporters []Reporter

	if dir := os.Getenv("JUNIT_OUTPUT"); dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}

		path := filepath.Join(dir, fmt.Sprintf("junit-%d.xml", time.Now().UnixNano()))
		specReporters = append(specReporters, reporters.NewJUnitReporter(path))
	}

	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "goldga", specReporters)
}
