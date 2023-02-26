package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	MSG_SIZE = 65536
	DEBUG    = false
)

var Input *bufio.Reader

func ternary(condition bool, iftrue any, iffalse any) any {
	if condition {
		return iftrue
	} else {
		return iffalse
	}
}

func send(message []byte, conn net.Conn) []byte {
	var err error

	if len(message) > MSG_SIZE {
		mError("Request is too big (", len(message), "/", MSG_SIZE, ")")
		return []byte("")
	}

	mDebug("Sending", string(message))
	_, err = conn.Write(message)
	if err != nil {
		mError("Error sending request:", err.Error())
		os.Exit(1)
	}

	response := make([]byte, MSG_SIZE)
	_, err = conn.Read(response)
	if err != nil {
		mError("Error receiving response:", err.Error())
		os.Exit(1)
	}

	return response
}

func main() {
	Input = bufio.NewReader(os.Stdin)
	prepare_configdir()
	var err error

	_password := password_read()

	iplist := config_read()
	mOk("Config file loaded")
	// mDebug(iplist, configdir)

	for i := 0; i < len(iplist); i++ {
		record := iplist[i]
		fmt.Printf("%d: %s:%d %s\n", i+1, record.Ip, record.Port, ternary(record.Secured, "Secured (TLS)", "Unsecured"))
	}

	var chosenrecord int
	for {
		_chosenrecord := input("Choose IP record to use (or new)", "1")
		if _chosenrecord == "new" {
			iplist = append(iplist, IP{})
			chosenrecord = len(iplist)
		} else {
			chosenrecord, err = strconv.Atoi(_chosenrecord)
			if err != nil {
				mError(err.Error() + ". You are supposed to type a number smaller than", len(iplist), "or 'new'")
				continue
			}
			if chosenrecord > len(iplist) {
				mError("You are supposed to type a number smaller than", len(iplist), "or 'new'")
				continue
			}
		}
		break
	}
	chosenrecord--

	host := input("Enter hostname", iplist[chosenrecord].Ip)
	port := input("Enter port", fmt.Sprint(iplist[chosenrecord].Port))
	secure := strings.ToLower(input("Secured? y/n", fmt.Sprint(ternary(iplist[chosenrecord].Secured, "yes", "no"))))[0] == 'y'
	password := input_p("Enter password (or 'no' for nologin)", fmt.Sprint(ternary(_password == "", "no", _password)))
	if password == "no" {
		mInfo("/registeradmin will not be executed")
	}

	_port, err := strconv.Atoi(port)
	if err != nil {
		mError("Given port was not a number")
		os.Exit(1)
	}

	iplist[chosenrecord] = IP{
		Ip:      host,
		Port:    _port,
		Secured: secure,
	}
	config_save(iplist)
	password_save(password)
	mOk("Config files saved")

	addr := net.JoinHostPort(host, port)
	mInfo("Connecting to", addr, "using", ternary(secure, "TCP TLS", "TCP"))

	var conn net.Conn
	if secure {
		conn, err = tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		conn, err = net.Dial("tcp", addr)
	}

	if err != nil {
		mError("Error connecting to server:", err.Error())
		os.Exit(1)
	} else {
		mOk("Successfully connected to server. Type /help for help. Type :q to exit.")
	}
	defer conn.Close()

	anotherLoginAttempt := false
	for {
		if password != "no" {
			message := []byte("/registeradmin "+password)
			response := string(send(message, conn))
			
			if strings.Trim(response, "\x00") == "[WRONG_PASSWORD]" {
				mError("Wrong password")
				password = input_p("Enter password (or 'no' for nologin)", fmt.Sprint(ternary(_password == "", "no", _password)))
				anotherLoginAttempt = true

				conn.Close()
				if secure {
					conn, err = tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true})
				} else {
					conn, err = net.Dial("tcp", addr)
				}
			
				if err != nil {
					mError("Error reconnecting to server:", err.Error())
					os.Exit(1)
				} else {
					mOk("Successfully reconnected to server. Type /help for help. Type :q to exit.")
				}

				continue
			}

			if strings.Trim(response, "\x00") == "[REGISTER_SUCCESS]"{
				mOk("Logged in successfully")
				if anotherLoginAttempt {
					password_save(password)
					mOk("Config files saved")
				}
			} else {
				mError("Unknown error: ", response, ". You are not logged in")
			}
			break
		} else {
			break
		}
	}

	for {
		fmt.Print("> ")
		_in, err := Input.ReadString('\n')
		if err != nil {
			mError("Error reading message:", err.Error())
			continue
		}

		line := strings.Trim(_in, NEWLINE)
		message := []byte(line)
		mDebug(len(message), ";", message)

		if line == ":q" {
			break
		}

		response := string(send(message, conn))
		fmt.Println(strings.Trim(response, "\x00"))
	}
	mInfo("Exitting")
}
