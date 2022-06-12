# sharedInformer

## Informer 设计思路

### Informer 设计中的关键点

为了让Client-go 更快地返回List/Get请求的结果、减少对 Kubernetes API的直接调用，Informer 被设计实现为一个依赖Kubernetes List/Watch API、可监听事件并触发回调函数的二级缓存工具包。

### 更快地返回 List/Get 请求，减少对 Kubernetes API 的直接调用

使用Informer实例的Lister()方法，List/Get Kubernetes 中的 Object时，Informer不会去请求Kubernetes API，而是直接查找缓存在本地内存中的数据(这份数据由Informer自己维护)。通过这种方式，Informer既可以更快地返回结果，又能减少对 Kubernetes API 的直接调用。


### 依赖 Kubernetes List/Watch API

Informer 只会调用Kubernetes List 和 Watch两种类型的 API。Informer在初始化的时，先调用Kubernetes List API 获得某种 resource的全部Object，缓存在内存中; 然后，调用 Watch API 去watch这种resource，去维护这份缓存; 最后，Informer就不再调用Kubernetes的任何 API。

用List/Watch去维护缓存、保持一致性是非常典型的做法，但令人费解的是，Informer 只在初始化时调用一次List API，之后完全依赖 Watch API去维护缓存，没有任何resync机制。

笔者在阅读Informer代码时候，对这种做法十分不解。按照多数人思路，通过 resync机制，重新List一遍 resource下的所有Object，可以更好的保证 Informer 缓存和 Kubernetes 中数据的一致性。

咨询过Google 内部 Kubernetes开发人员之后，得到的回复是:

在 Informer 设计之初，确实存在一个relist无法去执 resync操作， 但后来被取消了。原因是现有的这种 List/Watch 机制，完全能够保证永远不会漏掉任何事件，因此完全没有必要再添加relist方法去resync informer的缓存。这种做法也说明了Kubernetes完全信任etcd。


### 可监听事件并触发回调函数

Informer通过Kubernetes Watch API监听某种 resource下的所有事件。而且，Informer可以添加自定义的回调函数，这个回调函数实例(即 ResourceEventHandler 实例)只需实现 OnAdd(obj interface{}) OnUpdate(oldObj, newObj interface{}) 和OnDelete(obj interface{}) 三个方法，这三个方法分别对应informer监听到创建、更新和删除这三种事件类型。

在Controller的设计实现中，会经常用到 informer的这个功能。

### 关键逻辑介绍

1.Informer 在初始化时，Reflector 会先 List API 获得所有的 Pod

2.Reflect 拿到全部 Pod 后，会将全部 Pod 放到 Store 中

3.如果有人调用 Lister 的 List/Get 方法获取 Pod， 那么 Lister 会直接从 Store 中拿数据

4.Informer 初始化完成之后，Reflector 开始 Watch Pod，监听 Pod 相关 的所有事件;如果此时 pod_1 被删除，那么 Reflector 会监听到这个事件

5.Reflector 将 pod_1 被删除 的这个事件发送到 DeltaFIFO

6.DeltaFIFO 首先会将这个事件存储在自己的数据结构中(实际上是一个 queue)，然后会直接操作 Store 中的数据，删除 Store 中的 pod_1

7.DeltaFIFO 再 Pop 这个事件到 Controller 中

8.Controller 收到这个事件，会触发 Processor 的回调函数

9.LocalStore 会周期性地把所有的 Pod 信息重新放到 DeltaFIFO 中

## reference

[Client-go 源码分析之 SharedInformer 及实战](https://xie.infoq.cn/article/c1bac14389d533d5e84844f4e)

[理解 K8S 的设计精髓之 List-Watch机制和Informer模块](https://cloud.tencent.com/developer/article/1533221?from=article.detail.1717404)



