kubelet 起订是指定 容器运行时监听的socket

kubelet --container-runtime=remote --container-runtime-endpoint=unix:///var/run/runtime.sock --image-service-endpoint=unix:///var/run/runtime.sock