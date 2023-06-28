package archTest

import (
	"fmt"
	"strings"
	"testing"
)

func TestPackage_ShouldNotDependOn(t *testing.T) {

	t.Run("Succeeds on non dependencies", func(t *testing.T) {
		mockT := new(testingT)
		Package(t, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/nodependency")
		assertNoError(t, mockT)
	})

	t.Run("Fails on dependencies", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/dependency")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage",
			"github.com/ANDERSON1808/archTest/tradicionales/dependency")
	})

	t.Run("Supports testing against packages in the go root", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependOn("crypto")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage",
			"crypto")
	})

	t.Run("Fails on transative dependencies", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/transative")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage",
			"github.com/ANDERSON1808/archTest/tradicionales/dependency",
			"github.com/ANDERSON1808/archTest/tradicionales/transative")
	})

	t.Run("Supports wildcard matching", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/...").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/nodependency")

		assertNoError(t, mockT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/...").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/...")

		assertError(t, mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/dep", "github.com/ANDERSON1808/archTest/tradicionales/nesteddependency")
	})

	t.Run("Supports checking imports in test files", func(t *testing.T) {
		mockT := new(testingT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/...").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testonlydependency")

		assertNoError(t, mockT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/...").
			IncludeTests().
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testonlydependency")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/dep",
			"github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testonlydependency",
		)
	})

	t.Run("Supports checking imports from test packages", func(t *testing.T) {
		mockT := new(testingT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/...").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testpkgdependency")

		assertNoError(t, mockT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/...").
			IncludeTests().
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testpkgdependency")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/dep_test",
			"github.com/ANDERSON1808/archTest/tradicionales/testfiledeps/testpkgdependency",
		)
	})

	t.Run("Supports Ignoring packages", func(t *testing.T) {
		mockT := new(testingT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/dep").
			Ignoring("github.com/ANDERSON1808/archTest/tradicionales/testpackage/nested/dep").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/nesteddependency")

		assertNoError(t, mockT)
	})

	t.Run("Ignored packages ignore ignored transitive packages", func(t *testing.T) {
		mockT := new(testingT)

		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			Ignoring("github.com/this/is/verifying/multiple/exclusions", "github.com/ANDERSON1808/archTest/tradicionales/...").
			Ignoring("github.com/this/is/verifying/chaining").
			ShouldNotDependOn("github.com/ANDERSON1808/archTest/tradicionales/transative")

		assertNoError(t, mockT)
	})
}

func TestPackage_ShouldNotDependDirectly(t *testing.T) {

	t.Run("Fails on direct dependencies", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependDirectlyOn("github.com/ANDERSON1808/archTest/tradicionales/dependency")

		assertError(t, mockT,
			"github.com/ANDERSON1808/archTest/tradicionales/testpackage",
			"github.com/ANDERSON1808/archTest/tradicionales/dependency")
	})

	t.Run("Fails on transative dependencies", func(t *testing.T) {
		mockT := new(testingT)
		Package(mockT, "github.com/ANDERSON1808/archTest/tradicionales/testpackage").
			ShouldNotDependDirectlyOn("github.com/ANDERSON1808/archTest/tradicionales/transative")

		assertNoError(t, mockT)
	})
}

func assertNoError(t *testing.T, mockT *testingT) {
	t.Helper()
	if mockT.errored() {
		t.Fatalf("archtest should not have failed but, %s", mockT.message())
	}
}

func assertError(t *testing.T, mockT *testingT, dependencyTrace ...string) {
	t.Helper()
	if !mockT.errored() {
		t.Fatal("archtest did not fail on dependency")
	}

	if dependencyTrace == nil {
		return
	}

	s := strings.Builder{}
	s.WriteString("Error:\n")
	for i, v := range dependencyTrace {
		s.WriteString(strings.Repeat("\t", i))
		s.WriteString(v + "\n")
	}

	if mockT.message() != s.String() {
		t.Errorf("expected %s got error message: %s", s.String(), mockT.message())
	}
}

type testingT struct {
	errors [][]interface{}
}

func (t *testingT) Error(args ...interface{}) {
	t.errors = append(t.errors, args)
}

func (t *testingT) errored() bool {
	return len(t.errors) != 0
}

func (t *testingT) message() interface{} {
	return t.errors[0][0]
}

func (t *testingT) Errorf(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	t.Error(message)
}
