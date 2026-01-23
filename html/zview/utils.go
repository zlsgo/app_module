package zview

type Swap string

const (
	SwapInner      Swap = "inner"
	SwapReplace    Swap = "replace"
	SwapAppend     Swap = "append"
	SwapPrepend    Swap = "prepend"
	SwapBeforeBegin Swap = "beforebegin"
	SwapAfterEnd   Swap = "afterend"
	SwapMorph      Swap = "morph"
	SwapMorphAll   Swap = "morph-all"
	SwapSkip       Swap = "skip"
	SwapPush       Swap = "push"
	SwapPop        Swap = "pop"
)
