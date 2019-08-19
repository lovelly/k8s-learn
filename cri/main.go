package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	maxMsgSize = 65536
)

// Serivice implements runtime.ImageService and runtime.RuntimeService.
type MyService struct {
}

func NewMyService() *MyService {
	return &MyService{}
}

func (s *MyService) Version(context.Context, *runtime.VersionRequest) (*runtime.VersionResponse, error) {
	panic("implement me")
}

func (s *MyService) RunPodSandbox(context.Context, *runtime.RunPodSandboxRequest) (*runtime.RunPodSandboxResponse, error) {
	panic("implement me")
}

func (s *MyService) StopPodSandbox(context.Context, *runtime.StopPodSandboxRequest) (*runtime.StopPodSandboxResponse, error) {
	panic("implement me")
}

func (s *MyService) RemovePodSandbox(context.Context, *runtime.RemovePodSandboxRequest) (*runtime.RemovePodSandboxResponse, error) {
	panic("implement me")
}

func (s *MyService) PodSandboxStatus(context.Context, *runtime.PodSandboxStatusRequest) (*runtime.PodSandboxStatusResponse, error) {
	panic("implement me")
}

func (s *MyService) ListPodSandbox(context.Context, *runtime.ListPodSandboxRequest) (*runtime.ListPodSandboxResponse, error) {
	panic("implement me")
}

func (s *MyService) CreateContainer(context.Context, *runtime.CreateContainerRequest) (*runtime.CreateContainerResponse, error) {
	panic("implement me")
}

func (s *MyService) StartContainer(context.Context, *runtime.StartContainerRequest) (*runtime.StartContainerResponse, error) {
	panic("implement me")
}

func (s *MyService) StopContainer(context.Context, *runtime.StopContainerRequest) (*runtime.StopContainerResponse, error) {
	panic("implement me")
}

func (s *MyService) RemoveContainer(context.Context, *runtime.RemoveContainerRequest) (*runtime.RemoveContainerResponse, error) {
	panic("implement me")
}

func (s *MyService) ListContainers(context.Context, *runtime.ListContainersRequest) (*runtime.ListContainersResponse, error) {
	panic("implement me")
}

func (s *MyService) ContainerStatus(context.Context, *runtime.ContainerStatusRequest) (*runtime.ContainerStatusResponse, error) {
	panic("implement me")
}

func (s *MyService) UpdateContainerResources(context.Context, *runtime.UpdateContainerResourcesRequest) (*runtime.UpdateContainerResourcesResponse, error) {
	panic("implement me")
}

func (s *MyService) ReopenContainerLog(context.Context, *runtime.ReopenContainerLogRequest) (*runtime.ReopenContainerLogResponse, error) {
	panic("implement me")
}

func (s *MyService) ExecSync(context.Context, *runtime.ExecSyncRequest) (*runtime.ExecSyncResponse, error) {
	panic("implement me")
}

func (s *MyService) Exec(context.Context, *runtime.ExecRequest) (*runtime.ExecResponse, error) {
	panic("implement me")
}

func (s *MyService) Attach(context.Context, *runtime.AttachRequest) (*runtime.AttachResponse, error) {
	panic("implement me")
}

func (s *MyService) PortForward(context.Context, *runtime.PortForwardRequest) (*runtime.PortForwardResponse, error) {
	panic("implement me")
}

func (s *MyService) ContainerStats(context.Context, *runtime.ContainerStatsRequest) (*runtime.ContainerStatsResponse, error) {
	panic("implement me")
}

func (s *MyService) ListContainerStats(context.Context, *runtime.ListContainerStatsRequest) (*runtime.ListContainerStatsResponse, error) {
	panic("implement me")
}

func (s *MyService) UpdateRuntimeConfig(context.Context, *runtime.UpdateRuntimeConfigRequest) (*runtime.UpdateRuntimeConfigResponse, error) {
	panic("implement me")
}

func (s *MyService) Status(context.Context, *runtime.StatusRequest) (*runtime.StatusResponse, error) {
	panic("implement me")
}

type MyImgService struct {
}

func (img *MyImgService) ListImages(context.Context, *runtime.ListImagesRequest) (*runtime.ListImagesResponse, error) {
	panic("implement me")
}

func (img *MyImgService) ImageStatus(context.Context, *runtime.ImageStatusRequest) (*runtime.ImageStatusResponse, error) {
	panic("implement me")
}

func (img *MyImgService) PullImage(context.Context, *runtime.PullImageRequest) (*runtime.PullImageResponse, error) {
	panic("implement me")
}

func (img *MyImgService) RemoveImage(context.Context, *runtime.RemoveImageRequest) (*runtime.RemoveImageResponse, error) {
	panic("implement me")
}

func (img *MyImgService) ImageFsInfo(context.Context, *runtime.ImageFsInfoRequest) (*runtime.ImageFsInfoResponse, error) {
	panic("implement me")
}

func main() {
	service := &MyService{}
	imgService := &MyImgService{}
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize))
	runtime.RegisterRuntimeServiceServer(s, service)
	runtime.RegisterImageServiceServer(s, imgService)
	lis, err := net.Listen("unix", "/var/run/runtime.sock")
	if err != nil {
		log.Println(err)
		return
	}
	go s.Serve(lis)

	// Other codes
}
