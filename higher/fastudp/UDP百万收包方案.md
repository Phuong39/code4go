# UDP百万收包方案.md

参考文档：https://zhuanlan.zhihu.com/p/357902432

### 1.端口重用：SO_REUSEPORT

### 2.使用recvmmsg代替recvmsg

### 3.网卡多队列绑定CPU核心优化

### 4.系统调用分离

### 5.内核读缓冲区扩大

### 6.并发监听，利用多核优势