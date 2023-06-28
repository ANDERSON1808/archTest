# ArchTest

Project to validate clean architecture in golang

"archTest" is a Go package that provides a testing framework for architecture compliance in Go projects. It allows you
to define and enforce architectural rules and dependencies between packages.

### Installation

To use archTest, you need to have Go installed and set up on your machine. You can install the package using the
following go get command:

```golang
  go get github.com/ANDERSON1808/archTest.git
```

### Usage

To use archTest, you need to import the package in your test file:

```golang
  import "github.com/ANDERSON1808/archTest.git
```

### Package Function

The Package function is the entry point for defining architecture tests. It creates a new PackageTest instance.

```golang
  func Package(t TestingT, packageName ...string) *PackageTest
```

- t: A TestingT interface implementation that provides testing functionality (e.g., *testing.T).
- packageName: Variadic parameter(s) representing the package name(s) to be tested.

### PackageTest Methods

The PackageTest type provides the following methods:

IncludeTests

```golang
  func (t *PackageTest) IncludeTests() *PackageTest
```

This method includes test packages when analyzing dependencies.
Ignoring

```golang
func (t *PackageTest) Ignoring(e ...string) *PackageTest
```

This method allows you to specify packages to be ignored during the dependency analysis.

ShouldNotDependDirectlyOn

```golang
func (t *PackageTest) ShouldNotDependDirectlyOn(pegs ...string)
```

This method specifies that the tested package(s) should not depend directly on the provided package(s). It generates an
error if a direct dependency is found.

ShouldNotDependOn

```golang
func (t *PackageTest) ShouldNotDependOn(pkgs ...string)
```

This method specifies that the tested package(s) should not depend on the provided package(s) at any depth. It generates
an error if a dependency is found.

TestingT Interface
The TestingT interface is used for error reporting during testing. It is usually implemented by testing frameworks such
as *testing.T. The interface provides an Errorf method to report test failures.

### Example

Here's an example of how to use archTest to enforce architectural rules in a Go project:

```golang
package myproject_test

import (
	"testing"

	"github.com/ANDERSON1808/archTest"
)

func TestArchitecture(t *testing.T) {
	archTest.Package(t, "myproject").
		IncludeTests().
		ShouldNotDependDirectlyOn("externalpkg").
		ShouldNotDependOn("badpkg").
		ShouldNotDependOn("anotherpkg")
}
```

In the example above, the TestArchitecture function defines architectural rules for the myproject package. It ensures
that the myproject package should not depend directly on the externalpkg package and should not depend on the badpkg and
anotherpkg packages at any depth. If any violations are found, the test will fail.

---------------------

#### Checking for dependencies

```golang
archtest.Package(t, "github.com/myorg/myproject/....").
ShouldNotDependOn("github.com/some/package")
```

#### Checking for direct dependencies

```golang
archtest.Package(t, "github.com/myorg/myproject/....").
ShouldNotDependDirectlyOn("github.com/some/package")
```

#### Including Tests

```golang
archtest.Package(t, "github.com/myorg/myproject/....").
IncludeTests().
ShouldNotDependDirectlyOn("github.com/some/package")
```

### Conclusion
archTest provides a convenient way to enforce architectural rules and dependencies in Go projects. By using the provided methods, you can define your architectural constraints and ensure that your codebase adheres to them.