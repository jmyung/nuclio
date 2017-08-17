/*
Copyright 2017 The Nuclio Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runtime

import (
	"github.com/nuclio/nuclio-sdk"
)

type Runtime interface {
	ProcessEvent(event nuclio.Event) (interface{}, error)
}

type AbstractRuntime struct {
	Logger  nuclio.Logger
	Context *nuclio.Context
}

func NewAbstractRuntime(logger nuclio.Logger, configuration *Configuration) *AbstractRuntime {
	return &AbstractRuntime{
		Logger:  logger,
		Context: newContext(logger, configuration),
	}
}