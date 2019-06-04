/*
Copyright 2019 The Unity Scheduler Authors

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

package plugins

import (
	"github.com/golang/glog"
)

var plugins SchedulerPlugins

func init() {
	plugins = SchedulerPlugins{}
}

func RegisterSchedulerPlugin(plugin interface{}) {
	switch t := plugin.(type) {
	case PredicatesPlugin:
		glog.V(4).Info("register scheduler plugin, type: PredicatesPlugin")
		plugins.predicatesPlugin = t
	default:
		glog.V(4).Info("type incompatible, there is no scheduler plugin to register")
	}
}

func GetPredicatesPlugin() PredicatesPlugin {
	return plugins.predicatesPlugin
}

