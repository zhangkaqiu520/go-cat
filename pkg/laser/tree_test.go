package laser

import (
	"testing"
)

func TestTree(t *testing.T) {
	n := NewNode()
	n.Insert("/abc")
	n.Insert("/abc/ddd/bbb/ccc")
	n.Insert("/abc/ddd/bbb/ddd")
	n.Fetch()
}
