// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

// Package exec defines tasks for running commands.
//
// These are the supported tasks:
//
//	// Run executes the given shell command.
//	Run: {
//		$id: *"tool/exec.Run" | "exec" // exec for backwards compatibility
//
//		// cmd is the command to run.
//		cmd: string | [string, ...string]
//
//		// dir specifies the working directory of the command.
//		// The default is the current working directory.
//		dir?: string
//
//		// env defines the environment variables to use for this system.
//		// If the value is a list, the entries mus be of the form key=value,
//		// where the last value takes precendence in the case of multiple
//		// occurrances of the same key.
//		env: [string]: string | [...=~"="]
//
//		// stdout captures the output from stdout if it is of type bytes or string.
//		// The default value of null indicates it is redirected to the stdout of the
//		// current process.
//		stdout: *null | string | bytes
//
//		// stderr is like stdout, but for errors.
//		stderr: *null | string | bytes
//
//		// stdin specifies the input for the process. If stdin is null, the stdin
//		// of the current process is redirected to this command (the default).
//		// If it is of typ bytes or string, that input will be used instead.
//		stdin: *null | string | bytes
//
//		// success is set to true when the process terminates with a zero exit
//		// code or false otherwise. The user can explicitly specify the value
//		// force a fatal error if the desired success code is not reached.
//		success: bool
//	}
package exec

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/pkg/internal"
)

func init() {
	internal.Register("tool/exec", pkg)
}

var _ = adt.TopKind // in case the adt package isn't used

var pkg = &internal.Package{
	Native: []*internal.Builtin{},
	CUE: `{
	Run: {
		$id:  *"tool/exec.Run" | "exec"
		cmd:  string | [string, ...string]
		dir?: string
		env: {
			[string]: string | [...=~"="]
		}
		stdout:  *null | string | bytes
		stderr:  *null | string | bytes
		stdin:   *null | string | bytes
		success: bool
	}
}`,
}
