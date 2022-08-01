# go- distributed job schedular system


<!-- ABOUT THE PROJECT -->
## About The Project


This project is a distributed job schedular system with the Master-Worker pattern and goroutine in Golang.

### Features
• Built a front end as web management console to manage and schedule jobs.
• Leveraged etcd to store jobs as key-value pairs, monitor the changes of ongoing jobs, synchronize jobs across all.
workers with Raft, and implement a distributed lock for the job hold by a worker.
• Created REST APIs in Golang for CRUD operations, killing jobs, and job log queries from MongoDB.
• Implemented the Cron daemon with crontab commands to run jobs at the given time and frequencies.


### Built With

Built services on Centos7 Linux:

1. install etcd on Centos7
2. run etcd: 
   $ nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' -- advertise-client-urls 'http://0.0.0.0:2379' &
4. install mongoDB on Centos7
5. create a database : mkdir data
6. run mongoDB: 
   $ nohup bin/mongod --dbpath=./data --bind_ip=0.0.0.0 &
   $ bin/mongo

Install package for Go:
1.cron packadge
  $go get github.com/gorhill/cronexpr

2.install etcd package:
  $ go get GitHub.com/coreos/etcd/clientv3
  Build connection with your Linux service ip

3.install mongoDB package:


<!-- GETTING STARTED -->
## How to run

$ go build github.com/owenliang/crontab/master/main
$ go build github.com/owenliang/crontab/worker/main


