# goredisdemo
## a go demo to read redis config from config.ini and set a redis key periodically 

# use
~~~
git clone https://github.com/kalaGN/goredisdemo.git
~~~
## then edit the config.ini to you  reids config
~~~$
shardredis.host='127.0.0.1'
shardredis.port='6379'
shardredis.password='123456'
shardredis.db='1'
shardredis.timeout='0'
~~~
~~~
go build main.go
./main
~~~
## if it show 
~~~
set success to ins1:result:OK
set success to ins1:result:OK
~~~
## It works 
## enjoy it
