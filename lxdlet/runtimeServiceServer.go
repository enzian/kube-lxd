package lxdlet

import (
	"golang.org/x/net/context"

	runtimeApi "github.com/kubernetes/kubernetes/pkg/kubelet/apis/cri/v1alpha1/runtime"
)

type RuntimeServiceServer struct {
}

func (rss RuntimeServiceServer) Version(context.Context, *runtimeApi.VersionRequest) (*runtimeApi.VersionResponse, error) {
	name := "lxd"
	version := "0.1.0"
	return &runtimeApi.VersionResponse{
		Version:           version, // kubelet/remote version, must be 0.1.0
		RuntimeName:       name,
		RuntimeVersion:    version,
		RuntimeApiVersion: version,
	}, nil
}

func (rss RuntimeServiceServer) RunPodSandbox(context.Context, *runtimeApi.RunPodSandboxRequest) (*runtimeApi.RunPodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StopPodSandbox(context.Context, *runtimeApi.StopPodSandboxRequest) (*runtimeApi.StopPodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) RemovePodSandbox(context.Context, *runtimeApi.RemovePodSandboxRequest) (*runtimeApi.RemovePodSandboxResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) PodSandboxStatus(context.Context, *runtimeApi.PodSandboxStatusRequest) (*runtimeApi.PodSandboxStatusResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListPodSandbox(context.Context, *runtimeApi.ListPodSandboxRequest) (*runtimeApi.ListPodSandboxResponse, error) {
	var pods = []*runtimeApi.PodSandbox{}

	var resp = &runtimeApi.ListPodSandboxResponse{
		Items: pods,
	}

	return resp, nil
}

func (rss RuntimeServiceServer) CreateContainer(context.Context, *runtimeApi.CreateContainerRequest) (*runtimeApi.CreateContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StartContainer(context.Context, *runtimeApi.StartContainerRequest) (*runtimeApi.StartContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) StopContainer(context.Context, *runtimeApi.StopContainerRequest) (*runtimeApi.StopContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) RemoveContainer(context.Context, *runtimeApi.RemoveContainerRequest) (*runtimeApi.RemoveContainerResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListContainers(context.Context, *runtimeApi.ListContainersRequest) (*runtimeApi.ListContainersResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ContainerStatus(context.Context, *runtimeApi.ContainerStatusRequest) (*runtimeApi.ContainerStatusResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ExecSync(context.Context, *runtimeApi.ExecSyncRequest) (*runtimeApi.ExecSyncResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Exec(context.Context, *runtimeApi.ExecRequest) (*runtimeApi.ExecResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Attach(context.Context, *runtimeApi.AttachRequest) (*runtimeApi.AttachResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) PortForward(context.Context, *runtimeApi.PortForwardRequest) (*runtimeApi.PortForwardResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ContainerStats(context.Context, *runtimeApi.ContainerStatsRequest) (*runtimeApi.ContainerStatsResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) ListContainerStats(context.Context, *runtimeApi.ListContainerStatsRequest) (*runtimeApi.ListContainerStatsResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) UpdateRuntimeConfig(context.Context, *runtimeApi.UpdateRuntimeConfigRequest) (*runtimeApi.UpdateRuntimeConfigResponse, error) {
	return nil, nil
}

func (rss RuntimeServiceServer) Status(context.Context, *runtimeApi.StatusRequest) (*runtimeApi.StatusResponse, error) {
	// Use vendored strings
	runtimeReadyConditionString := runtimeApi.RuntimeReady
	networkReadyConditionString := runtimeApi.NetworkReady

	resp := &runtimeApi.StatusResponse{
		Status: &runtimeApi.RuntimeStatus{
			Conditions: []*runtimeApi.RuntimeCondition{
				{
					Type:   runtimeReadyConditionString,
					Status: true,
				},
				{
					Type:   networkReadyConditionString,
					Status: true,
				},
			},
		},
	}

	return resp, nil
}
