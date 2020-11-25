package repo

import (
	"errors"

	"github.com/ipfs/go-datastore"
)

var (
	ErrNoAPIEndpoint     = errors.New("API not running (no endpoint)")
	ErrNoAPIToken        = errors.New("API token not set")
	ErrRepoAlreadyLocked = errors.New("repo is already locked")
	ErrClosedRepo        = errors.New("repo is no longer open")

	// ErrInvalidBlockstoreDomain is returned by LockedRepo#Blockstore() when
	// an unrecognized domain is requested.
	ErrInvalidBlockstoreDomain = errors.New("invalid blockstore domain")
)

// Repo Repo
type Repo interface {
	// Lock locks the repo for exclusive use.
	Lock() (LockedRepo, error)
}

// LockedRepo LockedRepo
type LockedRepo interface {
	// Close closes repo and removes lock.
	Close() error

	// Returns datastore defined in this repo.
	Datastore(namespace string) (datastore.Batching, error)

	// Returns config in this repo
	Config() (interface{}, error)
	SetConfig(func(interface{})) error

	// Path returns absolute path of the repo
	Path() string
}
