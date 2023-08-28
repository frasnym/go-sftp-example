# Connect to an SFTP server using public key authentication

To connect to an SFTP server using public key authentication in Golang, you can modify the SSH client configuration to use the appropriate public key file. Here's an example code snippet:

```
key, err := ioutil.ReadFile("/path/to/private/key")
if err != nil {
    log.Fatal("Failed to read private key: ", err)
}

signer, err := ssh.ParsePrivateKey(key)
if err != nil {
    log.Fatal("Failed to parse private key: ", err)
}

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

```

In this example, we're reading the private key file from disk and parsing it using the ssh.ParsePrivateKey function. We then create an ssh.PublicKeys authentication method using the parsed private key, and add it to the Auth field of the ssh.ClientConfig struct.

When establishing the SSH connection, the client will use the public key authentication method to authenticate with the server. Note that you'll need to have the corresponding public key added to the authorized keys on the SFTP server for this to work.

