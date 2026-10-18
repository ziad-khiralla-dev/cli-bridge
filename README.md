# 🚀 CLI-to-Telegram Notification Bridge

A lightweight, secure, and production-ready Go utility that executes local command-line commands and instantly relays execution outputs, duration metrics, and status reports directly to your personal Telegram account via a custom Telegram Bot.

---

## ✨ Features

- **Local Command Execution:** Safely executes standard CLI commands using Go's standard `os/exec` package.
- **Performance Measurement:** Tracks precise command execution duration using high-precision stopwatch metrics.
- **Telegram Push Notifications:** Delivers structured success/failure reports directly to your mobile or desktop device.
- **Environment Security:** Protects sensitive credentials (Bot Token & Chat ID) using environment variables via `godotenv`.
- **Clean Modular Architecture:** Implemented with a clean directory structure separating the core execution engine (`internal/executor`) from the communication layer (`internal/notifier`).

---

## 📂 Project Structure

```text
cli-bridge/
├── cmd/
│   └── bridge/
│       └── main.go         # Application entry point and orchestrator
├── internal/
│   ├── executor/           # Command execution and timing logic
│   └── notifier/           # Telegram Bot API integration
├── .env                    # Local secret configuration (Git ignored)
├── .gitignore              # Version control exclusion rules
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

---

## 🛠️ Prerequisites

- **Go:** Version 1.18 or higher installed on your system.
- **Telegram Bot Token:** Obtained via Telegram's [@BotFather](https://t.me/BotFather).
- **Telegram Chat ID:** Verified via [@userinfobot](https://t.me/userinfobot) or Telegram Updates API.

---

## ⚙️ Setup & Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Ziad-khiralla-dev/cli-bridge.git
   cd cli-bridge
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Configure Environment Variables:**
   Create a `.env` file in the root directory of the project and populate your credentials:
   ```env
   TELEGRAM_BOT_TOKEN=your_bot_token_here
   TELEGRAM_CHAT_ID=your_chat_id_here
   ```

4. **Verify `.gitignore` Security:**
   Ensure your `.env` file is excluded from version control:
   ```text
   .env
   *.exe
   ```

---

## 🚀 Usage

Run the tool by passing any target CLI command as a string argument:

```bash
go run cmd/bridge/main.go "dir"
```

### Example Telegram Payload Formats

- **Success Output:**
  ```text
  [ Command Executed Successfully! ]

   --> Command: ' dir ' 
   --> Output: ' [Command output contents here] ' 
   --> Duration: ' 15.2ms '
  ```

- **Failure Output:**
  ```text
  [ Command Execution Failed! ]

   --> Command: ' invalid_cmd ' 
   --> Error: ' exec: "invalid_cmd": executable file not found in %PATH% ' 
   --> Duration: ' 1.1ms '
  ```

---

## 🔒 Security Practices

To adhere to security standards, credentials such as the Telegram Bot Token and Chat ID are kept out of source code and version control by relying on `.env` files loaded dynamically at runtime via `github.com/joho/godotenv`.

---

## 👨‍💻 Author

Developed by **Ziad Khiralla**  
- GitHub: [@ziad-khiralla-dev](https://github.com/ziad-khiralla-dev)