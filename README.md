# go- distributed job schedular system


<!-- ABOUT THE PROJECT -->
## About The Project


This project is a distributed job schedular system using Golang. 

### Features
Web management console to manage and schedule jobs



### Built With

Built services on Linux:

1. install etcd
2. run etcd: nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' -- advertise-client-urls 'http://0.0.0.0:2379' &

Front-end Framework:

React, Redux, JavaScript, HTML, CSS, Bootstrap

Back-end Framework:

Node.js, Express, MongoDB

<!-- GETTING STARTED -->
## How to run

To implement this application:

1. install Go on your machine
2. install npm files
3. command "npm run dev" to run this whole application
4. command "npm run server" to run the backend
5. command "npm run client" to run the frontend

cron packadge
$go get github.com/gorhill/cronexpr
