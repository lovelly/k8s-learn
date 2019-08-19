package main

import (
	"k8s-learn/cni/pkg"
	"runtime"

	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	bv "github.com/containernetworking/plugins/pkg/utils/buildversion"
)

func init() {
	// 这确保main只在主线程（线程组组长）上运行。
	// 因为命名空间操作（unshare，setns）是针对单个线程完成的。
	// 所以必须确保goroutine不会从OS线程跳转到线程 。
	runtime.LockOSThread()
}

func main() {
	skel.PluginMain(pkg.CmdAdd, pkg.CmdCheck, pkg.CmdDel, version.All, bv.BuildString("bridge"))
}
