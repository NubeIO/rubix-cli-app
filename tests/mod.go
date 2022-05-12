package main

import (
	"fmt"
	"github.com/grid-x/modbus"
	"log"
	"os"
	"time"
)

func main() {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler("192.168.15.202:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveID = 0xFF
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()

	fmt.Println(err)

	client := modbus.NewClient(handler)
	results, err := client.ReadDiscreteInputs(1, 2)
	fmt.Println(results)
	//results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	//results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})

}
