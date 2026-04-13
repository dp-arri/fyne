//go:build darwin

package glfw

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit

#import <stdbool.h>

void setFullScreen(bool full, void *window);
bool isFullScreen(void *window);
*/
import "C"

import (
	"runtime"

	"fyne.io/fyne/v2/driver"
)

// assert we are implementing driver.NativeWindow
var _ driver.NativeWindow = (*window)(nil)

func (w *window) RunNative(f func(any)) {
	context := driver.MacWindowContext{}
	if v := w.view(); v != nil {
		context.NSWindow = uintptr(v.GetCocoaWindow())
	}

	f(context)
}

func (w *window) doSetFullScreen(full bool) {
	if runtime.GOOS == "darwin" {
		win := w.view().GetCocoaWindow()
		C.setFullScreen(C.bool(full), win)
		return
	}
}

// isNativeFullScreen returns the NSWindow's current fullscreen state,
// which may diverge from w.fullScreen if the user toggled fullscreen via
// the macOS window controls rather than via Window.SetFullScreen.
//
// We read w.viewport directly instead of going through view() because view()
// returns nil once w.closing is set, which happens before the GLFW window is
// actually destroyed. Callers that run during teardown (a final save from
// OnExitedForeground, for example) still need the real state at that moment;
// the underlying NSWindow remains valid until glfw.Terminate() tears it down.
func (w *window) isNativeFullScreen() bool {
	if w.viewport == nil {
		return w.fullScreen
	}
	win := w.viewport.GetCocoaWindow()
	if win == nil {
		return w.fullScreen
	}
	return bool(C.isFullScreen(win))
}
