# Women who Go Workshop: IoT with GoBot

### Set up your environment
In order to be able to work with the micro-controller Wemos & and the GoLang framework GoBot we need to setup your machine first. Please have a look at our (preparational readme)[https://github.com/WATTx/wwg-gobot/blob/master/README.prep.md] and follow the instructions. 



#### Firmata
The `firmata` directory contains the sketch that we will flash the micro-controller with in order to be able to send instructions to the micro-controller.

We are using the Firmata library and its protocol to communicate with Wemos. (More information about Firmata in general)[https://www.arduino.cc/en/Reference/Firmata]

More information about how to use it you will find in the (prep readme)[https://github.com/WATTx/wwg-gobot/blob/master/README.prep.md].


#### Infra
The `infra` directory contains a docker-compose file that will help you to run the containers for influxDB (timeseries database), Chronograf (monitoring frontend) & NATS (messaging system).

```
cd infra
docker-compose up
```

(In order to use docker-compose up, you need to have docker & docker-compose installed.
(prep readme)[https://github.com/WATTx/wwg-gobot/blob/master/README.prep.md])


#### Publisher
The `publisher` will subscribe to the different topics in NATS (e.g. humidty or motion) and will write it to influxDB.

You need to build and start it like:
```
cd publisher
go build . && ./publisher --influx_url=http://DOCKER_IP:8086 --nats_url=nats://DOCKER_IP:4222
```

(Please note: We are using the default ports here. If you specify other ports in docker-compose file or the scripts, you need to adjust them here as well.)


#### Wemos
The `wemos` directory contains the source code for talking with Wemos via GoBot and the functions that it should execute, e.g. reading values from sensors that are attached to the Wemos or letting LED blink. 

You need to build and start it like:
```
cd wemos
go build . && ./wemos --firmata_url=FIRMATA_IP:3030 --nats_url=nats://DOCKER_IP:4222
```

(Please note: We are using the default ports here. If you specify other ports in docker-compose file or the scripts, you need to adjust them here as well.)
