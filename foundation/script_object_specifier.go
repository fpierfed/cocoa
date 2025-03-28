// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ScriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}

type _ScriptObjectSpecifierClass struct {
	objc.Class
}

type IScriptObjectSpecifier interface {
	objc.IObject
	IndicesOfObjectsByEvaluatingWithContainer_Count(container objc.IObject, count *int) *int
	ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object
	ObjectsByEvaluatingSpecifier() objc.Object
	ContainerClassDescription() ScriptClassDescription
	SetContainerClassDescription(value IScriptClassDescription)
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
	ContainerIsRangeContainerObject() bool
	SetContainerIsRangeContainerObject(value bool)
	ContainerSpecifier() ScriptObjectSpecifier
	SetContainerSpecifier(value IScriptObjectSpecifier)
	ChildSpecifier() ScriptObjectSpecifier
	SetChildSpecifier(value IScriptObjectSpecifier)
	Key() string
	SetKey(value string)
	KeyClassDescription() ScriptClassDescription
	EvaluationErrorSpecifier() ScriptObjectSpecifier
	EvaluationErrorNumber() int
	SetEvaluationErrorNumber(value int)
	Descriptor() AppleEventDescriptor
}

type ScriptObjectSpecifier struct {
	objc.Object
}

func MakeScriptObjectSpecifier(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ ScriptObjectSpecifier) InitWithContainerClassDescription_ContainerSpecifier_Key(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "initWithContainerClassDescription:containerSpecifier:key:", classDesc, container, property)
	return rv
}

func (s_ ScriptObjectSpecifier) InitWithContainerSpecifier_Key(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "initWithContainerSpecifier:key:", container, property)
	return rv
}

func (sc _ScriptObjectSpecifierClass) Alloc() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](sc, "alloc")
	return rv
}

func (s_ ScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "init")
	return rv
}

func (sc _ScriptObjectSpecifierClass) New() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScriptObjectSpecifier() ScriptObjectSpecifier {
	return ScriptObjectSpecifierClass.New()
}

func (sc _ScriptObjectSpecifierClass) ObjectSpecifierWithDescriptor(descriptor IAppleEventDescriptor) ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](sc, "objectSpecifierWithDescriptor:", descriptor)
	return rv
}

func (s_ ScriptObjectSpecifier) IndicesOfObjectsByEvaluatingWithContainer_Count(container objc.IObject, count *int) *int {
	rv := ffi.CallMethod[*int](s_, "indicesOfObjectsByEvaluatingWithContainer:count:", container, count)
	return rv
}

func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingWithContainers(containers objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "objectsByEvaluatingWithContainers:", containers)
	return rv
}

func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "objectsByEvaluatingSpecifier")
	return rv
}

func (s_ ScriptObjectSpecifier) ContainerClassDescription() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](s_, "containerClassDescription")
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerClassDescription(value IScriptClassDescription) {
	ffi.CallMethod[ffi.Void](s_, "setContainerClassDescription:", value)
}

func (s_ ScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	rv := ffi.CallMethod[bool](s_, "containerIsObjectBeingTested")
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setContainerIsObjectBeingTested:", value)
}

func (s_ ScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	rv := ffi.CallMethod[bool](s_, "containerIsRangeContainerObject")
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setContainerIsRangeContainerObject:", value)
}

func (s_ ScriptObjectSpecifier) ContainerSpecifier() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "containerSpecifier")
	return rv
}

func (s_ ScriptObjectSpecifier) SetContainerSpecifier(value IScriptObjectSpecifier) {
	ffi.CallMethod[ffi.Void](s_, "setContainerSpecifier:", value)
}

func (s_ ScriptObjectSpecifier) ChildSpecifier() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "childSpecifier")
	return rv
}

func (s_ ScriptObjectSpecifier) SetChildSpecifier(value IScriptObjectSpecifier) {
	ffi.CallMethod[ffi.Void](s_, "setChildSpecifier:", value)
}

func (s_ ScriptObjectSpecifier) Key() string {
	rv := ffi.CallMethod[string](s_, "key")
	return rv
}

func (s_ ScriptObjectSpecifier) SetKey(value string) {
	ffi.CallMethod[ffi.Void](s_, "setKey:", value)
}

func (s_ ScriptObjectSpecifier) KeyClassDescription() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](s_, "keyClassDescription")
	return rv
}

func (s_ ScriptObjectSpecifier) EvaluationErrorSpecifier() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "evaluationErrorSpecifier")
	return rv
}

func (s_ ScriptObjectSpecifier) EvaluationErrorNumber() int {
	rv := ffi.CallMethod[int](s_, "evaluationErrorNumber")
	return rv
}

func (s_ ScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	ffi.CallMethod[ffi.Void](s_, "setEvaluationErrorNumber:", value)
}

func (s_ ScriptObjectSpecifier) Descriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](s_, "descriptor")
	return rv
}
