package sshcontroller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	s := New("192.168.1.238:22", "pi", "/Users/vangel/sshKey/id_rsa")

	err := s.Open()

	assert.Nil(t, err)

	err = s.Run("ls -l")
	assert.Nil(t, err)

	err = s.Close()

	assert.Nil(t, err)
	//err = session.Run("ls -l")
	//sshConfig, err := newClientConfig("pi", "/Users/vangel/sshKey/id_rsa")

	//	if err != nil {
	//		return err
	//	}

	//	connection, err := ssh.Dial("tcp", "192.168.1.238:22", sshConfig)
	//assert.Nil(t, err)
}
