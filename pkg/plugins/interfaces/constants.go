// Copyright 2024 Deepgram CLI contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package interfaces

// CmdGroup is a group of CLI commands.
type CmdGroup string

const (
	// StartersCmdGroup are commands associated with starter apps.
	StartersCmdGroup CmdGroup = "Starters"

	// TestCmdGroup is the test command group.
	TestCmdGroup CmdGroup = "Test"
)
