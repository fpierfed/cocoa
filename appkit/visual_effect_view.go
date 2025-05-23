// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var VisualEffectViewClass = _VisualEffectViewClass{objc.GetClass("NSVisualEffectView")}

type _VisualEffectViewClass struct {
	objc.Class
}

type IVisualEffectView interface {
	IView
	Material() VisualEffectMaterial
	SetMaterial(value VisualEffectMaterial)
	BlendingMode() VisualEffectBlendingMode
	SetBlendingMode(value VisualEffectBlendingMode)
	IsEmphasized() bool
	SetEmphasized(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	MaskImage() Image
	SetMaskImage(value IImage)
	State() VisualEffectState
	SetState(value VisualEffectState)
}

type VisualEffectView struct {
	View
}

func MakeVisualEffectView(ptr unsafe.Pointer) VisualEffectView {
	return VisualEffectView{
		View: MakeView(ptr),
	}
}

func (v_ VisualEffectView) InitWithFrame(frameRect foundation.Rect) VisualEffectView {
	rv := ffi.CallMethod[VisualEffectView](v_, "initWithFrame:", frameRect)
	return rv
}

func (v_ VisualEffectView) Init() VisualEffectView {
	rv := ffi.CallMethod[VisualEffectView](v_, "init")
	return rv
}

func (vc _VisualEffectViewClass) Alloc() VisualEffectView {
	rv := ffi.CallMethod[VisualEffectView](vc, "alloc")
	return rv
}

func (vc _VisualEffectViewClass) New() VisualEffectView {
	rv := ffi.CallMethod[VisualEffectView](vc, "new")
	rv.Autorelease()
	return rv
}

func NewVisualEffectView() VisualEffectView {
	return VisualEffectViewClass.New()
}

func (v_ VisualEffectView) Material() VisualEffectMaterial {
	rv := ffi.CallMethod[VisualEffectMaterial](v_, "material")
	return rv
}

func (v_ VisualEffectView) SetMaterial(value VisualEffectMaterial) {
	ffi.CallMethod[ffi.Void](v_, "setMaterial:", value)
}

func (v_ VisualEffectView) BlendingMode() VisualEffectBlendingMode {
	rv := ffi.CallMethod[VisualEffectBlendingMode](v_, "blendingMode")
	return rv
}

func (v_ VisualEffectView) SetBlendingMode(value VisualEffectBlendingMode) {
	ffi.CallMethod[ffi.Void](v_, "setBlendingMode:", value)
}

func (v_ VisualEffectView) IsEmphasized() bool {
	rv := ffi.CallMethod[bool](v_, "isEmphasized")
	return rv
}

func (v_ VisualEffectView) SetEmphasized(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setEmphasized:", value)
}

func (v_ VisualEffectView) InteriorBackgroundStyle() BackgroundStyle {
	rv := ffi.CallMethod[BackgroundStyle](v_, "interiorBackgroundStyle")
	return rv
}

func (v_ VisualEffectView) MaskImage() Image {
	rv := ffi.CallMethod[Image](v_, "maskImage")
	return rv
}

func (v_ VisualEffectView) SetMaskImage(value IImage) {
	ffi.CallMethod[ffi.Void](v_, "setMaskImage:", value)
}

func (v_ VisualEffectView) State() VisualEffectState {
	rv := ffi.CallMethod[VisualEffectState](v_, "state")
	return rv
}

func (v_ VisualEffectView) SetState(value VisualEffectState) {
	ffi.CallMethod[ffi.Void](v_, "setState:", value)
}
