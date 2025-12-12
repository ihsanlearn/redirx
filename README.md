<div align="center">
<pre>
__________           .___.__      ____  ___
\______   \ ____   __| _/|__|_____\   \/  /
 |       _// __ \ / __ | |  \_  __ \     / 
 |    |   \  ___// /_/ | |  ||  | \/     \ 
 |____|_  /\___  >____ | |__||__| /___/\  \
        \/     \/     \/                \_/
</pre>  
<a href="https://github.com/ihsanlearn/redirx"><img src="https://img.shields.io/badge/RedirX-Open%20Redirect%20Scanner-blue?style=for-the-badge&logo=go" alt="RedirX"></a>
<div align="center">
  <img src="https://img.shields.io/badge/Go-1.24%2B-blue?style=for-the-badge&logo=go" alt="Go 1.24+">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License: MIT">
  <img src="https://img.shields.io/badge/Contribution-Welcome-red?style=for-the-badge" alt="Contribution: Welcome">
</div>
</div>

---

**RedirX** is a fast and concurrent tool designed to detect Open Redirect vulnerabilities. It supports single URL scanning, list scanning, and pipelining from other tools. With customizable payloads and multi-threading support, it is an essential tool for bug bounty hunters and penetration testers.

## Features

- ⚡ **Fast & Concurrent**: Multi-threaded scanning for high performance.
- 🎯 **Smart Detection**: Detects open redirects based on location headers.
- 🔧 **Custom Payloads**: Support for custom payloads and payload lists.
- 🛡️ **Bypass Techniques**: Includes HTTP Parameter Pollution (HPP) and other bypass methods.
- 🍾 **Pipeline Friendly**: Reads input from `stdin` for easy integration with other tools.
- 📊 **Flexible Output**: Save results to a file or print to stdout.

## Installation

Ensure you have **Go 1.24+** installed.

```bash
go install github.com/ihsanlearn/redirx/cmd/redirx@latest
```

## Usage

```bash
redirx -h
```

This will display help for the tool. Here are all the switches it supports.

| Flag                 | Description                                    | Example                  |
| -------------------- | ---------------------------------------------- | ------------------------ |
| **Input**            |                                                |                          |
| `-u, -url`           | Target URL for scanning (comma separated)      | `-u https://example.com` |
| `-l, -list`          | File containing list of target URLs            | `-l urls.txt`            |
| **Configuration**    |                                                |                          |
| `-t, -threads`       | Number of concurrent threads (default 25)      | `-t 50`                  |
| `-T, -timeout`       | Request timeout in seconds (default 10)        | `-T 5`                   |
| `-p, -payload`       | Custom payload for scanning                    | `-p https://evil.com`    |
| `-pl, -payload-list` | File containing list of custom payloads        | `-pl payloads.txt`       |
| `-H, -hpp`           | Enable HTTP Parameter Pollution                | `-H`                     |
| `-verify-ssl`        | Enable SSL verification (default false)        | `-verify-ssl`            |
| `-rl, -rate-limit`   | Maximum requests per second (default 10)       | `-rl 20`                 |
| `-d, -delay`         | Delay between requests in milliseconds         | `-d 100`                 |
| `-k, -keep-alive`    | Enable keep-alive connections (default true)   | `-k`                     |
| **Output**           |                                                |                          |
| `-o, -output`        | File for saving scan results                   | `-o results.txt`         |
| **Optimization**     |                                                |                          |
| `-s, -silent`        | Silent mode (only print found vulnerabilities) | `-s`                     |
| `-v, -verbose`       | Verbose mode (print error & debug messages)    | `-v`                     |
| `-V, -version`       | Display application version                    | `-V`                     |

## Examples

**1. Scan a single URL:**

```bash
redirx -u "https://example.com?redirect=test"
```

**2. Scan a list of URLs from a file:**

```bash
redirx -l urls.txt
```

**3. Use a custom payload:**

```bash
redirx -u "https://example.com?next=test" -p "https://evil.com"
```

**4. Pipeline with other tools (e.g., waybackurls):**

```bash
echo "https://example.com" | waybackurls | gf redirect | redirx
```

**5. Save results to a file:**

```bash
redirx -l urls.txt -o vulnerable.txt
```
