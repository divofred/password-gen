# Password Generator CLI

This is a command-line tool to generate secure, random passwords using cryptographic randomness in Go. The password generator allows you to customize the length of the password, and whether it includes digits, special characters, and a combination of uppercase and lowercase letters.

## Features
- Generate secure random passwords with a specified length.
- Optionally include digits and special characters.
- Ensure at least one uppercase letter and one lowercase letter in the password.
- Cryptographically secure random number generation using the `crypto/rand` package.

## Requirements
- Go 1.18 or higher
- Cobra CLI library

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/password-gen.git
    cd password-gen
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Build the application:
    ```bash
    go build
    ```

4. Optionally, install the application globally:
    ```bash
    go install
    ```

## Usage

### Command Syntax
```bash
password-gen generate --help
```
