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

## References
- Kubernetes image pull policy 
    - https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy
    - IfNotPresent, Always, Never

