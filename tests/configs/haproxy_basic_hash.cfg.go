/*
Copyright 2019 HAProxy Technologies

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
package configs

const configBasicHash = `# _md5hash=4cadde7159beea87ba17ba043201d3ad
# _version=1
# HAProxy Technologies
# https://www.haproxy.com/

global
  master-worker

defaults
  log global

frontend http
  mode http
  bind 0.0.0.0:80 name bind_1
  bind :::80 v4v6 name bind_2
  default_backend default_backend

backend default_backend
  mode http
  http-request deny deny_status 400 # deny
`
