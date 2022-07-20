// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************

// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"os"

	log "github.com/inconshreveable/log15"

	"github.com/chain4travel/camino-timestampvm/timestampvm"
	"github.com/chain4travel/caminogo/vms/rpcchainvm"
)

func main() {
	version, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	// Print VM ID and exit
	if version {
		fmt.Printf("%s@%s\n", timestampvm.Name, timestampvm.Version)
		os.Exit(0)
	}

	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.TerminalFormat())))

	rpcchainvm.Serve(&timestampvm.VM{})
}
