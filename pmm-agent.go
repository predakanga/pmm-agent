//go:generate make api

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/percona/pmm-agent/app"
	"github.com/percona/pmm-agent/cmd"
)

func main() {
	if err := cmd.New(context.Background(), app.App{}).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
