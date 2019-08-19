#### 生成证书
kubectl delete csr admission-webhook-example-svc.default
cfssl print-defaults csr > server-csr.json
cfssl genkey server-csr.json | cfssljson -bare server
export CSR_BUNDLE=$(cat server.csr |base64 | tr -d '\n')
sed -i "s|{CSR_BUNDLE}|${CSR_BUNDLE}|g" certReq.yaml
kubectl certificate approve  admission-webhook-example-svc.default
kubectl get csr admission-webhook-example-svc.default -o jsonpath='{.status.certificate}' | base64 -d  > server.pem

kubectl create secret generic  admission-webhook-example-svc.default \
        --from-file=key.pem=server-key.pem \
        --from-file=cert.pem=server.pem \
        --dry-run -o yaml  |
        kubectl apply -f -
        
export CA_BUNDLE=$(cat /etc/kubernetes/pki/ca.crt | base64 | tr -d "\n")
sed -i "s|{CA_BUNDLE}|${CA_BUNDLE}|g"

需要先创建 webhook的配置 在创建 webhook controller   