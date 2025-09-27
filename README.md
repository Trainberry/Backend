# Trainberry - Server

This project contains the central server in charge of controlling our trains.

## Prerequisites

You must have a Bluetooth antenna.

## Run it

Simply use the Golang CLI: `go run ./cmd`

There is a few environment variables that you can set to custom the behaviour:

* `TRAINBERRY_LOG_LEVEL` (Default: `INFO`): Set log level
* `TRAINBERRY_SANITY_RUN_INTERVAL` (Default: `5s`): Time between two executions of sanity run
* `TRAINBERRY_BLE_DETECTION_RUN_INTERVAL` (Default: `2s): Time between two network scan to find new trains

## APIs

### REST

There is only one REST endpoint on this server:

`GET` `/state` will return the current state of the server and associated chips. The format is the following:

```json
[
  {
    "name": "Trainberry::BB22001",
    "error_count": 0,
    "speed": 0,
    "light": true
  }
]
```

### WebSocket

99% of the trains are controlled through a WebSocket. I made this decision because communicating via WS is way much
quicker (and latency matters a lot!). Websocket is on `/ws` endpoint.

#### Server -> Client

The following information are sent by the server to the client(s):
* `added_device`: A new device have been added to the server
* `deleted_device`: A device have been removed after multiple healthcheck failure
* `fail_counter_increment`: A device missed a healthcheck
* `fail_counter_reset`: A device which was failing is back online
* `light_update`: The light setting (on/off) have been updated on this device
* `speed_update`: The speed setting have been updated on this device

This information are communicated through WS in a JSON format. The format is as it follows:

```json
{
  "operation":"deleted_device",
  "device": {
    "name":"Trainberry::BB22001",
    "error_count":0,
    "speed":0,
    "light":false
  }
}
```

The server will always send the full object, **except on deletion event, where only the name is sent**.

#### Client -> Server

The clients can contact the server to make updates on devices. The operations are:
* `set_light`: set light status for the desired device.
  * Payload is: `{"operation":"set_light","device":{"name":"Trainberry::BB22001","light":true}}`
* `set_speed`: set speed for the desired device.
  * Payload is: `{"operation":"set_speed","device":{"name":"Trainberry::BB22001","speed":80}}`
* `emerg_stop`: stop all trains simultaneously.
  * Payload is: `{"operation":"emerg_stop"}`

# License

This project is licensed under [GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/).
