# Gobot + Wemos

### Project structure

#### Firmata

Arduino IDE

#### Infra

`docker-compose up`

#### Publisher

`go build . && ./publisher --influx_url=http://DOCKER_IP:8086 --nats_url=nats://DOCKER_IP:4222`

#### Wemos

`go build . && ./wemos --firmata_url=FIRMATA_IP --nats_url=nats://DOCKER_IP:4222`
