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

package actions

import (
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/types"
)

type ExpectProxy struct {
	Cond     string
	CondTest string
	Comment  string
}

func (f *ExpectProxy) Parse(parts []string, parserType types.ParserType, comment string) error {
	if f.Comment != "" {
		f.Comment = comment
	}
	_, condition := common.SplitRequest(parts[2:])
	if len(condition) > 1 {
		f.Cond = condition[0]
		f.CondTest = strings.Join(condition[1:], " ")
	}
	return nil
}

func (f *ExpectProxy) String() string {
	var result strings.Builder
	result.WriteString("expect-proxy layer4")
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *ExpectProxy) GetComment() string {
	return f.Comment
}
