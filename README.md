# syslite

**`syslite`** is a lightweight Go module designed to quickly retrieve essential system information with minimal overhead. This module is tailored for developers who need fast, simple, and intuitive functions to check system details such as OS name, distribution, kernel version, and architecture. With built-in Boolean functions, **`syslite`** simplifies cross-platform development and CI/CD pipeline tasks.

---

## Comments
I'm converting this into single binary which works on cross platform to extract operating system information, return in yaml, json and env.
This binary will help to my other workflows (ansible, GHA) or code to make decision. 
Eg. Running this binary as part of worklow where it help to get more info and inject into GITHUB ENV of current runtime which will be useful for next steps decission making (if ENV.KERNEL_VERSION == 3.X)

## Why `go-syslite`?

- **Lightweight & Fast**: A minimalistic library with zero dependencies, optimized for high performance and low memory usage.
- **Cross-Platform Support**: Works seamlessly across Linux, Windows, macOS, and BSD systems.
- **Simple Boolean Functions**: Quickly check system properties without parsing system files or dealing with complex logic.
- **Optimized for CI/CD & Shell**: Use in CI pipelines, shell scripts, and other automated tools for efficient system detection.
- **Extendable**: Easily extend or customize the module for your own use cases.

---

## Features

- **System Information Retrieval**:
  - `os_name`: Get the name of the operating system.
  - `distro`: Detect Linux distributions.
  - `kernel_version`: Retrieve the kernel version.
  - `architecture`: Get the system architecture.

- **Boolean Utility Functions**:
  - `isLinux()`: Checks if the system is Linux.
  - `isWindows()`: Checks if the system is Windows.
  - `is64()`: Checks if the system is 64-bit.
  - `isDebian()`: Checks if the system is a Debian-based Linux distro.
  - `isUbuntu()`: Checks if the system is Ubuntu.
  - `isCentOS()`: Checks if the system is CentOS.
  - `isDarwin()`: Checks if the system is macOS.
  - `isFreeBSD()`: Checks if the system is FreeBSD.
  - `is86_64()`: Checks if the system architecture is x86_64.
  - `isAMD64()`: Checks if the system architecture is AMD64.

---

## Installation

You can install **`go-syslite`** using Go modules:

```bash
go get github.com/thineshsubramani/go-syslite
```

## Contact
