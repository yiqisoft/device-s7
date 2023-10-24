# Device Service for Siemens S7 PLC

<!---
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-s7-go/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-s7-go/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/device-s7-go/branch/main/graph/badge.svg?token=IUywg34zfH)](https://codecov.io/gh/edgexfoundry/device-s7-go) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/device-s7-go)](https://goreportcard.com/report/github.com/edgexfoundry/device-s7-go) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-s7-go?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/device-s7-go/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-s7-go?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/device-s7-go)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/device-s7-go) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/device-s7-go)](https://github.com/edgexfoundry/device-s7-go/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/device-s7-go)](https://github.com/edgexfoundry/device-s7-go/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/device-s7-go-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/device-s7-go)](https://github.com/edgexfoundry/device-s7-go/commits))
-->

## Overview

S7 Micro Service - device service for connecting Siemens S7(S7-200, S7-300, S7-400, S7-1200, S7-1500) devices by `ISO-on-TCP` to EdgeX.

### Branch

This branch is for EdgeX v2.

## Features

- Single Read and Write
- Multiple Read and Write

## Prerequisites

- A Siemens S7 series device with network interface
- Enable ISO-on-TCP connection

### Install and Deploy

Clone and Build

```shell
git clone https://github.com/edgexfoundy-holding/device-s7.git
cd device-s7
make build
```

To start the device service:

```shell
export EDGEX_SECURITY_SECRET_STORE=false
make run
```

To rebuild after making changes to source:

```shell
make clean
make build
```

## Packaging

This component is packaged as Docker image.

For docker, please refer to the `Dockerfile`.

### Build Docker image

```shell
make docker
```

The docker image looks like:

```
edgexfoundry/device-s7    2.3.0    91125917a73e    12 days ago    31.3MB
```

### Docker compose file

Add to your docker-compose.yml.

```yaml
device-s7:
  container_name: edgex-device-s7
  depends_on:
    consul:
      condition: service_started
    data:
      condition: service_started
    metadata:
      condition: service_started
  environment:
    CLIENTS_CORE_COMMAND_HOST: edgex-core-command
    CLIENTS_CORE_DATA_HOST: edgex-core-data
    CLIENTS_CORE_METADATA_HOST: edgex-core-metadata
    CLIENTS_SUPPORT_NOTIFICATIONS_HOST: edgex-support-notifications
    CLIENTS_SUPPORT_SCHEDULER_HOST: edgex-support-scheduler
    DATABASES_PRIMARY_HOST: edgex-redis
    EDGEX_SECURITY_SECRET_STORE: 'false'
    MESSAGEQUEUE_HOST: edgex-redis
    REGISTRY_HOST: edgex-core-consul
    SERVICE_HOST: edgex-device-s7
  hostname: edgex-device-modbus
  image: edgexfoundry/device-s7:2.3.0
  networks:
    edgex-network: {}
  ports:
    - 127.0.0.1:59994:59994/tcp
  read_only: true
  restart: always
  security_opt:
    - no-new-privileges:true
  user: 2002:2001
```

## Usage

### Device Profile Sample

You should change all `valueType` and `NodeName` to your real `configuration`.

```yaml
name: S7-Device
manufacturer: YIQISOFT
description: Example of S7 Device
model: Siemens S7
labels: [ISO-on-TCP]
deviceResources:
  - description: PLC bool
    name: bool
    isHidden: false
    properties:
      valueType: Bool
      readWrite: RW
    attributes:
      NodeName: DB4.DBX0.0
  - description: PLC byte
    name: byte
    isHidden: false
    properties:
      valueType: Uint8
      readWrite: RW
    attributes:
      NodeName: DB4.DBB1
  - description: PLC word
    name: word
    isHidden: false
    properties:
      valueType: Int16
      readWrite: RW
    attributes:
      NodeName: DB4.DBW2
  - description: PLC dword
    name: dword
    isHidden: false
    properties:
      valueType: Int32
      readWrite: RW
    attributes:
      NodeName: DB4.DBD4
  - description: PLC int
    name: int
    isHidden: false
    properties:
      valueType: Int16
      readWrite: RW
    attributes:
      NodeName: DB4.DBW8
  - description: PLC dint
    name: dint
    isHidden: false
    properties:
      valueType: Int32
      readWrite: RW
    attributes:
      NodeName: DB4.DBW10
  - description: PLC real
    name: real
    isHidden: false
    properties:
      valueType: Float32
      readWrite: RW
    attributes:
      NodeName: DB4.DBD14
  - description: PLC heartbeat
    name: heartbeat
    isHidden: false
    properties:
      valueType: Int16
      readWrite: RW
    attributes:
      NodeName: DB1.DBW160
deviceCommands:
  - name: AllResource
    isHidden: false
    readWrite: RW
    resourceOperations:
      - deviceResource: bool
      - deviceResource: byte
      - deviceResource: word
      - deviceResource: dword
      - deviceResource: int
      - deviceResource: dint
      - deviceResource: real
      - deviceResource: heartbeat
```

### Device Sample

Change `Host`, `Port`, `Rack`, `Slot` and `Interval` to your real `Configuration`.

```
[[DeviceList]]
  Name = "S7-Device01"
  ProfileName = "S7-Device1"
  Description = "Example of S7 Device"
  Labels = [ "industrial" ]
  [DeviceList.Protocols]
    [DeviceList.Protocols.s7]
      Host = "192.168.0.1"
      Port = "102"
      Rack = "0"
      Slot = "2"
  [[DeviceList.AutoEvents]]
    Interval = "300ms"
    OnChange = false
    SourceName = "AllResource"
```

### Service status

#### Sevice Ping

```shell
curl http://localhost:59994/api/v2/ping
```

```json
{
  "apiVersion": "v2",
  "serviceName": "device-s7",
  "timestamp": "Wed Oct 18 08:57:16 UTC 2023"
}
```

#### Get version

```shell
curl http://localhost:59994/api/v2/version
```

```json
{
  "apiVersion": "v2",
  "sdk_version": "0.0.0",
  "serviceName": "device-s7",
  "version": "2.3.0"
}
```

### Execute Commands

#### All device

```shell
curl http://localhost:59882/api/v2/device/all
```

```
{
  "apiVersion" : "v2",
  "deviceCoreCommands" : [
    {
      "coreCommands" : [
        {
          "get" : true,
          "name" : "heartbeat",
          "parameters" : [
            {
              "resourceName" : "heartbeat",
              "valueType" : "Int16"
            }
          ],
          "path" : "/api/v2/device/name/plc/heartbeat",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "AllResource",
          "parameters" : [
            {
              "resourceName" : "bool",
              "valueType" : "Bool"
            },
            {
              "resourceName" : "byte",
              "valueType" : "Uint8"
            },
            {
              "resourceName" : "word",
              "valueType" : "Int16"
            },
            {
              "resourceName" : "dword",
              "valueType" : "Int32"
            },
            {
              "resourceName" : "int",
              "valueType" : "Int16"
            },
            {
              "resourceName" : "dint",
              "valueType" : "Int32"
            },
            {
              "resourceName" : "real",
              "valueType" : "Float32"
            },
            {
              "resourceName" : "heartbeat",
              "valueType" : "Int16"
            }
          ],
          "path" : "/api/v2/device/name/plc/AllResource",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "bool",
          "parameters" : [
            {
              "resourceName" : "bool",
              "valueType" : "Bool"
            }
          ],
          "path" : "/api/v2/device/name/plc/bool",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "byte",
          "parameters" : [
            {
              "resourceName" : "byte",
              "valueType" : "Uint8"
            }
          ],
          "path" : "/api/v2/device/name/plc/byte",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "word",
          "parameters" : [
            {
              "resourceName" : "word",
              "valueType" : "Int16"
            }
          ],
          "path" : "/api/v2/device/name/plc/word",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "dword",
          "parameters" : [
            {
              "resourceName" : "dword",
              "valueType" : "Int32"
            }
          ],
          "path" : "/api/v2/device/name/plc/dword",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "int",
          "parameters" : [
            {
              "resourceName" : "int",
              "valueType" : "Int16"
            }
          ],
          "path" : "/api/v2/device/name/plc/int",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "dint",
          "parameters" : [
            {
              "resourceName" : "dint",
              "valueType" : "Int32"
            }
          ],
          "path" : "/api/v2/device/name/plc/dint",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        },
        {
          "get" : true,
          "name" : "real",
          "parameters" : [
            {
              "resourceName" : "real",
              "valueType" : "Float32"
            }
          ],
          "path" : "/api/v2/device/name/plc/real",
          "set" : true,
          "url" : "http://edgex-core-command:59882"
        }
      ],
      "deviceName" : "plc",
      "profileName" : "S7-Device"
    }
  ],
  "statusCode" : 200,
  "totalCount" : 1
}
```

#### Set command

```shell
curl http://localhost:59882/api/v2/device/name/plc/heartbeat \
-X PUT \
-H "Content-Type:application/json" \
-d '{"heartbeat": "1"}'
```

```json
{
  "apiVersion": "v2",
  "statusCode": 200
}
```

#### Get command

```shell
curl http://localhost:59882/api/v2/device/name/plc/heartbeat
```

```json
{
  "apiVersion": "v2",
  "event": {
    "apiVersion": "v2",
    "deviceName": "plc",
    "id": "ee208561-4f9b-48ea-beb7-f54aaec579d7",
    "origin": 1697618905997597628,
    "profileName": "S7-Device",
    "readings": [
      {
        "deviceName": "plc",
        "id": "3505d09a-450b-4fc4-991a-1b3737a4fa8d",
        "origin": 1697618905997591878,
        "profileName": "S7-Device",
        "resourceName": "heartbeat",
        "value": "1",
        "valueType": "Int16"
      }
    ],
    "sourceName": "heartbeat"
  },
  "statusCode": 200
}
```

### For EdgeX v3

Checkout to branch v3.

## Reference

- [Gos7](https://github.com/robinson/gos7)

## License

Apache-2.0
