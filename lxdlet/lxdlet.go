package lxdlet

import runtimeApi "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"

func New() (ContainerAndImageService, error) {
	return combinedRuntimes{
		RuntimeServiceServer: RuntimeServiceServer{},
		ImageServiceServer:   ImageServiceServer{},
	}, nil
}

type combinedRuntimes struct {
	runtimeApi.RuntimeServiceServer
	runtimeApi.ImageServiceServer
}

type ContainerAndImageService interface {
	runtimeApi.RuntimeServiceServer
	runtimeApi.ImageServiceServer
}
