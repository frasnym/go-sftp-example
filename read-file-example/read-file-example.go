package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Create SFTP Connection
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "sftp.example.com:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("Failed to create SFTP client: ", err)
	}
	defer client.Close()

	bufioNewReaderReadString(client) // Method 1
	bufioNewScannerScanLines(client) // Method 2
	readFull(client)                 // Method 3
}

func bufioNewReaderReadString(client *sftp.Client) {
	remoteFile, err := client.Open("/path/to/remote/file")
	if err != nil {
		log.Fatal("Failed to open remote file: ", err)
	}
	defer remoteFile.Close()

	reader := bufio.NewReader(remoteFile)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read remote file: ", err)
		}
		// Process the line here
		// ...
		fmt.Printf("line: %v\n", line)
	}
}

func bufioNewScannerScanLines(client *sftp.Client) {
	remoteFile, err := client.Open("/path/to/remote/file")
	if err != nil {
		log.Fatal("Failed to open remote file: ", err)
	}
	defer remoteFile.Close()

	scanner := bufio.NewScanner(remoteFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line here
		// ...
		fmt.Printf("line: %v\n", line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read remote file: ", err)
	}
}

func readFull(client *sftp.Client) {
	remoteFile, err := client.Open("/path/to/remote/file")
	if err != nil {
		log.Fatal("Failed to open remote file: ", err)
	}
	defer remoteFile.Close()

	bufferSize := 4096 // Adjust the buffer size as needed
	buffer := make([]byte, bufferSize)

	var remainder []byte // Keep track of any remaining bytes from the previous buffer

	for {
		bytesRead, err := io.ReadFull(remoteFile, buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Failed to read remote file: ", err)
		}
		// Combine the remainder from the previous buffer with the current buffer
		combinedBuffer := append(remainder, buffer[:bytesRead]...)

		// Split the combined buffer into lines
		lines := bytes.Split(combinedBuffer, []byte{'\n'})

		// Process each line
		for _, line := range lines {
			// Process the line here
			// ...
			fmt.Printf("line: %v\n", line)
		}

		// Keep track of any remaining bytes that don't form a complete line
		remainder = lines[len(lines)-1]
	}

	// Process any remaining bytes that don't form a complete line
	if len(remainder) > 0 {
		// Process the remainder here
		// ...
	}
}
