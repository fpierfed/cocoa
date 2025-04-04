// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}

type _GestureRecognizerClass struct {
	objc.Class
}

type IGestureRecognizer interface {
	objc.IObject
	LocationInView(view IView) foundation.Point
	Reset()
	MouseDown(event IEvent)
	MouseDragged(event IEvent)
	MouseUp(event IEvent)
	OtherMouseDown(event IEvent)
	OtherMouseDragged(event IEvent)
	OtherMouseUp(event IEvent)
	RightMouseDown(event IEvent)
	RightMouseDragged(event IEvent)
	RightMouseUp(event IEvent)
	MagnifyWithEvent(event IEvent)
	RotateWithEvent(event IEvent)
	CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool
	CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool
	ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool
	KeyDown(event IEvent)
	KeyUp(event IEvent)
	TabletPoint(event IEvent)
	FlagsChanged(event IEvent)
	PressureChangeWithEvent(event IEvent)
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() GestureRecognizerState
	View() View
	IsEnabled() bool
	SetEnabled(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	Delegate() GestureRecognizerDelegateWrapper
	SetDelegate(value GestureRecognizerDelegate)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
}

type GestureRecognizer struct {
	objc.Object
}

func MakeGestureRecognizer(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{
		Object: objc.MakeObject(ptr),
	}
}

func (g_ GestureRecognizer) InitWithTarget_Action(target objc.IObject, action objc.Selector) GestureRecognizer {
	rv := ffi.CallMethod[GestureRecognizer](g_, "initWithTarget:action:", target, action)
	return rv
}

func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := ffi.CallMethod[GestureRecognizer](gc, "alloc")
	return rv
}

func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := ffi.CallMethod[GestureRecognizer](g_, "init")
	return rv
}

func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := ffi.CallMethod[GestureRecognizer](gc, "new")
	rv.Autorelease()
	return rv
}

func NewGestureRecognizer() GestureRecognizer {
	return GestureRecognizerClass.New()
}

func (g_ GestureRecognizer) LocationInView(view IView) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](g_, "locationInView:", view)
	return rv
}

func (g_ GestureRecognizer) Reset() {
	ffi.CallMethod[ffi.Void](g_, "reset")
}

func (g_ GestureRecognizer) MouseDown(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "mouseDown:", event)
}

func (g_ GestureRecognizer) MouseDragged(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "mouseDragged:", event)
}

func (g_ GestureRecognizer) MouseUp(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "mouseUp:", event)
}

func (g_ GestureRecognizer) OtherMouseDown(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "otherMouseDown:", event)
}

func (g_ GestureRecognizer) OtherMouseDragged(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "otherMouseDragged:", event)
}

func (g_ GestureRecognizer) OtherMouseUp(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "otherMouseUp:", event)
}

func (g_ GestureRecognizer) RightMouseDown(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "rightMouseDown:", event)
}

func (g_ GestureRecognizer) RightMouseDragged(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "rightMouseDragged:", event)
}

func (g_ GestureRecognizer) RightMouseUp(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "rightMouseUp:", event)
}

func (g_ GestureRecognizer) MagnifyWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "magnifyWithEvent:", event)
}

func (g_ GestureRecognizer) RotateWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "rotateWithEvent:", event)
}

func (g_ GestureRecognizer) CanBePreventedByGestureRecognizer(preventingGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "canBePreventedByGestureRecognizer:", preventingGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) CanPreventGestureRecognizer(preventedGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "canPreventGestureRecognizer:", preventedGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) ShouldBeRequiredToFailByGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "shouldBeRequiredToFailByGestureRecognizer:", otherGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) ShouldRequireFailureOfGestureRecognizer(otherGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "shouldRequireFailureOfGestureRecognizer:", otherGestureRecognizer)
	return rv
}

func (g_ GestureRecognizer) KeyDown(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "keyDown:", event)
}

func (g_ GestureRecognizer) KeyUp(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "keyUp:", event)
}

func (g_ GestureRecognizer) TabletPoint(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "tabletPoint:", event)
}

func (g_ GestureRecognizer) FlagsChanged(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "flagsChanged:", event)
}

func (g_ GestureRecognizer) PressureChangeWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "pressureChangeWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "touchesBeganWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "touchesCancelledWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "touchesEndedWithEvent:", event)
}

func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](g_, "touchesMovedWithEvent:", event)
}

func (g_ GestureRecognizer) Action() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](g_, "action")
	return rv
}

func (g_ GestureRecognizer) SetAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](g_, "setAction:", value)
}

func (g_ GestureRecognizer) Target() objc.Object {
	rv := ffi.CallMethod[objc.Object](g_, "target")
	return rv
}

func (g_ GestureRecognizer) SetTarget(value objc.IObject) {
	ffi.CallMethod[ffi.Void](g_, "setTarget:", value)
}

func (g_ GestureRecognizer) State() GestureRecognizerState {
	rv := ffi.CallMethod[GestureRecognizerState](g_, "state")
	return rv
}

func (g_ GestureRecognizer) View() View {
	rv := ffi.CallMethod[View](g_, "view")
	return rv
}

func (g_ GestureRecognizer) IsEnabled() bool {
	rv := ffi.CallMethod[bool](g_, "isEnabled")
	return rv
}

func (g_ GestureRecognizer) SetEnabled(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setEnabled:", value)
}

func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysPrimaryMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysPrimaryMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysSecondaryMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysSecondaryMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysOtherMouseButtonEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysOtherMouseButtonEvents:", value)
}

func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysKeyEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysKeyEvents:", value)
}

func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysMagnificationEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysMagnificationEvents:", value)
}

func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := ffi.CallMethod[bool](g_, "delaysRotationEvents")
	return rv
}

func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	ffi.CallMethod[ffi.Void](g_, "setDelaysRotationEvents:", value)
}

func (g_ GestureRecognizer) Delegate() GestureRecognizerDelegateWrapper {
	rv := ffi.CallMethod[GestureRecognizerDelegateWrapper](g_, "delegate")
	return rv
}

func (g_ GestureRecognizer) SetDelegate(value GestureRecognizerDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(g_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](g_, "setDelegate:", po)
}

func (g_ GestureRecognizer) PressureConfiguration() PressureConfiguration {
	rv := ffi.CallMethod[PressureConfiguration](g_, "pressureConfiguration")
	return rv
}

func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	ffi.CallMethod[ffi.Void](g_, "setPressureConfiguration:", value)
}

func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := ffi.CallMethod[TouchTypeMask](g_, "allowedTouchTypes")
	return rv
}

func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	ffi.CallMethod[ffi.Void](g_, "setAllowedTouchTypes:", value)
}

type GestureRecognizerDelegate interface {
	ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool
	// optional
	GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool
	ImplementsGestureRecognizerShouldBegin() bool
	// optional
	GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool
	// optional
	GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	ImplementsGestureRecognizer_ShouldReceiveTouch() bool
	// optional
	GestureRecognizer_ShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool
}

type GestureRecognizerDelegateImpl struct {
	_GestureRecognizer_ShouldAttemptToRecognizeWithEvent                  func(gestureRecognizer GestureRecognizer, event Event) bool
	_GestureRecognizerShouldBegin                                         func(gestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldRequireFailureOfGestureRecognizer            func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer          func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool
	_GestureRecognizer_ShouldReceiveTouch                                 func(gestureRecognizer GestureRecognizer, touch Touch) bool
}

func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool {
	return di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldAttemptToRecognizeWithEvent(f func(gestureRecognizer GestureRecognizer, event Event) bool) {
	di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer GestureRecognizer, event Event) bool {
	return di._GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer, event)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizerShouldBegin() bool {
	return di._GestureRecognizerShouldBegin != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizerShouldBegin(f func(gestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizerShouldBegin = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizerShouldBegin(gestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizerShouldBegin(gestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool {
	return di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(f func(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool) {
	di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer GestureRecognizer, otherGestureRecognizer GestureRecognizer) bool {
	return di._GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer, otherGestureRecognizer)
}
func (di *GestureRecognizerDelegateImpl) ImplementsGestureRecognizer_ShouldReceiveTouch() bool {
	return di._GestureRecognizer_ShouldReceiveTouch != nil
}

func (di *GestureRecognizerDelegateImpl) SetGestureRecognizer_ShouldReceiveTouch(f func(gestureRecognizer GestureRecognizer, touch Touch) bool) {
	di._GestureRecognizer_ShouldReceiveTouch = f
}

func (di *GestureRecognizerDelegateImpl) GestureRecognizer_ShouldReceiveTouch(gestureRecognizer GestureRecognizer, touch Touch) bool {
	return di._GestureRecognizer_ShouldReceiveTouch(gestureRecognizer, touch)
}

type GestureRecognizerDelegateWrapper struct {
	objc.Object
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldAttemptToRecognizeWithEvent() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldAttemptToRecognizeWithEvent(gestureRecognizer IGestureRecognizer, event IEvent) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizer:shouldAttemptToRecognizeWithEvent:", gestureRecognizer, event)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizerShouldBegin() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizerShouldBegin:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizerShouldBegin(gestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizerShouldBegin:", gestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldRequireFailureOfGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldRequireFailureOfGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizer:shouldRequireFailureOfGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer IGestureRecognizer, otherGestureRecognizer IGestureRecognizer) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:", gestureRecognizer, otherGestureRecognizer)
	return rv
}

func (g_ *GestureRecognizerDelegateWrapper) ImplementsGestureRecognizer_ShouldReceiveTouch() bool {
	return g_.RespondsToSelector(objc.GetSelector("gestureRecognizer:shouldReceiveTouch:"))
}

func (g_ GestureRecognizerDelegateWrapper) GestureRecognizer_ShouldReceiveTouch(gestureRecognizer IGestureRecognizer, touch ITouch) bool {
	rv := ffi.CallMethod[bool](g_, "gestureRecognizer:shouldReceiveTouch:", gestureRecognizer, touch)
	return rv
}
