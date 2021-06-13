package solution

import (
	"time"
)

var VersionList []bool
var BadVersion int

func init() {
	VersionList = make([]bool, 10, 10)
	BadVersion = int(time.Now().Unix()) % len(VersionList)

	for i := BadVersion; i < len(VersionList); i++ {
		VersionList[i] = true
	}
}

func isBadVersion(version int) bool {
	return VersionList[version-1]
}

func FirstBadVersion(n int) int {
	l, h := 1, n
	for l < h {
		mid := l + (h-l)/2
		if isBadVersion(mid) {
			h = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
