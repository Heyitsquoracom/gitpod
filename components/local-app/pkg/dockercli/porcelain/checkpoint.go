// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by porcelain/generator.go; DO NOT EDIT.
package porcelain

import "os/exec"

type CheckpointCreateOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool

	// Leave the container running after checkpoint.
	LeaveRunning bool
}

// Create a checkpoint from a running container.
func CheckpointCreate(opts *CheckpointCreateOpts, container string, checkpoint string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "create"},
		[]string{container, checkpoint},
		opts,
		0,
	)
}

type CheckpointLsOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool
}

// List checkpoints for a container.
func CheckpointLs(opts *CheckpointLsOpts, container string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "ls"},
		[]string{container},
		opts,
		0,
	)
}

type CheckpointRmOpts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

	// Use a custom checkpoint storage directory.
	CheckpointDir string

	// Print usage.
	Help bool
}

// Remove a checkpoint.
func CheckpointRm(opts *CheckpointRmOpts, container string, checkpoint string) (string, error) {
	return runCtrCmd(
		[]string{"checkpoint", "rm"},
		[]string{container, checkpoint},
		opts,
		0,
	)
}
