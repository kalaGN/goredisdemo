/**
feature
read config from cofig.ini
cron set redis heartbeat

 */
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-ini/ini"
	"time"
)

func main() {
	TimerDemo := time.NewTimer(time.Duration(5) * time.Second)
	//循环监听定时器
	for {
		select { case <-TimerDemo.C:
			redishearbeat()
			//超时后重置定时器
			TimerDemo.Reset(time.Duration(5) * time.Second) }
	}
}

func redishearbeat() {

	config := getconfig("config.ini")
	c, err := redis.Dial("tcp", config.host+":"+config.port)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	tpass := config.password
	//解密密码
	//auth认证
	if _, err := c.Do("AUTH", tpass); err != nil {

	}
	c.Do("select", config.db)
	defer c.Close()
	//update the redis ins key value to new time
	timeUnix := time.Now().Unix()
	res, err := redis.String(c.Do("set", config.insid, timeUnix))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Println("set success to ins1:result:" + res)
	}
}

//redis config return
type config struct {
	host     string
	port     string
	password string
	db       string
	timeout  string
	insid    string
}

func getconfig(inifile string) config {
	//获取配置
	cfg, err := ini.Load(inifile)
	if err != nil {
		fmt.Printf("fail to read file: $v", err)
	}
	return config{
		host:     cfg.Section("production").Key("shardredis.host").String(),
		port:     cfg.Section("production").Key("shardredis.port").String(),
		password: cfg.Section("production").Key("shardredis.password").String(),
		db:       cfg.Section("production").Key("shardredis.db").String(),
		timeout:  cfg.Section("production").Key("shardredis.timeout").String(),
		insid:    cfg.Section("production").Key("system.instanceid").String(),
	}

}
