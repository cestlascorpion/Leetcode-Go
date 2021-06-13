package solution

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	fmt.Println(BadVersion, VersionList)
	exp := BadVersion + 1
	res := FirstBadVersion(len(VersionList))
	if exp != res {
		t.Fail()
	}
}
