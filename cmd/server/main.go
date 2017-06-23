package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/enzian/kube-lxd/lxdlet"
	"github.com/golang/glog"
	runtimeapi "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
	"github.com/kubernetes/kubernetes/pkg/util/logs"
	"google.golang.org/grpc"
)

const defaultUnixSock = "/var/run/lxdlet.sock"

func main() {

	logs.InitLogs()
	defer logs.FlushLogs()

	glog.Warning("This lxd CRI server implementation is for development use only; we recommend using the copy of this code included in the kubelet")

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)

	socketPath := defaultUnixSock
	defer os.Remove(socketPath)

	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		glog.Fatalf("Error listening on sock %q: %v ", socketPath, err)
	}
	defer sock.Close()

	grpcServer := grpc.NewServer()

	rktruntime, err := lxdlet.New()
	if err != nil {
		glog.Fatalf("could not create lxdlet: %v", err)
	}

	runtimeapi.RegisterImageServiceServer(grpcServer, rktruntime)
	runtimeapi.RegisterRuntimeServiceServer(grpcServer, rktruntime)

	glog.Infof("Starting to serve on %q", socketPath)
	go grpcServer.Serve(sock)

	<-exitCh

	glog.Infof("lxdlet service exiting...")
}
