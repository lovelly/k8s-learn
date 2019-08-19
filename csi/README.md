开启参数

kube-apiserver:
--allow-privileged=true \
--feature-gates=BlockVolume=true,CSIBlockVolume=true,CSIPersistentVolume=true,MountPropagation=true,VolumeSnapshotDataSource=true, \
KubeletPluginsWatcher=true,CSINodeInfo=true,CSIDriverRegistry=true

kubelet:
--allow-privileged=true --feature-gates=BlockVolume=true,CSIBlockVolume=true,CSIPersistentVolume=true,MountPropagation=true,VolumeSnapshotDataSource=true, KubeletPluginsWatcher=true,CSINodeInfo=true,CSIDriverRegistry=true

controller-manager:
--feature-gates=BlockVolume=true,CSIBlockVolume=true