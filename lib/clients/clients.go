package clients

import (
	"net"

	"golang.org/x/xerrors"
)

type Client struct {
	Name       string
	PrivateKey string
	IP         net.IP
}

func NewClient(name, privateKey, IP string) (Client, error) {
	c := Client{name, privateKey, net.ParseIP(IP)}

	if c.Name == "" {
		return c, xerrors.New("name cannot be empty")
	}
	if c.PrivateKey == "" {
		return c, xerrors.New("privateKey cannot be empty")
	}
	if IP == "" {
		return c, xerrors.New("IP cannot be empty")
	}
	if c.IP == nil {
		return c, xerrors.New("IP or invalid")
	}

	return c, nil

}
