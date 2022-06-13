# controller-runtime
Controller-runtime是一个用于开发Kubernetes Controller的库，包含了各种Controller常用的模块，兼顾了灵活性和模块化。

一开始做Kubernetes Controller开发时，是学习`simple-controller`使用`client-go`进行开发，中间会有很多与业务无关的重复工作。后来社区推出了`kubebuilder`，它可以方便的渲染出Controller的整个框架，让开发者只用专注Controller本身的业务逻辑，特别是在开发CRD时，极为方便，而kubebuilder渲染出的框架使用的则是`controller-runtime`。

Controller-runtime中为Controller的开发提供了各种功能模块，主要包括：
- Client：用于读写Kubernetes资源
- Cache：本地缓存，可供Client直接读取资源。
- Manager：可以管理协调多个Controller，提供Controller共用的依赖。
- Controller：“组装”多个模块（例如Source、Queue、Reconciler），实现Kubernetes Controller的通用逻辑：
- Reconciler：状态同步的逻辑所在，是开发者需要实现的主要接口，供Controller调用。Reconciler的重点在于“状态同步”，由于Reconciler传入的参数是资源的Namespace和Name，而非event，Reconciler并非用于“处理事件”，而是根据指定资源的状态，来同步“预期集群状态”与“当前集群状态”。
- Webhook：用于开发webhook server，实现Kubernetes Admission Webhooks机制。
- Source：source of event，Controller从中获取event。
- EventHandler：顾名思义，event的处理方法，决定了一个event是否需要入队列、如何入队列。
- Predicate：相当于event的过滤器。

## Test
我们的示例代码的逻辑：监听指定namespace下一个名为nginx-clean的pod，当该pod被delete后，执行一些清理task。当然，我们的目的是演示controller的reconcile功能，所以cleanup task是伪代码。

### 编译
```shell
go build -o cleanup
```

### 运行
```shell
$ ./cleanup 
I0613 01:37:02.330795   18697 main.go:37] CleanUpController Manager init success.
I0613 01:37:02.330920   18697 main.go:47] CleanUpController Reconciler init success.
I0613 01:37:02.432154   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432519   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432550   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432621   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432673   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432700   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432720   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432739   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:29.305273   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:30.539335   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:31.553537   18697 controller.go:35] CleanUpController Reconcile is running...
```

### 删除nginx-clean pod
```shell
$ kubectl get po 
NAME                              READY   STATUS    RESTARTS   AGE
details-v1-79f774bdb9-rwlcf       2/2     Running   0          10d
nginx-clean                       1/1     Running   0          12h
nginx-pod-1654079527724902515     1/1     Running   0          11d
productpage-v1-6b746f74dc-qc7v7   2/2     Running   0          10d
ratings-v1-b6994bb9-h9srv         2/2     Running   0          10d
reviews-v1-545db77b95-rlvsd       2/2     Running   0          10d
reviews-v2-7bf8c9648f-fnkxc       2/2     Running   0          10d
reviews-v3-84779c7bbc-jkkjc       2/2     Running   0          10d

$ kubectl delete po nginx-clean 
pod "nginx-clean" deleted
```

cleanup的日志：

```shell
[root@biz-master-48 ~]# ./cleanup 
I0613 01:37:02.330795   18697 main.go:37] CleanUpController Manager init success.
I0613 01:37:02.330920   18697 main.go:47] CleanUpController Reconciler init success.
I0613 01:37:02.432154   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432519   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432550   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432621   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432673   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432700   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432720   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:02.432739   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:29.305273   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:30.539335   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:31.553537   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:31.555745   18697 controller.go:35] CleanUpController Reconcile is running... 
I0613 01:37:31.555773   18697 controller.go:39] nginx-clean pod was deleted. 
I0613 01:37:31.555778   18697 controller.go:41] Let's start clean up task ...
```

正常打印出了 `nginx-clean pod was deleted. Let's start clean up task ... ` , 说明 reconcile 正常起作用了。

