// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ScriptMessageClass = _ScriptMessageClass{objc.GetClass("WKScriptMessage")}

type _ScriptMessageClass struct {
	objc.Class
}

type IScriptMessage interface {
	objc.IObject
	Body() objc.Object
	FrameInfo() FrameInfo
	WebView() WebView
	World() ContentWorld
	Name() string
}

type ScriptMessage struct {
	objc.Object
}

func MakeScriptMessage(ptr unsafe.Pointer) ScriptMessage {
	return ScriptMessage{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScriptMessageClass) Alloc() ScriptMessage {
	rv := ffi.CallMethod[ScriptMessage](sc, "alloc")
	return rv
}

func (s_ ScriptMessage) Init() ScriptMessage {
	rv := ffi.CallMethod[ScriptMessage](s_, "init")
	return rv
}

func (sc _ScriptMessageClass) New() ScriptMessage {
	rv := ffi.CallMethod[ScriptMessage](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScriptMessage() ScriptMessage {
	return ScriptMessageClass.New()
}

func (s_ ScriptMessage) Body() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "body")
	return rv
}

func (s_ ScriptMessage) FrameInfo() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](s_, "frameInfo")
	return rv
}

func (s_ ScriptMessage) WebView() WebView {
	rv := ffi.CallMethod[WebView](s_, "webView")
	return rv
}

func (s_ ScriptMessage) World() ContentWorld {
	rv := ffi.CallMethod[ContentWorld](s_, "world")
	return rv
}

func (s_ ScriptMessage) Name() string {
	rv := ffi.CallMethod[string](s_, "name")
	return rv
}

type ScriptMessageHandler interface {
	// required
	UserContentController_DidReceiveScriptMessage(userContentController UserContentController, message ScriptMessage)
}

type ScriptMessageHandlerWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWrapper) UserContentController_DidReceiveScriptMessage(userContentController IUserContentController, message IScriptMessage) {
	ffi.CallMethod[ffi.Void](s_, "userContentController:didReceiveScriptMessage:", userContentController, message)
}

type ScriptMessageHandlerWithReply interface {
	// required
	UserContentController_DidReceiveScriptMessage_ReplyHandler(userContentController UserContentController, message ScriptMessage, replyHandler func(reply objc.IObject, errorMessage string))
}

type ScriptMessageHandlerWithReplyWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWithReplyWrapper) UserContentController_DidReceiveScriptMessage_ReplyHandler(userContentController IUserContentController, message IScriptMessage, replyHandler func(reply objc.Object, errorMessage string)) {
	ffi.CallMethod[ffi.Void](s_, "userContentController:didReceiveScriptMessage:replyHandler:", userContentController, message, replyHandler)
}
