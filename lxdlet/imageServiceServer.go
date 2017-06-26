package lxdlet

import (
	"golang.org/x/net/context"

	runtimeApi "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

type ImageServiceServer struct {
}

func (iss ImageServiceServer) ListImages(context.Context, *runtimeApi.ListImagesRequest) (*runtimeApi.ListImagesResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) ImageStatus(context.Context, *runtimeApi.ImageStatusRequest) (*runtimeApi.ImageStatusResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) PullImage(context.Context, *runtimeApi.PullImageRequest) (*runtimeApi.PullImageResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) RemoveImage(context.Context, *runtimeApi.RemoveImageRequest) (*runtimeApi.RemoveImageResponse, error) {
	return nil, nil
}

func (iss ImageServiceServer) ImageFsInfo(context.Context, *runtimeApi.ImageFsInfoRequest) (*runtimeApi.ImageFsInfoResponse, error) {
	return nil, nil
}
