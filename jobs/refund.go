package jobs

import (
	"fmt"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	infra "github.com/masterzcw/resk.wen"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/go-utils"
	"resk.com/core/envelopes"
)

type RefundExpiredJobStarter struct {
	infra.BaseStarter
	ticker *time.Ticker
	mutex  *redsync.Mutex
}

func (r *RefundExpiredJobStarter) Init(ctx infra.StarterContext) {
	d := ctx.Props().GetDurationDefault("jobs.refund.interval", time.Minute)
	r.ticker = time.NewTicker(d)
	maxIdle := ctx.Props().GetIntDefault("redis.maxIdle", 2)                   // 最大空闲连接数
	maxActive := ctx.Props().GetIntDefault("redis.maxActive", 5)               // 最大活动连接数
	timeout := ctx.Props().GetDurationDefault("redis.timeout", 20*time.Second) // 超时时间
	addr := ctx.Props().GetDefault("redis.addr", "127.0.0.1:6379")             // 连接地址
	pools := make([]redsync.Pool, 0)                                           // 连接池切片
	pool := &redis.Pool{
		MaxIdle:     maxIdle,   // 最大空闲连接数
		MaxActive:   maxActive, // 最大活动连接数
		IdleTimeout: timeout,   // 超时时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", addr)
		}, // 创建连接
	} // 配置属性
	pools = append(pools, pool) // 把链接加入连接池
	rsync := redsync.New(pools) // 把连接池对象绑进去,new一个分布式互斥锁
	ip, err := utils.GetExternalIP()
	if err != nil {
		ip = "127.0.0.1"
	}
	// 加锁
	r.mutex = rsync.NewMutex("lock:RefundExpired",
		redsync.SetExpiry(50*time.Second), // 50秒过期
		redsync.SetRetryDelay(3),
		redsync.SetGenValueFunc(func() (s string, e error) {
			now := time.Now()
			log.Infof("节点%s正在执行过期红包的退款任务", ip)
			return fmt.Sprintf("%d:%s", now.Unix(), ip), nil
		}),
	)
}

func (r *RefundExpiredJobStarter) Start(ctx infra.StarterContext) {
	go func() {
		for {
			c := <-r.ticker.C

			err := r.mutex.Lock() // 拿锁: 如果拿到锁就进行红包退款, 没有拿到就不执行任何操作,只记录一条日志出来
			if err == nil {
				log.Debug("过期红包退款开始...", c)
				//红包过期退款的业务逻辑代码
				domain := envelopes.ExpiredEnvelopeDomain{}
				domain.Expired()
			} else {
				log.Info("已经有节点在运行该任务了")
			}
			r.mutex.Unlock()

		}
	}()
}

func (r *RefundExpiredJobStarter) Stop(ctx infra.StarterContext) {
	r.ticker.Stop()
}
