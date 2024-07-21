
# FinalShell Activation Code Generator Documentation

## Introduction

This document provides instructions on using the FinalShell activation code generator written in Go. The program generates activation codes for a specific machine code, runs with administrator privileges, and supports both Windows and Linux/macOS systems.

## Features

- Requests and runs with administrator privileges.
- Accepts user input for the machine code.
- Generates activation codes for both old and new versions of FinalShell.
- Keeps the program running after displaying the activation codes, waiting for the user to press the Enter key to exit.

## Prerequisites

### Dependencies

Ensure that you have the Go programming environment installed. Install the required dependencies using the following command:

```bash
go get golang.org/x/crypto/sha3
```

## Code Structure

The program is structured into multiple files to support cross-platform functionality:

- `main.go`: Main entry point for Linux/macOS platforms.
- `main_windows.go`: Main entry point for Windows platforms.
- `common.go`: Contains shared logic and functions.

## How to Use

1. **Build the Program**: Build the program using the Go compiler.

    ```bash
    go build -o finalshell-activator
    ```

2. **Run the Program**: Execute the program. Ensure you run it with the necessary permissions (e.g., using `sudo` on Linux/macOS or as an administrator on Windows).

    ```bash
    ./finalshell-activator
    ```

3. **Input Machine Code**: When prompted, enter the machine code for which you want to generate activation codes.

    ```plaintext
    Please enter the machine code: <input your machine code>
    ```

4. **View Activation Codes**: The program will display the activation codes for both old and new versions of FinalShell.

    ```plaintext
    Input Machine Code: <your machine code>
    Version < 3.9.6 (Old Version)
    Premium: <premium code>
    Professional: <professional code>
    Version >= 3.9.6 (New Version)
    Premium: <premium code>
    Professional: <professional code>
    ```

5. **Exit the Program**: After viewing the activation codes, press the Enter key to exit the program.

    ```plaintext
    Press Enter to exit...
    ```

## Additional Notes

- **Administrator Privileges**: The program checks for administrator privileges at the start. If it is not running with the required permissions, it will attempt to restart itself with elevated privileges.
- **Cross-Platform Support**: The program includes platform-specific code to handle privilege elevation appropriately on both Windows and Linux/macOS.

This document provides a comprehensive guide to using the FinalShell activation code generator, ensuring you can generate the necessary activation codes effectively.
