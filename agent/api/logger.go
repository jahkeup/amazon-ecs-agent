// Copyright 2014-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package api

import (
	"github.com/aws/aws-sdk-go/aws"
)

// debugLogger is satisfied by most loggers capable of emitting
// leveled logs.
type debugLogger interface {
	Debug(...interface{})
}

// NewDebugLoggerConfig returns configuration suitable for handling
// SDK request debug messages on LogDebugWithRequestRetries and
// LogDebugWithRequestErrors.
func NewDebugLoggerConfig(logger debugLogger) *aws.Config {
	debug := aws.LoggerFunc(logger.Debug)
	return aws.NewConfig().
		WithLogger(&debug).
		WithLogLevel(aws.LogDebugWithRequestRetries | aws.LogDebugWithRequestErrors)
}
