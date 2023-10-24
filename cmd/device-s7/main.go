// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
// Copyright (C) 2023 YIQISOFT
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"

	"github.com/edgexfoundry/device-s7-go"
	"github.com/edgexfoundry/device-s7-go/driver"
)

const (
	serviceName string = "device-s7"
)

func main() {
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device.Version, sd)
}
