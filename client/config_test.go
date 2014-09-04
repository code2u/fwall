package main

import (
	"testing"
)

var testConf string = `
{
    "server":"example.com",
    "server_port":443,
    "master_key":"foobar",
    "method":"aes-cfb",
    "tunnel":"raw",

    "local_address":"127.0.0.1",
    "local_port":1080,
    "username":"test",
    "password":"bzrfoo"
}
`

func Test_doParseConfigFile(t *testing.T) {
	cf, err := doParseConfigFile([]byte(testConf))
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if cf.Server != "example.com" {
		t.Errorf("server wrong: %s", cf.Server)
	}

	if cf.ServerPort != 443 {
		t.Errorf("server_port wrong: %d", cf.ServerPort)
	}

	if cf.Method != "aes-cfb" {
		t.Errorf("method wrong: %s", cf.Method)
	}
}
