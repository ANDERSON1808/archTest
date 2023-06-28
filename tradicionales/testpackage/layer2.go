package testpackage

import (
	"crypto"
	"fmt"

	"github.com/ANDERSON1808/archTest/tradicionales/dependency"
)

func What(a crypto.Decrypter) {
	fmt.Println(dependency.Item)
}
