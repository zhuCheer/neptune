package connect

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"fmt"
	"github.com/zhuCheer/pool"
	"github.com/zhuCheer/neptune/helper"
	"errors"
	"log"
	"sync"
)


var GlobalRedisHandler = &redisHandler{PoolMap:make(map[string]pool.Pool)}

// 连接池全局存储变量
type redisHandler struct {
	mu          sync.Mutex
	PoolMap 	map[string]pool.Pool
}

// 连接池注册节点
type RegistCnf struct {
	Address string
	Auth    string
	InitialCap int
	MaxCap		int
	IdleTimeout time.Duration
	DBNum		int
}

// 连接信息节点
type ConnItem struct{
	PoolHashId string
	Conn redis.Conn
}

// 获取一个连接对象
func (rs *redisHandler) GetConn(poolHashId string) (item ConnItem, err error){
	poolItem, ok := rs.PoolMap[poolHashId]
	if ok == false{
		log.Println("[error] redis pool is not found")
		return item, errors.New("redis pool is not found")
	}
	poolConn, err := poolItem.Get()
	if err != nil{
		log.Println("[error] get conn have an error in redis pool")
		return item, errors.New("get conn have an error in redis pool")
	}
	conn := poolConn.(redis.Conn)
	item = ConnItem{
		poolHashId,
		conn,
	}

	return item,nil
}

// 将一个连接放回连接池
func (rs *redisHandler) PutConn(item ConnItem) (err error){
	poolHashId := item.PoolHashId
	poolItem, ok := rs.PoolMap[poolHashId]
	if ok == false{
		log.Println("[error] redis pool is not found")
		return errors.New("redis pool is not found")
	}
	err = poolItem.Put(item.Conn)
	if err != nil{
		log.Println("[error] put conn have an error in redis pool")
		return errors.New("put conn have an error in redis pool")
	}
	return nil
}

// 释放一个连接池
func (rs *redisHandler) ReleasePool(poolHashId string) (err error){
	poolItem, ok := rs.PoolMap[poolHashId]

	if ok == false{
		log.Println("[warning] redis pool is not found, do not anything")
		return
	}
	poolItem.Release()
	delete(rs.PoolMap, poolHashId)

	return
}


// 注册一个池
func Regist(cnf *RegistCnf) (poolHashId string, err error){
	poolHashId = "QI:"+helper.String2md5(cnf.Address)
	GlobalRedisHandler.mu.Lock()
	defer GlobalRedisHandler.mu.Unlock()

	if _, ok := GlobalRedisHandler.PoolMap[poolHashId]; ok {
		log.Println("pool has exists:"+poolHashId)
		return poolHashId, nil
	}
	fmt.Println(cnf)

	//factory 创建连接的方法
	factory := func() (interface{}, error) {
		fmt.Println("open link" + poolHashId)
		c, err := redis.Dial("tcp", cnf.Address)
		if cnf.Auth != ""{
			if _, err := c.Do("AUTH", cnf.Auth); err != nil {
				c.Close()
				return nil, fmt.Errorf("[factory]redis auth error: %s", err)
			}
		}
		if cnf.DBNum > 0 {
			if _, err := c.Do("SELECT", cnf.DBNum); err != nil {
				c.Close()
				return nil, fmt.Errorf("[factory]redis select db num error: %s", err)
			}
			fmt.Println("select db num:",cnf.DBNum)
		}

		return c, err
	}

	//close 关闭连接的方法
	close := func(v interface{}) error {
		fmt.Println("close link " + poolHashId)
		return v.(redis.Conn).Close()
	}

	poolCnf := &pool.Config{
		InitialCap: cnf.InitialCap,
		MaxCap:     cnf.MaxCap,
		Factory:    factory,
		Close:      close,
		//Ping:       ping,
		IdleTimeout: cnf.IdleTimeout * time.Second,
	}
	pool, err := getPool(poolCnf)
	if err !=nil {
		log.Println("[error]regist pool error: ", err)
		return poolHashId, fmt.Errorf("regist pool error: %s", err)
	}

	GlobalRedisHandler.PoolMap[poolHashId] = pool
	return poolHashId,nil
}


func getPool(poolConfig *pool.Config)(pool.Pool, error){
	fmt.Println("create a new pool")

	//factory 创建连接的方法
	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		log.Println("[error]create pool have an error", err)
		return nil, fmt.Errorf("create pool have an error: %s", err)
	}
	return p, nil
}
