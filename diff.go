// Package dockerdiff +build linux
package dockerdiff

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/distribution"
	"github.com/docker/docker/client"
	"github.com/docker/docker/image"
	"github.com/docker/docker/layer"
	"github.com/docker/docker/pkg/archive"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type manifestItem struct {
	Config       string
	RepoTags     []string
	Layers       []string
	Parent       image.ID                                 `json:",omitempty"`
	LayerSources map[layer.DiffID]distribution.Descriptor `json:",omitempty"`
}

// ImageExport 导出镜像到dst路径下(已解压)
func ImageExport(cli *client.Client, ImageID string, dst string) error {
	responseBody, err := cli.ImageSave(context.Background(), []string{ImageID})
	if err != nil {
		return err
	}
	defer responseBody.Close()

	if err = archive.Untar(responseBody, dst, nil); err != nil {
		return err
	}
	return nil
}

// ImageTar 镜像包解包
func ImageTar(path string, output io.Writer) error {
	fs, err := archive.Tar(path, archive.Uncompressed)
	if err != nil {
		return err
	}
	defer fs.Close()

	if _, err = io.Copy(output, fs); err != nil {
		return err
	}
	return nil
}

// DiffExport 经过对比删减后的镜像
func DiffExport(cli *client.Client, image1Name string, image2Name string, output io.Writer) error {
	image1, _, err := cli.ImageInspectWithRaw(context.Background(), image1Name)
	if err != nil {
		return err
	}
	image2, _, err := cli.ImageInspectWithRaw(context.Background(), image2Name)
	if err != nil {
		return err
	}

	if len(image2.RepoTags) == 0 {
		return fmt.Errorf("%s repo tags is empty", image2Name)
	}

	// 查找重复文件层
	var duplicateLayer []string
	for i, l := range image1.RootFS.Layers {
		if i > len(image2.RootFS.Layers)-1 || l != image2.RootFS.Layers[i] {
			break
		}
		duplicateLayer = append(duplicateLayer, l)
	}

	tmpDir, err := os.MkdirTemp("", "docker-import-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	// 导出image2
	if err = ImageExport(cli, image2.RepoTags[0], tmpDir); err != nil {
		return err
	}

	manifestFile, err := os.Open(filepath.Join(tmpDir, "manifest.json"))
	if err != nil {
		return err
	}
	defer manifestFile.Close()
	var manifest []manifestItem
	if err = json.NewDecoder(manifestFile).Decode(&manifest); err != nil {
		return err
	}
	var layers []string
	for _, t := range manifest {
		if t.Config == strings.TrimPrefix(image2.ID, "sha256:")+".json" {
			layers = t.Layers
		}
	}

	// 删除image2重复层
	for i := range duplicateLayer {
		unlessLayer := layers[i]
		unlessLayer = strings.TrimPrefix(unlessLayer, "sha256:")
		unlessLayerFile := filepath.Join(tmpDir, unlessLayer)
		if _, err = os.OpenFile(unlessLayerFile, os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
			return err
		}
	}

	// 压包
	if err = ImageTar(tmpDir, output); err != nil {
		return err
	}

	return nil
}
