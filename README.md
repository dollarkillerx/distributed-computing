# distributed-computing
distributed-computing

## Google三剑客:
- GFS: Google File System
- BigTable
- MapReduce
 
### GFS    
架构Master Work
Master存储index
Work存储对应文件

热点数据热平衡
![README](./README/c1.png)

![WRITE](./README/c2.png)

### Bigtable
传统
![传统](./README/c3.png)
拆分
![j1](./README/c4.png)
再拆分
![j1](./README/c5.png)
写数据    < buf size 写入内存  打log防止数据丢失   > bu size 写文件
![j1](./README/c6.png)
读数据
![j1](./README/c7.png)
与GFS结合 
![j1](./README/c8.png)
将逻辑视图转换为物理存储   
![j1](./README/c9.png)
架构
![j1](./README/c10.png)

### MapReduce
![j1](./README/c11.png)
wordCount
![j1](./README/c12.png)
build inverted index
![j1](./README/c13.png)
架构
![j1](./README/c14.png)


### CAP 
一致性（Consistency）、可用性（Availability）、分区容错性（Partition tolerance）。CAP 原则指的是，这三个要素最多只能同时实现两点，不可能三者兼顾。

### Raft 分布式一致性协议
图示 http://thesecretlivesofdata.com/raft/
![j1](./README/p1.png)
我们来用Raft实现一个强一致性的分布式MAP




