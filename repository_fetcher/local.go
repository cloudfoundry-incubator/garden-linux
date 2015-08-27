package repository_fetcher

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/layercake"
	"github.com/cloudfoundry-incubator/garden-linux/process"
	"github.com/docker/docker/image"
	"github.com/docker/docker/pkg/archive"
	"github.com/pivotal-golang/lager"
)

type ContainerIDProvider interface {
	ProvideID(path string) layercake.IDer
}

type Local struct {
	Cake              layercake.Cake
	DefaultRootFSPath string
	IDProvider        ContainerIDProvider

	mu sync.RWMutex
}

func (l *Local) Fetch(
	logger lager.Logger,
	repoURL *url.URL,
	tag string,
	_ int64,
) (string, process.Env, []string, error) {

	path := repoURL.Path
	if len(path) == 0 {
		path = l.DefaultRootFSPath
	}

	if len(path) == 0 {
		return "", nil, nil, errors.New("RootFSPath: is a required parameter, since no default rootfs was provided to the server. To provide a default rootfs, use the --rootfs flag on startup.")
	}

	id, err := l.fetch(path)
	return id, nil, nil, err
}

func (l *Local) fetch(path string) (string, error) {
	path, err := resolve(path)
	if err != nil {
		panic(err)
	}

	id := l.IDProvider.ProvideID(path)

	// synchronize all downloads, we could optimize by only mutexing around each
	// particular rootfs path, but in practice importing local rootfses is decently fast,
	// and concurrently importing local rootfses is rare.
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, err := l.Cake.Get(id); err == nil {
		return id.ID(), nil // use cache
	}

	tar, err := archive.Tar(path, archive.Uncompressed)
	if err != nil {
		return "", fmt.Errorf("repository_fetcher: fetch local rootfs: untar rootfs: %v", err)
	}
	defer tar.Close()

	if err := l.Cake.Register(&image.Image{ID: id.ID()}, tar); err != nil {
		return "", fmt.Errorf("repository_fetcher: fetch local rootfs: register rootfs: %v", err)
	}

	return id.ID(), nil
}

func resolve(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("repository_fetcher: stat file: %s", err)
	}

	if (fileInfo.Mode() & os.ModeSymlink) == os.ModeSymlink {
		if path, err = os.Readlink(path); err != nil {
			return "", fmt.Errorf("repository_fetcher: read link: %s", err)
		}
	}
	return path, nil
}

type LayerIDProvider struct {
}

func (LayerIDProvider) ProvideID(id string) layercake.IDer {

	info, err := os.Lstat(id)
	if err != nil {
		return layercake.ContainerID(fmt.Sprintf("%s-file-doesn't-yet-exist", id))
	}

	return layercake.ContainerID(fmt.Sprintf("%s-%d", id, info.ModTime().UnixNano()))
}
