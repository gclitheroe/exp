package main

import (
	tk "github.com/gclitheroe/exp/mtr/credentials/token"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"testing"
)

var testServer *grpc.Server
var conn *grpc.ClientConn
var connRead *grpc.ClientConn
var connNoCreds *grpc.ClientConn

// TestMain starts the testServer and connections to it.
// To allow testing with insecure connections (no TLS) run the tests with
// go test -tags devmode
// this is needed for the token auth.
func TestMain(m *testing.M) {
	testServer = grpc.NewServer()

	register(testServer)

	lis, err := net.Listen("tcp", ":"+os.Getenv("MTR_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Print("starting test server")
	go testServer.Serve(lis)

	conn, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithPerRPCCredentials(tk.New(os.Getenv("MTR_TOKEN_WRITE"))),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connRead, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithPerRPCCredentials(tk.New(os.Getenv("MTR_TOKEN_READ"))),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connNoCreds, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	code := m.Run()

	conn.Close()
	connNoCreds.Close()
	testServer.Stop()

	os.Exit(code)
}
