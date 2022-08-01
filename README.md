# go-crontab


<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot]
![Screen Shot 2022-08-01 at 1 50 33 PM](https://user-images.githubusercontent.com/85295969/182211218-09724c1f-92d5-46b0-b6b0-f5cbf063a6e6.png)

TooTasty is a full-stack website application, which can search, rate and review restaurants.
https://tootasty.herokuapp.com/

TooTasty implements authentication and authorization. Usera can only review restaurants and view their profile after logging in. What's more, 
there are three diffrent roles: customer, owner, and admin, who have diffrent authority and view diffrente pages.

### Built With

Built services on Linux:

1. install etcd
2. run etcd: nohup ./etcd --listen-client-urls 'http://0.0.0.0:2379' -- advertise-client-urls 'http://0.0.0.0:2379' &

Front-end Framework:

React, Redux, JavaScript, HTML, CSS, Bootstrap

Back-end Framework:

Node.js, Express, MongoDB

<!-- GETTING STARTED -->
## Getting Started

To implement this application:

1. install Go on your machine
2. install npm files
3. command "npm run dev" to run this whole application
4. command "npm run server" to run the backend
5. command "npm run client" to run the frontend

cron packadge
$go get github.com/gorhill/cronexpr
