apiserver 启动参数  --etcd-servers=http://localhost:2379 --secure-port=443 --tls-cert-file=./apiserver/config/certificates/apiserver.crt --tls-private-key-file=./apiserver/config/certificates/apiserver.key --delegated-auth=true --kubeconfig=./config --authentication-kubeconfig=./config --authorization-kubeconfig=./config --audit-webhook-config-file=./config


controller 启动参数  --kubeconfig=./config