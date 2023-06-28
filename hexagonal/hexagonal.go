package hexagonal

import "fmt"

type HexagonalValidator struct {
	allowedDependencies map[string][]string
}

func NewHexagonalValidator() *HexagonalValidator {
	return &HexagonalValidator{
		allowedDependencies: make(map[string][]string),
	}
}

func (v *HexagonalValidator) AddPackage(packageName string, dependencies []string) {
	v.allowedDependencies[packageName] = dependencies
}

func (v *HexagonalValidator) Validate(dependencies map[string][]string) error {
	for packageName, allowedDeps := range v.allowedDependencies {
		actualDeps, ok := dependencies[packageName]
		if !ok {
			return fmt.Errorf("missing package: %s", packageName)
		}

		for _, actualDep := range actualDeps {
			if !contains(allowedDeps, actualDep) {
				return fmt.Errorf("invalid dependency: %s -> %s", packageName, actualDep)
			}
		}
	}

	return nil
}

func (v *HexagonalValidator) AddPackageDependency(packageName string, dependency string) {
	dependencies, ok := v.allowedDependencies[packageName]
	if !ok {
		dependencies = []string{}
	}
	dependencies = append(dependencies, dependency)
	v.allowedDependencies[packageName] = dependencies
}

func (v *HexagonalValidator) RemovePackage(packageName string) {
	delete(v.allowedDependencies, packageName)
}

func (v *HexagonalValidator) RemovePackageDependency(packageName string, dependency string) {
	dependencies, ok := v.allowedDependencies[packageName]
	if !ok {
		return
	}

	var updatedDependencies []string
	for _, dep := range dependencies {
		if dep != dependency {
			updatedDependencies = append(updatedDependencies, dep)
		}
	}

	v.allowedDependencies[packageName] = updatedDependencies
}

func (v *HexagonalValidator) GetAllowedDependencies(packageName string) ([]string, bool) {
	dependencies, ok := v.allowedDependencies[packageName]
	return dependencies, ok
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
