package repo

import (
	"os"
	"path/filepath"

	logging "github.com/ipfs/go-log/v2"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"

	config "github.com/lyswifter/processing/config"
)

var log = logging.Logger("repo")

const (
	fsAPI       = "api"
	fsConfig    = "config.toml"
	fsDatastore = "datastore"
	fsKeystore  = "keystore"
	fsLock      = "repo.lock"
)

// FsRepo FsRepo
type FsRepo struct {
	path       string
	configPath string
}

// NewFS creates a repo instance based on a path on file system
func NewFS(path string) (*FsRepo, error) {
	path, err := homedir.Expand(path)
	if err != nil {
		return nil, err
	}

	return &FsRepo{
		path:       path,
		configPath: filepath.Join(path, fsConfig),
	}, nil
}

// Path path
func (fsr *FsRepo) Path() string {
	return fsr.path
}

// SetConfigPath SetConfigPath
func (fsr *FsRepo) SetConfigPath(cfgPath string) {
	fsr.configPath = cfgPath
}

// Exists Exists
func (fsr *FsRepo) Exists() (bool, error) {
	// _, err := os.Stat(filepath.Join(fsr.path, fsDatastore))
	// notexist := os.IsNotExist(err)
	// if notexist {
	// 	err = nil
	// }
	// return !notexist, err
	return true, nil
}

func (fsr *FsRepo) Init() error {
	exist, err := fsr.Exists()
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	log.Infof("Initializing repo at '%s'", fsr.path)
	err = os.MkdirAll(fsr.path, 0755) //nolint: gosec
	if err != nil && !os.IsExist(err) {
		return err
	}

	if err := fsr.initConfig(); err != nil {
		return xerrors.Errorf("init config: %w", err)
	}

	return nil

}

func (fsr *FsRepo) initConfig() error {
	_, err := os.Stat(fsr.configPath)
	if err == nil {
		// exists
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	c, err := os.Create(fsr.configPath)
	if err != nil {
		return err
	}

	comm, err := config.ConfigComment(config.DefaultProcessingNode())
	if err != nil {
		return err
	}

	_, err = c.Write(comm)
	if err != nil {
		return xerrors.Errorf("write config: %w", err)
	}

	if err := c.Close(); err != nil {
		return xerrors.Errorf("close config: %w", err)
	}
	return nil
}
