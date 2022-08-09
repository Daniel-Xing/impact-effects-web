# Impact Effect Web

> English | [中文](./README_zh.md)

![Docker Build Status badge](https://img.shields.io/badge/docker%20build-passing-brightgreen)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

Impact-effect-web project is an asteroid/comet Impact simulation system based on [Impact-Effect](https://github.com/acse-dx121/impact-effects) under the guidance of [Professor Gareth Collins](http://www.imperial.ac.uk/people/g.collins) and [Dr Thomas M Guide to build Davison](https://www.imperial.ac.uk/people/thomas.davison). According to the parameters given by the user, the Web program will give the possible corresponding impact consequences.

The project is based on the front and back end separation architecture, the front end is built based on vue.js, the back end is built based on Golang&Python, and docker-compose is supported for rapid deployment.

## :crystal_ball: Visuals

**Annotation Platform**
![fore-end-show](doc/img/fore-end-show.png)

**Architecture**
![webArch](doc/img/webArch.png)

## 🍞 Features

- It supports user-defined input data and computes relevant impact results.
- Docker Compose installation is supported
  - Higher service stability: The container can be restarted immediately after exiting. After the host crashes, it can quickly move to another machine.​
  - Good isolation: Similar to Python's virtual environment, it is also running in a container in a completely isolated environment. Users do not need to run an additional configuration for the application.​
  - Easy to expand:  Once the container is constructed successfully, it is very easy to run multiple applications and form a cluster.​

## 🍕 Requirements

### Back-end

- python > 3.7
- Golang >= 1.14
- Gin v1
- Gorm v1
- Redis
- GRPC
- Docker(optional)

### Fore-end

- node.js
- npm/cnpm
- vue.js/webPack/etc

## 🚍 Installation

### 🚀 Quick Start (Docker)

All backend services are containerized, and the project provides a cluster.yml file in the root directory. One-click startup with docker-compose is supported.

```shell
docker-compose -f cluster.yml up // create the cluster
docker-compose -f cluster.yml down //distory the cluster
```

Since the front-end service needs to access the background service, the fixed IP address is specified when the background is containerized, so generally speaking, it can run without modification.

```bash
cd front-web
npm install
npm run dev
```

### manul install

**克隆项目** First clone the project

```bash
# clone the project
git clone https://github.com/acse-dx121/impact-effects-web.git
```

**Function Service** Go to the Function service directory and build the project. Service will listen on port 50051. Make sure the firewall port is developed, otherwise you will not be able to access the service.

```bash
cd function-service
# create the virtural env
conda env create -f environment.yml
# activate env
conda activate functions-service
# run the service
python service.py


```

**Back-end Service** Go to the back-web directory and run the backend application. The server will listen on port 50052. Make sure the firewall port is developed, otherwise you will not be able to access the service. In addition, if you manually changed the listening port of the function service and Redis service, you need to enter the corresponding file to make changes.

```bash
# make sure you already install golang
cd back-web && go mode tidy 
# run the service
go run main.go
```

**Fore-end Service** Go to the front-web directory and run the frontend application. The server will listen on port 9999. Also, make sure the firewall is set up correctly, and the port number for the backend service is correct.

```bash
cd front-web
npm install
npm run dev
```

## 🚩  License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2013-present, Yuxi (Evan) You


