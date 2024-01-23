# DNS Synthetic Job

## Build and Run Locally

```bash
$ docker build -t hgsgtk/dnsynthetic .
$ docker tag hgsgtk/dnsynthetic hgsgtk/dnsynthetic:latest
$ docker push hgsgtk/dnsynthetic:latest
```

## Remote repository

https://hub.docker.com/r/hgsgtk/dnsynthetic

## Execute on Kubernetes

```bash
$ $ kubectl create -f dnsynthetic.yaml --save-config
cronjob.batch/dnsynthetic created

$ kubectl get dnsynthetic dnsynthetic

$ kubectl get pods

$ kubectl logs <pod-name>

$ kubectl delete cronjob hello
```

- When you apply a change to the kubernetes config

```bash
$ kubectl apply -f dnsynthetic.yaml                                                                   1 ↵
cronjob.batch/dnsynthetic configured
```

## CoreDNS

```bash
$ kubectl get --namespace=kube-system deployments
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
coredns   0/1     0            0           3m21s
```

### If fail

```bash
$ kubectl logs dnsynthetic-28433155-2cqwk
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:36555->10.96.0.10:53: i/o timeout
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:42122->10.96.0.10:53: i/o timeout
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:35494->10.96.0.10:53: i/o timeout
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:49327->10.96.0.10:53: i/o timeout
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:60912->10.96.0.10:53: i/o timeout
lookup golang.org on 10.96.0.10:53: read udp 10.244.0.240:54600->10.96.0.10:53: i/o timeout
```

## References
- Kubernetes image pull policy 
    - https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
    - IfNotPresent, Always, Never

