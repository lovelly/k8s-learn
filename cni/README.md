创建配置文件给 kubelet加载

cat > /etc/cni/net.d/10-mymet.conf <<EOF
{
	"name": "mynet",
	"type": "bridge",
	"bridge": "mynet0",
	"isDefaultGateway": true,
	"forceAddress": false,
	"ipMasq": true,
	"hairpinMode": true,
	"ipam": {
		"type": "host-local",
		"subnet": "10.10.0.0/16"
	}
}
EOF

确保kubelet 开启了参数 --network-plugin=cni
--network-plugin=cni
--cni-conf-dir=/etc/cni/net.d 默认 
--cni-bin-dir=/opt/cni/bin   默认

把插件放到 --cni-bin=dir 指定的/opt/cni/bin目录下面 