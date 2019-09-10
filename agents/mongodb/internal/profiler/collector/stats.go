// pmm-agent
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

type stats struct {
	In                     int64  `name:"in"`
	Out                    int64  `name:"out"`
	IteratorCreated        string `name:"iterator-created"`
	IteratorCounter        int64  `name:"iterator-counter"`
	IteratorRestartCounter int64  `name:"iterator-restart-counter"`
	IteratorErrLast        string `name:"iterator-err-last"`
	IteratorErrCounter     int64  `name:"iterator-err-counter"`
}
