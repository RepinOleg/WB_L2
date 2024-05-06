package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Утилита telnet

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

type telnet struct {
	timeout *time.Duration
	address string
	port    string
}

func main() {
	var t telnet
	err := t.parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = t.connection()
	if err != nil {
		log.Fatal(err)
	}

}

func (t *telnet) parseArgs() error {
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout for connection")
	flag.Parse()

	if flag.NArg() != 2 {
		return fmt.Errorf("not enough args")
	}

	t.address = flag.Arg(0)
	t.port = flag.Arg(1)
	t.timeout = timeout

	return nil
}

func (t *telnet) connection() error {
	fmt.Println("Trying", t.address+"...")
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(t.address, t.port), *t.timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Connected to", t.address)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nConnection closed.")
		os.Exit(0)
	}()

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)

	return nil
}
