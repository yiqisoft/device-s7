// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2023 YIQISOFT
//
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/robinson/gos7"
)

type ServiceConfig struct {
	S7Info S7Info
}

// UpdateFromRaw updates the service's full configuration from raw data received from
// the Service Provider.
func (sw *ServiceConfig) UpdateFromRaw(rawConfig interface{}) bool {
	configuration, ok := rawConfig.(*ServiceConfig)
	if !ok {
		return false //errors.New("unable to cast raw config to type 'ServiceConfig'")
	}

	*sw = *configuration

	return true
}

type S7Info struct {
	// PLC connection info
	Host string
	Port int
	Rack int
	Slot int

	// DB address, start, size
	DbAddress    int
	StartAddress int
	ReadSize     int

	ConnEstablishingRetry int
	ConnRetryWaitTime     int

	Writable WritableInfo
}

type S7Client struct {
	DeviceName string
	Client     gos7.Client
}

// Validate ensures your custom configuration has proper values.
func (info *S7Info) Validate() errors.EdgeX {
	if info.Writable.ResponseFetchInterval == 0 {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "S7Info.Writable.ResponseFetchInterval configuration setting can not be blank", nil)
	}
	return nil
}

type WritableInfo struct {
	// ResponseFetchInterval specifies the retry interval(milliseconds) to fetch the command response from the MQTT broker
	ResponseFetchInterval int
}

func fetchCommandOpts(protocols map[string]models.ProtocolProperties) (string, errors.EdgeX) {
	properties, ok := protocols[Protocol]
	if !ok {
		return "", errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("'%s' protocol properties is not defined", Protocol), nil)
	}
	commandTopic, ok := properties[CommandTopic]
	if !ok {
		return "", errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("'%s' not found in the '%s' protocol properties", CommandTopic, Protocol), nil)
	}
	return commandTopic, nil
}
