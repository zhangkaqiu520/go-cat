package laser

import (
	"testing"
)

func TestTree(t *testing.T) {
	n := NewNode()
	n.Insert("/abc")
	n.Insert("/abc/ddd/:room_id/ccc")
	n.Insert("/abc/ddd/bbb/ddd")
	n.Search("/abc/ddd/5/ccc")
}
