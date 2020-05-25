# 传统方案 - crontab

- 配置任务时，需要 ssh 登录脚本服务器进行操作

- 服务器宕机，任务将终止调度，需要人工迁移

- 排查问题低效，无法方便的查看状态与错误输出

# 分布式任务调度

- 可视化 web 后台，方便进行任务管理
- 分布式结构、集群化调度，不存在单点故障
- 追踪任务执行状态，采集任务输出，可视化 log 查看

# golang shell

# etcd 的用处

- k8s

  - 服务发现
  - 集群状态存储
  - 配置同步

- CLOUD FOUNDRY
  - 集群状态存储
  - 配置同步
  - 分布式锁

# etcd 与 Raft 的关系

- Raft 是强一致的集群日志同步算法
- etcd 是一个分布式 KV 存储
- etcd 利用 raft 算法在集群中同步 key-value

      nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379' &
