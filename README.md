# go- distributed job schedular system


<!-- ABOUT THE PROJECT -->

## About The Project

This project is a distributed job schedular system with the Master-Worker pattern and goroutine in Golang.
Here is the demo:
* https://user-images.githubusercontent.com/85295969/182254985-1d8a3042-4169-4b45-b04d-e2238578cb80.mp4

![demogif](https://user-images.githubusercontent.com/85295969/182524245-a6ef4648-2126-48b4-9522-0620737ea14f.gif)
![ezgif com-gif-maker](https://user-images.githubusercontent.com/85295969/182540078-70b87680-736c-4474-9dc8-945c500fc78e.gif)


### Features
![crontab](https://user-images.githubusercontent.com/85295969/182517389-8973484b-ffde-472c-acee-1ad10e4c3ea0.jpg)
 
* Built a front end as web management console to manage and schedule jobs.
* Leveraged etcd to store jobs as key-value pairs, monitor the changes of ongoing jobs, synchronize jobs across all.
workers with Raft, and implement a distributed lock for the job hold by a worker.
* Created REST APIs in Golang for CRUD operations, killing jobs, and job log queries from MongoDB.
* Implemented the Cron daemon with crontab commands to run jobs at the given time and frequencies.


### Built With

Built services on Centos7 Linux:

1. install etcd on Centos7
2. run etcd: 
   ```sh
   $ nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' -- advertise-client-urls 'http://0.0.0.0:2379' &
   ```
4. install mongoDB on Centos7
5. create a database : mkdir data
6. run mongoDB: 
   ```sh
   $ nohup bin/mongod --dbpath=./data --bind_ip=0.0.0.0 &
   ```
   ```sh
   $ bin/mongo
   ```

Install package for Go:
* cron packadge
  ```sh
  $go get github.com/gorhill/cronexpr
  ```
  

* etcd package:
  ```sh
  $ go get GitHub.com/coreos/etcd/clientv3
  ```
  Build connection with your Linux service ip

* install mongoDB package:


<!-- GETTING STARTED -->
## How to run
```sh
$ go build github.com/owenliang/crontab/master/main
```
```sh
$ go build github.com/owenliang/crontab/worker/main
```

