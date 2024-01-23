# CronJob single Yaml sample

```bash
$ kubectl create -f cronjob.yaml
cronjob.batch/hello created

$ kubectl get cronjob hello

NAME    SCHEDULE    SUSPEND   ACTIVE   LAST SCHEDULE   AGE
hello   * * * * *   False     0        54s             83s

$ kubectl get jobs --watch

NAME             COMPLETIONS   DURATION   AGE
hello-28432874   1/1           3s         66s
hello-28432875   1/1           2s         6s
hello-28432876   0/1                      0s
hello-28432876   0/1           0s         0s
hello-28432876   0/1           3s         3s
hello-28432876   1/1           3s         3s

$ kubectl get pods
hello-28432997-cx5hf         0/1     Completed   0              2m3s
hello-28432998-bxtz5         0/1     Completed   0              63s
hello-28432999-jmrt6         0/1     Completed   0              3s

$ kubectl logs hello-28432999-jmrt6
Tue Jan 23 03:19:00 UTC 2024
Hello from the minukube cluster

$ kubectl delete cronjob hello
cronjob.batch "hello" deleted
```