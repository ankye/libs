package grpc

import (
	log "github.com/gonethopper/libs/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientOptions struct {
	Address string
	Secure  *ClientSecureOptions
	Auth    credentials.PerRPCCredentials
}

type Client struct {
	*grpc.ClientConn
	options *ClientOptions
}

type ClientSecureOptions struct {
	File string
	Name string
}

//NewClient create grpc client, if connection is refused, then return error
func NewClient(options *ClientOptions) (*Client, error) {
	c := &Client{}
	opts := make([]grpc.DialOption, 0)
	if options.Secure == nil {
		opts = append(opts, grpc.WithInsecure())
	} else {
		creds, err := credentials.NewClientTLSFromFile(options.Secure.File, options.Secure.Name)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	if options.Auth != nil {
		grpc.WithPerRPCCredentials(options.Auth)
	}

	conn, err := grpc.Dial(options.Address, opts...)
	if err != nil {
		return nil, err
	}
	c.ClientConn = conn

	state := conn.GetState()
	log.Info("rpc connect:[%s] status[%s]", options.Address, state.String())
	return c, nil

}
