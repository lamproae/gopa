/*
Copyright 2016 Medcl (m AT medcl.net)

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

package modules

import (
	log "github.com/cihub/seelog"
	. "github.com/medcl/gopa/core/env"
	apiModule "github.com/medcl/gopa/modules/api"
	crawlerModule "github.com/medcl/gopa/modules/crawler"
	storageModule "github.com/medcl/gopa/modules/storage"
	"github.com/medcl/gopa/modules/checker"
)

type Modules struct {
	env     *Env
	modules map[string]ModuleInterface
}

func New(env *Env) *Modules {
	modules := Modules{}
	modules.env = env
	return &modules
}

func (this *Modules) Start() {

	//start modules
	storageModule.Start(this.env)
	apiModule.Start(this.env)
	crawlerModule.Start(this.env)
	//parserModule.Start(this.env)
	url_checker.Start(this.env)
}

func (this *Modules) Stop() {
	//parserModule.Stop()
	url_checker.Stop()
	crawlerModule.Stop()
	apiModule.Stop()
	storageModule.Stop()
	this.env.RuntimeConfig.Storage.Close()
	log.Info("all modules stopeed")
}
