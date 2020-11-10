# proxy-protocol-debugger

This simple application listens for HTTP connections on TCP8080. After establishing a connection, the application logs and returns address of the connecting client.

When the application receives a request that has [PROXY protocol](https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt) header, it will use the client information from that header.

## Example

### Request

```
$ curl 7c7b4f52-us-south.lb.appdomain.cloud
169.51.88.35:57812
```

### Logs

```
$ kubectl logs proxy-protocol-debugger-7cf657675b-f5b8h | jq .
{
  "level": "info",
  "ts": 1605037999.625721,
  "caller": "proxy-protocol-debugger/main.go:40",
  "msg": "successfully retrieved remote address",
  "remote": "169.51.88.35:57812"
}
{
  "level": "info",
  "ts": 1605038052.4959483,
  "caller": "proxy-protocol-debugger/main.go:40",
  "msg": "successfully retrieved remote address",
  "remote": "87.99.85.140:60361"
}
{
  "level": "info",
  "ts": 1605038278.7522736,
  "caller": "proxy-protocol-debugger/main.go:40",
  "msg": "successfully retrieved remote address",
  "remote": "172.17.3.8:55058"
}
```

## Deploy

You can use the following manifests to deploy and expose the application on a Kubernetes cluster running in IBM Cloud.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: proxy-protocol-debugger
spec:
  replicas: 1
  selector:
    matchLabels:
      app: proxy-protocol-debugger
  template:
    metadata:
      labels:
        app: proxy-protocol-debugger
    spec:
      containers:
        - name: proxy-protocol-debugger
          image: attiss/proxy-protocol-debugger:latest
          imagePullPolicy: Always
          ports:
            - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: proxy-protocol-debugger
  annotations:
    service.kubernetes.io/ibm-load-balancer-cloud-provider-enable-features: "proxy-protocol"
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
      name: http
  selector:
    app: proxy-protocol-debugger
```

## Credits

This application uses the [pires/go-proxyproto](https://github.com/pires/go-proxyproto) project.
