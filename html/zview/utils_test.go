package zview

import "testing"

func TestSwapValues(t *testing.T) {
	tests := map[Swap]string{
		SwapInner: "inner", SwapReplace: "replace", SwapAppend: "append",
		SwapPrepend: "prepend", SwapBeforeBegin: "beforebegin", SwapAfterEnd: "afterend",
		SwapMorph: "morph", SwapMorphAll: "morph-all", SwapSkip: "skip",
		SwapPush: "push", SwapPop: "pop",
	}
	for value, want := range tests {
		if string(value) != want {
			t.Errorf("swap %q = %q, want %q", value, value, want)
		}
	}
}
