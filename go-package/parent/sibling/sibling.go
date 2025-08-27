package sibling

import (
	"test-go-module-package/parent/internal"
	internalchild "test-go-module-package/parent/internal/internal_child"
)

func CallInternal(){
	internal.InternalMessage("getting called from sibling package")
	internalchild.InternalChildMessage("getting called from sibling package")
}
