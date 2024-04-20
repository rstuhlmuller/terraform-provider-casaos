package conns

import "context"

type CasaOSClient struct {
	Host     string
	Username string
	Password string
}

func (c *CasaOSClient) CredentialsProvider(context.Context) (*CasaOSClient, error) {
	return c, nil
}
