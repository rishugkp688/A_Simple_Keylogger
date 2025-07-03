
# Windows Low-Level Keyboard Hook in Go

This project implements a low-level global keyboard hook using Go and Windows API calls. It logs key presses by intercepting keyboard events at the system level using `user32.dll`.

> ⚠️ **Warning:** This code demonstrates how to create a keylogger. It should only be used for educational or authorized purposes. Running or distributing keyloggers without consent is illegal and unethical.

---

## 🧠 How It Works

- This Go program sets a **low-level keyboard hook** using the Windows API (`SetWindowsHookExW`) to listen for `WM_KEYDOWN` events.
- When a key is pressed, it reads the `scanCode`, converts it to a human-readable key name using `GetKeyNameTextW`, and prints it.
- The hook runs in a loop using `GetMessageW` to process Windows messages continuously.

### Key Functions

- `SetWindowsHookExW`: Installs the keyboard hook.
- `CallNextHookEx`: Passes the hook information to the next hook procedure.
- `GetMessageW`: Retrieves messages from the thread’s message queue.
- `GetKeyNameTextW`: Converts a scan code into a readable string.

---

## 🧪 Example Output

```
Key pressed: A
Key pressed: B
Key pressed: C
```

---

## 🛠 How to Build & Run

### ✅ Requirements

- Go 1.24 or later
- Windows OS (tested on Windows 10/11)

### 🧾 Instructions

1. **Save the Code**

   Save the Go code to a file named `main.go`.

2. **Open Command Prompt / PowerShell**

3. **Build the Binary**

   ```sh
   go build -o keylogger.exe main.go
   ```

4. **Run the Program**

   Run the compiled binary in a terminal **as administrator**:

   ```sh
   keylogger.exe
   ```

   > 📢 The hook will now capture and print every key pressed on the keyboard.

---

## ❌ Troubleshooting

- **"Failed to set hook" error**: This usually means the program doesn't have enough privileges. Try running it as administrator.
- **Antivirus blocking it**: Due to its behavior, some antivirus programs may flag the binary. Only run it in a safe, test environment.

---

## ⚖️ Legal Disclaimer

This tool is for educational and authorized use **only**. Unauthorized use to monitor keystrokes can be considered illegal under various computer crime laws.

---

## 📄 License

This project is released under the MIT License.
