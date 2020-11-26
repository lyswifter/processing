package db

import (
	"github.com/ipfs/go-datastore"
	"github.com/lyswifter/processing/repo"
	"golang.org/x/xerrors"
)

// DealDs DealDs
var DealDs datastore.Batching

// PowerDs PowerDs
var PowerDs datastore.Batching

// WindowDs WindowDs
var WindowDs datastore.Batching

// WinningDs WinningDs
var WinningDs datastore.Batching

// SlaveDs SlaveDs
var SlaveDs datastore.Batching

// OpenDs OpenDs
func OpenDs(repoPath string, namespace string) (datastore.Batching, error) {

	r, err := repo.NewFS(repoPath)
	if err != nil {
		return nil, err
	}

	ok, err := r.Exists()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, xerrors.Errorf("repo at '%s' is not initialized", repoPath)
	}

	lr, err := r.Lock()
	if err != nil {
		return nil, err
	}

	// defer lr.Close()

	ds, err := lr.Datastore(namespace)
	if err != nil {
		return nil, err
	}

	return ds, nil
}
