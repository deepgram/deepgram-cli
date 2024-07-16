// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/deepgram-devs/deepgram-cli/cmd"

	_ "github.com/deepgram-devs/deepgram-cli/cmd/analyze"
	_ "github.com/deepgram-devs/deepgram-cli/cmd/listen"
	_ "github.com/deepgram-devs/deepgram-cli/cmd/manage"
	_ "github.com/deepgram-devs/deepgram-cli/cmd/selfhosted"
	_ "github.com/deepgram-devs/deepgram-cli/cmd/speak"
)

func main() {
	cmd.Execute()
}
