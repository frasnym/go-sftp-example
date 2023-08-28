package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	// Open SFTP connection
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

	// Copy local file content to remote file
	localFile, err := os.Open("/path/to/local/file")
	if err != nil {
		log.Fatal("Failed to open local file: ", err)
	}
	defer localFile.Close()

	remoteFile, err := client.Create("/path/to/remote/file")
	if err != nil {
		log.Fatal("Failed to create remote file: ", err)
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		log.Fatal("Failed to upload file: ", err)
	}
}
