// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ScrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}

type _ScrollLayerClass struct {
	objc.Class
}

type IScrollLayer interface {
	ILayer
	ScrollToPoint(p coregraphics.Point)
	ScrollToRect(r coregraphics.Rect)
	ScrollMode() ScrollLayerScrollMode
	SetScrollMode(value ScrollLayerScrollMode)
}

type ScrollLayer struct {
	Layer
}

func MakeScrollLayer(ptr unsafe.Pointer) ScrollLayer {
	return ScrollLayer{
		Layer: MakeLayer(ptr),
	}
}

func (sc _ScrollLayerClass) Layer() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](sc, "layer")
	return rv
}

func (s_ ScrollLayer) Init() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](s_, "init")
	return rv
}

func (s_ ScrollLayer) InitWithLayer(layer objc.IObject) ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](s_, "initWithLayer:", layer)
	return rv
}

func (s_ ScrollLayer) PresentationLayer() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](s_, "presentationLayer")
	return rv
}

func (s_ ScrollLayer) ModelLayer() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](s_, "modelLayer")
	return rv
}

func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](sc, "alloc")
	return rv
}

func (sc _ScrollLayerClass) New() ScrollLayer {
	rv := ffi.CallMethod[ScrollLayer](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScrollLayer() ScrollLayer {
	return ScrollLayerClass.New()
}

func (s_ ScrollLayer) ScrollToPoint(p coregraphics.Point) {
	ffi.CallMethod[ffi.Void](s_, "scrollToPoint:", p)
}

func (s_ ScrollLayer) ScrollToRect(r coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](s_, "scrollToRect:", r)
}

func (s_ ScrollLayer) ScrollMode() ScrollLayerScrollMode {
	rv := ffi.CallMethod[ScrollLayerScrollMode](s_, "scrollMode")
	return rv
}

func (s_ ScrollLayer) SetScrollMode(value ScrollLayerScrollMode) {
	ffi.CallMethod[ffi.Void](s_, "setScrollMode:", value)
}
