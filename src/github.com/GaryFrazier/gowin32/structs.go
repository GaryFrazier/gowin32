package gowin32

type (
	COLORREF     uint32
	HBITMAP      HGDIOBJ
	HBRUSH       HGDIOBJ
	HDC          HANDLE
	HFONT        HGDIOBJ
	HGDIOBJ      HANDLE
	HENHMETAFILE HANDLE
	HPALETTE     HGDIOBJ
	HPEN         HGDIOBJ
	HRGN         HGDIOBJ
	CLIPFORMAT   uint16
)

type (
	ATOM          uint16
	HANDLE        uintptr
	HGLOBAL       HANDLE
	HINSTANCE     HANDLE
	LCID          uint32
	LCTYPE        uint32
	LANGID        uint16
	HMODULE       uintptr
	HWINEVENTHOOK HANDLE
	HRSRC         uintptr
)

type (
	HACCEL    HANDLE
	HCURSOR   HANDLE
	HDWP      HANDLE
	HICON     HANDLE
	HMENU     HANDLE
	HMONITOR  HANDLE
	HRAWINPUT HANDLE
	HWND      HANDLE
)

type WNDCLASSEX struct {
	cbSize        uint32
	style         uint32
	lpfnWndProc   uintptr
	cbClsExtra    int32
	cbWndExtra    int32
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  *uint16
	lpszClassName *uint16
	hIconSm       HICON
}

// Windows API Types
type (
	WNDPROC  uintptr
	LPARAM   uintptr
	WPARAM   uintptr
	UINT     uint32
	ULONG    uint32
	WORD     uint16
	DWORD    uint32
	BOOL     int32
	LRESULT  uintptr
	WNDCLASS struct {
		style         UINT
		lpfnWndProc   uintptr
		cbClsExtra    int32
		cbWndExtra    int32
		hInstance     HINSTANCE
		hIcon         HICON
		hCursor       HCURSOR
		hbrBackground HBRUSH
		lpszMenuName  *uint16
		lpszClassName *uint16
	}
	MSG struct {
		hWnd    HWND
		message UINT
		wParam  WPARAM
		lParam  LPARAM
		time    DWORD
		pt      POINT
	}
	POINT struct {
		x LONG
		y LONG
	}
	RECT struct {
		Left   LONG
		Top    LONG
		Right  LONG
		Bottom LONG
	}
	PAINTSTRUCT struct {
		hdc         HDC
		fErase      BOOL
		rcPaint     RECT
		fRestore    BOOL
		fIncUpdate  BOOL
		rgbReserved [32]byte
	}
	LONG int32
)
