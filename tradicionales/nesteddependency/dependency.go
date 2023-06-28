package nesteddependency

import (
	"fmt"

	"github.com/ANDERSON1808/archTest/tradicionales/transative"
)

const Item = "depend on me"

func Somemethod() {
	fmt.Println(transative.NowYouDependOnMe)
}
