// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ProcessPoolClass = _ProcessPoolClass{objc.GetClass("WKProcessPool")}

type _ProcessPoolClass struct {
	objc.Class
}

type IProcessPool interface {
	objc.IObject
}

type ProcessPool struct {
	objc.Object
}

func MakeProcessPool(ptr unsafe.Pointer) ProcessPool {
	return ProcessPool{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _ProcessPoolClass) Alloc() ProcessPool {
	rv := ffi.CallMethod[ProcessPool](pc, "alloc")
	return rv
}

func (p_ ProcessPool) Init() ProcessPool {
	rv := ffi.CallMethod[ProcessPool](p_, "init")
	return rv
}

func (pc _ProcessPoolClass) New() ProcessPool {
	rv := ffi.CallMethod[ProcessPool](pc, "new")
	rv.Autorelease()
	return rv
}

func NewProcessPool() ProcessPool {
	return ProcessPoolClass.New()
}
