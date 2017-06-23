package lxdlet

import (
	"golang.org/x/net/context"

	runtime "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

type RuntimeServiceServer struct {
}

func (rss RuntimeServiceServer) Version(context.Context, *runtime.VersionRequest) (*runtime.VersionResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) RunPodSandbox(context.Context, *runtime.RunPodSandboxRequest) (*runtime.RunPodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StopPodSandbox(context.Context, *runtime.StopPodSandboxRequest) (*runtime.StopPodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) RemovePodSandbox(context.Context, *runtime.RemovePodSandboxRequest) (*runtime.RemovePodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) PodSandboxStatus(context.Context, *runtime.PodSandboxStatusRequest) (*runtime.PodSandboxStatusResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListPodSandbox(context.Context, *runtime.ListPodSandboxRequest) (*runtime.ListPodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) CreateContainer(context.Context, *runtime.CreateContainerRequest) (*runtime.CreateContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StartContainer(context.Context, *runtime.StartContainerRequest) (*runtime.StartContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StopContainer(context.Context, *runtime.StopContainerRequest) (*runtime.StopContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) RemoveContainer(context.Context, *runtime.RemoveContainerRequest) (*runtime.RemoveContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListContainers(context.Context, *runtime.ListContainersRequest) (*runtime.ListContainersResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ContainerStatus(context.Context, *runtime.ContainerStatusRequest) (*runtime.ContainerStatusResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ExecSync(context.Context, *runtime.ExecSyncRequest) (*runtime.ExecSyncResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Exec(context.Context, *runtime.ExecRequest) (*runtime.ExecResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Attach(context.Context, *runtime.AttachRequest) (*runtime.AttachResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) PortForward(context.Context, *runtime.PortForwardRequest) (*runtime.PortForwardResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ContainerStats(context.Context, *runtime.ContainerStatsRequest) (*runtime.ContainerStatsResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListContainerStats(context.Context, *runtime.ListContainerStatsRequest) (*runtime.ListContainerStatsResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) UpdateRuntimeConfig(context.Context, *runtime.UpdateRuntimeConfigRequest) (*runtime.UpdateRuntimeConfigResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Status(context.Context, *runtime.StatusRequest) (*runtime.StatusResponse, error) {
	return nil, nil
}
