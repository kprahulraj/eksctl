package builder_test

import (
	"testing"

	"github.com/weaveworks/eksctl/pkg/testutils"
)

func TestCFNBuilder(t *testing.T) {
	testutils.RegisterAndRun(t, "API Suite")
}
