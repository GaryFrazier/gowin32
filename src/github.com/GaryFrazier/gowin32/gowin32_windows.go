package gowin32

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const NULLPTR = 0x0
const WINDOW_CLASS = "MyWindowClass"

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	createWindow        = user32.NewProc("CreateWindowExW")
	defWindowProc       = user32.NewProc("DefWindowProcW")
	dispatchMessage     = user32.NewProc("DispatchMessageW")
	translateMessage    = user32.NewProc("TranslateMessage")
	beginPaint          = user32.NewProc("BeginPaint")
	endPaint            = user32.NewProc("EndPaint")
	postQuitMessage     = user32.NewProc("PostQuitMessage")
	registerClass       = user32.NewProc("RegisterClassExW")
	showWindow          = user32.NewProc("ShowWindow")
	drawTextW           = user32.NewProc("DrawTextW")
	releaseCapture      = user32.NewProc("ReleaseCapture")
	getCursorPos        = user32.NewProc("GetCursorPos")
	screenToClient      = user32.NewProc("ScreenToClient")
	getMessage          = user32.NewProc("GetMessageW")
	procGetModuleHandle = kernel32.NewProc("GetModuleHandleW")
)

func getModuleHandle() (uintptr, error) {
	ret, _, err := procGetModuleHandle.Call(0)
	if ret == 0 {
		return 0, err
	}
	return ret, nil
}

func registerWindowClass() error {
	wndClass := WNDCLASSEX{
		cbSize:        uint32(unsafe.Sizeof(WNDCLASSEX{})),
		style:         CS_HREDRAW | CS_VREDRAW,
		lpfnWndProc:   syscall.NewCallback(WndProc),
		cbClsExtra:    0,
		cbWndExtra:    0,
		hInstance:     0,
		hIcon:         0,
		hCursor:       0,
		hbrBackground: 0,
		lpszMenuName:  nil,
		lpszClassName: (*uint16)(unsafe.Pointer(StringToUTF16Ptr(WINDOW_CLASS))),
		hIconSm:       0,
	}

	ret, _, err := registerClass.Call(uintptr(unsafe.Pointer(&wndClass)))
	if ret == 0 {
		return err
	}
	return nil
}

func WndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		fmt.Println("Window Created!")
	case WM_PAINT:
		var ps PAINTSTRUCT
		hdc, _, _ := beginPaint.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&ps)))
		defer endPaint.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&ps)))
		text := uintptr(unsafe.Pointer(StringToUTF16Ptr("TEXT")))
		rect := RECT{Left: 20, Top: 20}
		drawTextW.Call(hdc, text, 0, uintptr(unsafe.Pointer(&rect)), DT_SINGLELINE|DT_CENTER|DT_VCENTER)
	case WM_LBUTTONDOWN:
		fmt.Println("WM_LBUTTONDOWN received")
		// Release the capture to allow dragging the window
		releaseCapture.Call()
		// Get the cursor position
		var pt POINT
		getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
		// Convert screen coordinates to client coordinates
		screenToClient.Call(hWnd, uintptr(unsafe.Pointer(&pt)))
		fmt.Printf("Mouse clicked at (%d, %d) in client coordinates\n", pt.x, pt.y)
	case WM_CLOSE:
		postQuitMessage.Call(0)
		os.Exit(0)
	}
	r1, _, _ := defWindowProc.Call(uintptr(hWnd), uintptr(msg), wParam, lParam)

	return r1
}

// CreateWindow uses the CreateWindowEx Windows API Call to create and return a Handle to a new window https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexa
func CreateWindow(windowName string) {
	hInstance, err := getModuleHandle()
	if err != nil {
		fmt.Println("Error getting module handle:", err)
		return
	}

	classErr := registerWindowClass()
	if classErr != nil {
		fmt.Println("Error registering window class:", classErr)
		return
	}

	hwind, _, createWinErr := createWindow.Call(
		0,
		uintptr(unsafe.Pointer(StringToUTF16Ptr(WINDOW_CLASS))), //LPCSTR    lpClassName
		uintptr(unsafe.Pointer(StringToUTF16Ptr(windowName))),   //LPCSTR    lpWindowName
		WS_OVERLAPPEDWINDOW|WS_VISIBLE|WS_POPUP,                 //DWORD     dwStyle https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
		0,         //int       X
		0,         //int       Y
		800,       //int       nWidth
		600,       //int       nHeight
		NULLPTR,   //HWND      hWndParent
		NULLPTR,   //HMENU     hMenu
		hInstance, //HINSTANCE hInstance
		NULLPTR,   //LPVOID    lpParam
	)
	if createWinErr != syscall.Errno(0) {
		return
	}

	showWindow.Call(hwind, 1)

	// Message loop
	var msg MSG
	for {
		msgR1, _, _ := getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
		if msgR1 == 0 {
			break
		}
		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}
}
