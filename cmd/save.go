package main

import (
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"github.com/zn-chen/dockerdiff"
	"io"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:     "save",
	Example: "dockerdiff save IMAGE1 IMAGE2 -o dst.tar",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			_ = cmd.Help()
			log.Fatal(err)
		}
		if err := DiffExport(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
		}
	},
}

func DiffExport(image1, image2 string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	var outputStream io.Writer
	if output == "" {
		outputStream = os.Stdout
	} else {
		if fd, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 600); err != nil {
			return err
		} else {
			defer fd.Close()
			outputStream = fd
		}
	}

	if err = dockerdiff.DiffExport(cli, image1, image2, outputStream); err != nil {
		return err
	}
	return nil
}
