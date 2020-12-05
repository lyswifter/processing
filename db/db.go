package db

import (
	"github.com/ipfs/go-datastore"
	"github.com/lyswifter/processing/db/sectorstore"
	"github.com/lyswifter/processing/db/winning"
	"github.com/lyswifter/processing/repo"
	"golang.org/x/xerrors"
)

// SectorStore SectorStore
var SectorStore *sectorstore.SectorLifecycle

// WinningStore WinningStore
var WinningStore *winning.WinningLoop

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
