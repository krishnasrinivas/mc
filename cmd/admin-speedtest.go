// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"context"
	"fmt"

	humanize "github.com/dustin/go-humanize"
	"github.com/minio/cli"
	"github.com/minio/mc/pkg/probe"
)

var adminSpeedtestFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "duration",
		Usage: "Duration in seconds for Uploads/Downloads",
		Value: 10,
	},
	cli.IntFlag{
		Name:  "size",
		Usage: "Object size",
		Value: 64 * humanize.MiByte,
	},
	cli.IntFlag{
		Name:  "threads",
		Usage: "number of threads per server",
		Value: 32,
	},
}

var adminSpeedtestCmd = cli.Command{
	Name:            "speedtest",
	Usage:           "Server side speed test",
	Action:          mainAdminSpeedtest,
	OnUsageError:    onUsageError,
	Before:          setGlobalsFromContext,
	Flags:           append(adminSpeedtestFlags, globalFlags...),
	HideHelpCommand: true,
	CustomHelpTemplate: `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} [FLAGS] TARGET

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}
EXAMPLES:
  1. Speedtest with default values:
     {{.Prompt}} {{.HelpName}} speedtest
  2. Speedtest for 10 seconds with object size of 128MB with 32 threads per server:
     {{.Prompt}} {{.HelpName}} speedtest --duration 20 --size 128000000 --threads 32

`,
}

func mainAdminSpeedtest(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		cli.ShowCommandHelpAndExit(ctx, "speedtest", 1) // last argument is exit code
	}

	// Get the alias parameter from cli
	args := ctx.Args()
	aliasedURL := args.Get(0)

	client, perr := newAdminClient(aliasedURL)
	if perr != nil {
		fatalIf(perr.Trace(aliasedURL), "Unable to initialize admin client.")
		return nil
	}

	ctxt, cancel := context.WithCancel(globalContext)
	defer cancel()

	size := ctx.Int("size")
	threads := ctx.Int("threads")
	durationSecs := ctx.Int("duration")

	results, err := client.Speedtest(ctxt, size, threads, durationSecs)
	if err != nil {
		fatalIf(probe.NewError(err), "speedtest error")
		return nil
	}
	uploads := uint64(0)
	downloads := uint64(0)
	for _, result := range results {
		uploads += result.Uploads
		downloads += result.Downloads
		// JSON mode should print stats for individual hosts:
		// fmt.Printf("Host: %s Upload: %s/s %d objs/s Download: %s/s %d objs/s\n", result.Endpoint, humanize.Bytes(result.Uploads*uint64(size)/uint64(durationSecs)), result.Uploads/uint64(durationSecs), humanize.Bytes(result.Downloads*uint64(size)/uint64(durationSecs)), result.Downloads/uint64(durationSecs))
	}
	fmt.Printf("PUT: %s/s %d objs/s\nGET: %s/s %d objs/s\n", humanize.Bytes(uploads*uint64(size)/uint64(durationSecs)), uploads/uint64(durationSecs), humanize.Bytes(downloads*uint64(size)/uint64(durationSecs)), downloads/uint64(durationSecs))
	return nil
}
