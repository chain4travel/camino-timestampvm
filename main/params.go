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
	"flag"

	"github.com/chain4travel/camino-timestampvm/timestampvm"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	versionKey = "version"
)

func buildFlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet(timestampvm.Name, flag.ContinueOnError)

	fs.Bool(versionKey, false, "If true, prints Version and quit")

	return fs
}

// getViper returns the viper environment for the plugin binary
func getViper() (*viper.Viper, error) {
	v := viper.New()

	fs := buildFlagSet()
	pflag.CommandLine.AddGoFlagSet(fs)
	pflag.Parse()
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}

	return v, nil
}

func PrintVersion() (bool, error) {
	v, err := getViper()
	if err != nil {
		return false, err
	}

	return v.GetBool(versionKey), nil
}
