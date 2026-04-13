#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

void setFullScreen(bool full, void *win) {
    NSWindow *window = (NSWindow*)win;

    NSUInteger masks = [window styleMask];
    bool isFull = masks & NSWindowStyleMaskFullScreen;
    if (isFull == full) {
        return;
    }

    [window toggleFullScreen:NULL];
}

// isFullScreen queries the NSWindow's current style mask and reports whether
// the window is presently in native fullscreen. This is the source of truth
// on macOS because the user can enter/exit fullscreen via the green traffic
// light (or cmd-ctrl-F) without Fyne's Window.SetFullScreen being called,
// which would leave the internal fullScreen flag stale.
bool isFullScreen(void *win) {
    NSWindow *window = (NSWindow*)win;
    return ([window styleMask] & NSWindowStyleMaskFullScreen) != 0;
}
