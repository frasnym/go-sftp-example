package main

import (
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func AuthWithPublicKey() {
	// Read key file
	key, err := os.ReadFile("/path/to/private/key")
	if err != nil {
		log.Fatal("Failed to read private key: ", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("Failed to parse private key: ", err)
	}

	// Open SFTP connection
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
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
}
