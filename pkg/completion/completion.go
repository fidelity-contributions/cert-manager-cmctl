/*
Copyright 2021 The cert-manager Authors.

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

package completion

import (
	"context"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/cert-manager/cmctl/v2/pkg/build"
)

func NewCmdCompletion(setupCtx context.Context, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "completion",
		Short: "Generate completion scripts for the cert-manager CLI",
		Long:  "Generate completion for the cert-manager CLI so arguments and flags can be suggested and auto-completed",
	}

	if build.IsKubectlPlugin(setupCtx) {
		cmds.AddCommand(newCmdCompletionKubectl(setupCtx, ioStreams))
	} else {
		cmds.AddCommand(newCmdCompletionBash(setupCtx, ioStreams))
		cmds.AddCommand(newCmdCompletionZSH(setupCtx, ioStreams))
		cmds.AddCommand(newCmdCompletionFish(setupCtx, ioStreams))
		cmds.AddCommand(newCmdCompletionPowerShell(setupCtx, ioStreams))
	}

	return cmds
}
