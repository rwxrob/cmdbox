package fmt

import "testing"

func TestEmphasize(t *testing.T) {
	// TODO better test
	t.Log(Emphasize("here is *italic*"))
	t.Log(Emphasize("here is **bold**"))
	t.Log(Emphasize("here is ***bolditalic***"))
}
