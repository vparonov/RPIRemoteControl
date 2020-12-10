package sshcontroller

import (
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

//SSHController - ...........
type SSHController struct {
	host       string
	user       string
	sshKeyFile string
	connection *ssh.Client
	session    *ssh.Session
	stdin      io.Reader
	stdout     io.Writer
	stderr     io.Writer
}

//New - ...........
func New(host string, user string, sshKeyFile string) *SSHController {
	return &SSHController{
		host:       host,
		user:       user,
		sshKeyFile: sshKeyFile,
		stdin:      os.Stdin,
		stdout:     os.Stdout,
		stderr:     os.Stderr,
	}
}

//Open - --------
func (sshc *SSHController) Open() error {
	sshConfig, err := newClientConfig(sshc.user, sshc.sshKeyFile)

	if err != nil {
		return err
	}

	connection, err := ssh.Dial("tcp", sshc.host, sshConfig)

	if err != nil {
		return err
	}

	sshc.connection = connection

	session, err := connection.NewSession()

	if err != nil {
		return err
	}

	sshc.session = session

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	go io.Copy(stdin, sshc.stdin)

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	go io.Copy(sshc.stdout, stdout)

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	go io.Copy(sshc.stderr, stderr)

	return nil
}

//Close - .........
func (sshc *SSHController) Close() error {

	err := sshc.session.Close()
	if err != nil {
		return err
	}

	return sshc.connection.Close()
}

//Run - ...........
func (sshc *SSHController) Run(cmd string) error {
	return sshc.session.Run(cmd)
}

//WithStdin - .........
func (sshc *SSHController) WithStdin(stdin io.Reader) *SSHController {
	sshc.stdin = stdin

	return sshc
}

//WithStdout - .........
func (sshc *SSHController) WithStdout(stdout io.Writer) *SSHController {
	sshc.stdout = stdout

	return sshc
}

//WithStderr - .........
func (sshc *SSHController) WithStderr(stderr io.Writer) *SSHController {
	sshc.stderr = stderr

	return sshc
}

func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	//key, err := ssh.ParsePrivateKeyWithPassphrase(buffer, []byte("#####"))
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func newClientConfig(user string, file string) (*ssh.ClientConfig, error) {
	publicKeyFile, err := publicKeyFile(file)
	if err != nil {
		return nil, err
	}

	return &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			publicKeyFile,
		},
	}, nil
}
