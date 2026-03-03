# 🦆 Ducky-clip v0.1.0

> **Your personal snippet pond.** A powerful, cross-platform CLI snippet manager for developers who keep repeating the same commands.

---

## ✨ Features

- **Quick Save:** Store commands or code blocks with a simple alias.
- **Searchable Pond:** Browse your library with the `list` command.
- **Cross-Platform:** Native binaries for **Linux** and **Windows**.
- **Fast & Local:** Powered by a local SQLite database (no cloud required).
- **Custom ASCII Art:** Because every CLI tool needs a duck.

---

## 🚀 Installation

### 🐧 Linux (Fedora, Ubuntu, Debian, etc.)

**Manual Installation:**

1. Download the `ducky` binary from the [Releases](#) section.
2. Grant execution permissions:
    ```bash
    chmod +x ducky
    ```
3. Move the binary to your path (renaming it to just `ducky` for convenience):
    ```bash
    sudo mv ducky /usr/local/bin/ducky
    ```
4. Verify the installation:
    ```bash
    ducky --help
    ```

---

## 📂 Data Storage & Persistence

`ducky` values your privacy and speed. All your snippets and settings are stored locally in your user's home directory to ensure they persist between updates:

- **Folder Path:** `~/.ducky-clip/`
- **Database File:** `~/.ducky-clip/snippets.db`

> **Note:** To backup your snippets or move them to a new machine, simply copy the `.ducky-clip` folder.

---

## 🗑️ Uninstallation

To completely remove `ducky` from your system, follow these steps:

1. **Remove the binary:**

    ```bash
    sudo rm /usr/local/bin/ducky
    ```

2. **(Optional) Wipe your data:**
   By default, your saved snippets remain in `~/.ducky-clip`. If you wish to delete all your data as well, run:
    ```bash
    rm -rf ~/.ducky-clip
    ```

---

## 🛠️ Development

If you want to build the project from source:

```bash
# Clone the repository
git clone [https://github.com/tu-usuario/ducky-clip.git](https://github.com/tu-usuario/ducky-clip.git)
cd ducky-clip

# Build the binary
go build -o build/ducky main.go
```
