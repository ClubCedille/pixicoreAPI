package sshclient

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"github.com/tmc/scp"
)

//SSHClient used for ssh client
type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

//RunCommand run ssh command in the remote server and retrun output
func (client *SSHClient) RunCommand(command string) (string, error) {
	var (
		session *ssh.Session
		err     error
	)

	if session, err = client.newSession(); err != nil {
		return "", err
	}
	defer session.Close()

	out, err := session.CombinedOutput(command)
	return string(out), err
}

func (client *SSHClient) newSession() (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	modes := ssh.TerminalModes{
		// ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		return nil, fmt.Errorf("request for pseudo terminal failed: %s", err)
	}

	return session, nil
}

//PublicKeyFile get public key with private key
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//Copy permit copy localFile to remote path using active ssh session on SSHClient
func (client SSHClient) Copy(sourceFilePath string, destinationFilePath string) error {
	session, err := client.newSession()
	if err != nil {
		return err
	}
	defer session.Close()

	err = scp.CopyPath(sourceFilePath, destinationFilePath, session)
	return err
}
