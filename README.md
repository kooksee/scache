# scache
暂时用于用户内容的地址的分享,这是一个去中心化的全同步的http应用

## 数据同步方案
1. 保留每个节点的数据同步状态
2. 然后定时或者实时检查数据
3. 每个客户端只跟自己的节点进行交互
4. 每个节点都有自己的ID,每个节点的数据都带有版本信息
5. 并且版本信息都是连续的,所以要同步所有的数据，就需要检查每个节点的版本信息
6. 然后，把缺失的信息发送给对应的节点
7. 如果一个客户端并没有在当前节点查找大数据，那么，节点会异步的去查找一下，然后再把数据缓存下来，等待下次查找