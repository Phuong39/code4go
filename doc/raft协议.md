# Raft协议

[参考文档](https://zhuanlan.zhihu.com/p/91288179)

#### Raft中的角色：
1. 领导人（leader）
2. 追随者（followers）
3. 候选人（candidate）

> 正常模式下只有leader和followers，当leader宕机之后followers会成为candidate，选举新的leader

#### 如何保证数据一致性

> 通过复制状态机模型，leader收到写入请求，都会复制一份操作日志给followers同步

#### 选举流程

1. 集群启动场景
> 