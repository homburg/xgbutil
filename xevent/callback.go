/*
    Does all the plumbing to allow a simple callback interface for users.

    This file is automatically generated using `scripts/write-events callbacks`.

    Edit it at your peril.
*/
package xevent

import "burntsushi.net/go/x-go-binding/xgb"
import "burntsushi.net/go/xgbutil"

type KeyPressFun func(xu *xgbutil.XUtil, event KeyPressEvent)

func (callback KeyPressFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(KeyPress, win, callback)
}

func (callback KeyPressFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(KeyPressEvent))
}

type KeyReleaseFun func(xu *xgbutil.XUtil, event KeyReleaseEvent)

func (callback KeyReleaseFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(KeyRelease, win, callback)
}

func (callback KeyReleaseFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(KeyReleaseEvent))
}

type ButtonPressFun func(xu *xgbutil.XUtil, event ButtonPressEvent)

func (callback ButtonPressFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ButtonPress, win, callback)
}

func (callback ButtonPressFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ButtonPressEvent))
}

type ButtonReleaseFun func(xu *xgbutil.XUtil, event ButtonReleaseEvent)

func (callback ButtonReleaseFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ButtonRelease, win, callback)
}

func (callback ButtonReleaseFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ButtonReleaseEvent))
}

type MotionNotifyFun func(xu *xgbutil.XUtil, event MotionNotifyEvent)

func (callback MotionNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(MotionNotify, win, callback)
}

func (callback MotionNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(MotionNotifyEvent))
}

type EnterNotifyFun func(xu *xgbutil.XUtil, event EnterNotifyEvent)

func (callback EnterNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(EnterNotify, win, callback)
}

func (callback EnterNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(EnterNotifyEvent))
}

type LeaveNotifyFun func(xu *xgbutil.XUtil, event LeaveNotifyEvent)

func (callback LeaveNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(LeaveNotify, win, callback)
}

func (callback LeaveNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(LeaveNotifyEvent))
}

type FocusInFun func(xu *xgbutil.XUtil, event FocusInEvent)

func (callback FocusInFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(FocusIn, win, callback)
}

func (callback FocusInFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(FocusInEvent))
}

type FocusOutFun func(xu *xgbutil.XUtil, event FocusOutEvent)

func (callback FocusOutFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(FocusOut, win, callback)
}

func (callback FocusOutFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(FocusOutEvent))
}

type KeymapNotifyFun func(xu *xgbutil.XUtil, event KeymapNotifyEvent)

func (callback KeymapNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(KeymapNotify, win, callback)
}

func (callback KeymapNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(KeymapNotifyEvent))
}

type ExposeFun func(xu *xgbutil.XUtil, event ExposeEvent)

func (callback ExposeFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(Expose, win, callback)
}

func (callback ExposeFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ExposeEvent))
}

type GraphicsExposureFun func(xu *xgbutil.XUtil, event GraphicsExposureEvent)

func (callback GraphicsExposureFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(GraphicsExposure, win, callback)
}

func (callback GraphicsExposureFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(GraphicsExposureEvent))
}

type NoExposureFun func(xu *xgbutil.XUtil, event NoExposureEvent)

func (callback NoExposureFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(NoExposure, win, callback)
}

func (callback NoExposureFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(NoExposureEvent))
}

type VisibilityNotifyFun func(xu *xgbutil.XUtil, event VisibilityNotifyEvent)

func (callback VisibilityNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(VisibilityNotify, win, callback)
}

func (callback VisibilityNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(VisibilityNotifyEvent))
}

type CreateNotifyFun func(xu *xgbutil.XUtil, event CreateNotifyEvent)

func (callback CreateNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(CreateNotify, win, callback)
}

func (callback CreateNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(CreateNotifyEvent))
}

type DestroyNotifyFun func(xu *xgbutil.XUtil, event DestroyNotifyEvent)

func (callback DestroyNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(DestroyNotify, win, callback)
}

func (callback DestroyNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(DestroyNotifyEvent))
}

type UnmapNotifyFun func(xu *xgbutil.XUtil, event UnmapNotifyEvent)

func (callback UnmapNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(UnmapNotify, win, callback)
}

func (callback UnmapNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(UnmapNotifyEvent))
}

type MapNotifyFun func(xu *xgbutil.XUtil, event MapNotifyEvent)

func (callback MapNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(MapNotify, win, callback)
}

func (callback MapNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(MapNotifyEvent))
}

type MapRequestFun func(xu *xgbutil.XUtil, event MapRequestEvent)

func (callback MapRequestFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(MapRequest, win, callback)
}

func (callback MapRequestFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(MapRequestEvent))
}

type ReparentNotifyFun func(xu *xgbutil.XUtil, event ReparentNotifyEvent)

func (callback ReparentNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ReparentNotify, win, callback)
}

func (callback ReparentNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ReparentNotifyEvent))
}

type ConfigureNotifyFun func(xu *xgbutil.XUtil, event ConfigureNotifyEvent)

func (callback ConfigureNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ConfigureNotify, win, callback)
}

func (callback ConfigureNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ConfigureNotifyEvent))
}

type ConfigureRequestFun func(xu *xgbutil.XUtil, event ConfigureRequestEvent)

func (callback ConfigureRequestFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ConfigureRequest, win, callback)
}

func (callback ConfigureRequestFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ConfigureRequestEvent))
}

type GravityNotifyFun func(xu *xgbutil.XUtil, event GravityNotifyEvent)

func (callback GravityNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(GravityNotify, win, callback)
}

func (callback GravityNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(GravityNotifyEvent))
}

type ResizeRequestFun func(xu *xgbutil.XUtil, event ResizeRequestEvent)

func (callback ResizeRequestFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ResizeRequest, win, callback)
}

func (callback ResizeRequestFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ResizeRequestEvent))
}

type CirculateNotifyFun func(xu *xgbutil.XUtil, event CirculateNotifyEvent)

func (callback CirculateNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(CirculateNotify, win, callback)
}

func (callback CirculateNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(CirculateNotifyEvent))
}

type CirculateRequestFun func(xu *xgbutil.XUtil, event CirculateRequestEvent)

func (callback CirculateRequestFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(CirculateRequest, win, callback)
}

func (callback CirculateRequestFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(CirculateRequestEvent))
}

type PropertyNotifyFun func(xu *xgbutil.XUtil, event PropertyNotifyEvent)

func (callback PropertyNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(PropertyNotify, win, callback)
}

func (callback PropertyNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(PropertyNotifyEvent))
}

type SelectionClearFun func(xu *xgbutil.XUtil, event SelectionClearEvent)

func (callback SelectionClearFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(SelectionClear, win, callback)
}

func (callback SelectionClearFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(SelectionClearEvent))
}

type SelectionRequestFun func(xu *xgbutil.XUtil, event SelectionRequestEvent)

func (callback SelectionRequestFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(SelectionRequest, win, callback)
}

func (callback SelectionRequestFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(SelectionRequestEvent))
}

type SelectionNotifyFun func(xu *xgbutil.XUtil, event SelectionNotifyEvent)

func (callback SelectionNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(SelectionNotify, win, callback)
}

func (callback SelectionNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(SelectionNotifyEvent))
}

type ColormapNotifyFun func(xu *xgbutil.XUtil, event ColormapNotifyEvent)

func (callback ColormapNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ColormapNotify, win, callback)
}

func (callback ColormapNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ColormapNotifyEvent))
}

type ClientMessageFun func(xu *xgbutil.XUtil, event ClientMessageEvent)

func (callback ClientMessageFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(ClientMessage, win, callback)
}

func (callback ClientMessageFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(ClientMessageEvent))
}

type MappingNotifyFun func(xu *xgbutil.XUtil, event MappingNotifyEvent)

func (callback MappingNotifyFun) Connect(xu *xgbutil.XUtil, win xgb.Id) {
    xu.AttachCallback(MappingNotify, win, callback)
}

func (callback MappingNotifyFun) Run(xu *xgbutil.XUtil, event interface{}) {
    callback(xu, event.(MappingNotifyEvent))
}

