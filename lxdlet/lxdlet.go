package lxdlet

import runtimeapi "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"

func New() (ContainerAndImageService, error) {
	return combinedRuntimes{
		RuntimeServiceServer: RuntimeServiceServer{},
		ImageServiceServer:   ImageServiceServer{},
	}, nil
}

type combinedRuntimes struct {
	runtimeapi.RuntimeServiceServer
	runtimeapi.ImageServiceServer
}

type ContainerAndImageService interface {
	runtimeapi.RuntimeServiceServer
	runtimeapi.ImageServiceServer
}
