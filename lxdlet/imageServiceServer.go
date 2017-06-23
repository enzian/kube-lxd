package lxdlet

import (
	"golang.org/x/net/context"

	runtime "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

type ImageServiceServer struct {
}

func (iss ImageServiceServer) ListImages(context.Context, *runtime.ListImagesRequest) (*runtime.ListImagesResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) ImageStatus(context.Context, *runtime.ImageStatusRequest) (*runtime.ImageStatusResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) PullImage(context.Context, *runtime.PullImageRequest) (*runtime.PullImageResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) RemoveImage(context.Context, *runtime.RemoveImageRequest) (*runtime.RemoveImageResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) ImageFsInfo(context.Context, *runtime.ImageFsInfoRequest) (*runtime.ImageFsInfoResponse, error) {
	return nil, nil
}
