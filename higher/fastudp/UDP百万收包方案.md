# UDP百万收包方案.md

参考文档：https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/114957102

### 1.端口重用：SO_REUSEPORT

### 2.使用recvmmsg代替recvmsg

### 3.网卡多队列绑定CPU核心优化

### 4.系统调用分离

### 5.内核读缓冲区扩大