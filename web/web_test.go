package web

import (
	"chaos-slave/chaoslogger"
	"chaos-slave/proto"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
	"time"
)

var (
	logger = getLogger()
)

func TestHealthCheckGrpcCheckSuccess(t *testing.T) {
	done := make(chan struct{})
	errOnGrpcHandlerRun := make(chan error)
	port := "8080"

	withTestGrpcServer(port, done, errOnGrpcHandlerRun)

	clientConn, err := withTestGrpcClientConn(port)
	if err != nil {
		t.Fatalf("Can not create client connection")
	}

	client := proto.NewHealthClient(clientConn)
	resp := performHChReqOnCheck(client)

	done <- struct{}{}
	clientConn.Close()

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Nil(t, <-errOnGrpcHandlerRun)
	assert.Equal(t, proto.HealthCheckResponse_SERVING, resp.Status)
}

func TestHealthCheckGrpcCheckInvalidPort(t *testing.T) {
	done := make(chan struct{})
	errOnGrpcHandlerRun := make(chan error)
	port := "-1"

	withTestGrpcServer(port, done, errOnGrpcHandlerRun)

	clientConn, err := withTestGrpcClientConn(port)
	if err != nil {
		t.Fatalf("Can not create client connection")
	}

	client := proto.NewHealthClient(clientConn)
	resp := performHChReqOnCheck(client)

	done <- struct{}{}
	clientConn.Close()

	assert.Nil(t, resp)
	assert.EqualError(t, <-errOnGrpcHandlerRun, fmt.Sprintf("listen tcp: address %s: invalid port", port))
}

// === End to end testing ===
func TestStartServiceGrpcSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	done := make(chan struct{})
	errOnGrpcHandlerRun := make(chan error)
	port := "8080"

	withTestGrpcServer(port, done, errOnGrpcHandlerRun)

	clientConn, err := withTestGrpcClientConn(port)
	if err != nil {
		t.Fatalf("Can not create client connection")
	}

	client := proto.NewServiceClient(clientConn)
	resp := performStartServiceReq(client)

	done <- struct{}{}
	clientConn.Close()

	assert.Nil(t, err)
	assert.Nil(t, <-errOnGrpcHandlerRun)
	assert.Equal(t, proto.StatusResponse_SUCCESS, resp.Status)
	assert.NotNil(t, resp.Message)
}

// === End to end testing ===
func TestStopServiceGrpcSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	done := make(chan struct{})
	errOnGrpcHandlerRun := make(chan error)
	port := "8080"

	withTestGrpcServer(port, done, errOnGrpcHandlerRun)

	clientConn, err := withTestGrpcClientConn(port)
	if err != nil {
		t.Fatalf("Can not create client connection")
	}

	client := proto.NewServiceClient(clientConn)
	resp := performStopServiceReq(client)

	done <- struct{}{}
	clientConn.Close()

	assert.Nil(t, err)
	assert.Nil(t, <-errOnGrpcHandlerRun)
	assert.Equal(t, proto.StatusResponse_SUCCESS, resp.Status)
	assert.NotNil(t, resp.Message)
}


func withTestGrpcServer(port string, done chan struct{}, errOnGrpcHandlerRun chan error) {
	grpcHandler := NewGrpcHandler(port, logger)
	go func(errOnGrpcHandlerRun chan error) {
		if err := grpcHandler.Run(); err != nil {
			errOnGrpcHandlerRun <- err
		} else {
			errOnGrpcHandlerRun <- nil
		}
	}(errOnGrpcHandlerRun)
	time.Sleep(2 * time.Second)

	go func(done chan struct{}) {
		select {
		case <-done:
			grpcHandler.Stop()
		}
	}(done)
}

func withTestGrpcClientConn(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", port), grpc.WithInsecure())
	if err != nil {
		_ = level.Error(logger).Log("msg", "failed to dial server address", "err", err)
		return nil, err
	}

	return conn, nil
}

func performHChReqOnCheck(client proto.HealthClient) *proto.HealthCheckResponse {
	resp, err := client.Check(context.Background(), &proto.HealthCheckRequest{})
	if err != nil {
		_ = level.Error(logger).Log("msg", "Failed to call the Check method on the Health-check", "err", err)
		return nil
	}
	return resp
}

func performStartServiceReq(client proto.ServiceClient) *proto.StatusResponse {
	resp, err := client.Start(context.Background(), &proto.ServiceRequest{Name: "simple"})
	if err != nil {
		_ = level.Error(logger).Log("msg", "Failed to call the Start method on the Service", "err", err)
		return resp
	}
	return resp
}

func performStopServiceReq(client proto.ServiceClient) *proto.StatusResponse {
	resp, err := client.Stop(context.Background(), &proto.ServiceRequest{Name: "simple"})
	if err != nil {
		_ = level.Error(logger).Log("msg", "Failed to call the Start method on the Service", "err", err)
		return resp
	}
	return resp
}

func getLogger() log.Logger {
	allowLevel := &chaoslogger.AllowedLevel{}
	if err := allowLevel.Set("debug"); err != nil {
		fmt.Printf("%v", err)
	}
	return chaoslogger.New(allowLevel)
}
