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
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/types"
)

// tcp-check send-binary <hexstring> [comment <msg>]
type CheckSendBinary struct {
	HexString    string
	CheckComment string
	Comment      string
}

func (c *CheckSendBinary) Parse(parts []string, parserType types.ParserType, comment string) error {
	if comment != "" {
		c.Comment = comment
	}
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}
	c.HexString = parts[2]
	for i := 3; i < len(parts); i++ {
		el := parts[i]
		if el == "comment" {
			checkParsePair(parts, &i, &c.CheckComment)
		}
	}
	return nil
}

func (c *CheckSendBinary) String() string {
	sb := &strings.Builder{}
	sb.WriteString("send-binary")
	sb.WriteString(" ")
	sb.WriteString(c.HexString)
	checkWritePair(sb, "comment", c.CheckComment)
	return sb.String()
}

func (c *CheckSendBinary) GetComment() string {
	return c.Comment
}
