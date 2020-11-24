package repo

import (
	dgbadger "github.com/dgraph-io/badger/v2"
	ldbopts "github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/ipfs/go-datastore"
	badger "github.com/ipfs/go-ds-badger2"
	levelds "github.com/ipfs/go-ds-leveldb"
)

type dsCtor func(path string, readonly bool) (datastore.Batching, error)

// ChainBadgerOptions ChainBadgerOptions
func ChainBadgerOptions() badger.Options {
	opts := badger.DefaultOptions
	opts.GcInterval = 0 // disable GC for chain datastore

	opts.Options = dgbadger.DefaultOptions("").WithTruncate(true).
		WithValueThreshold(128)
	return opts
}

var fsDatastores = map[string]dsCtor{
	"deal":         chainBadgerDs,
	"power":        chainBadgerDs,
	"slave":        chainBadgerDs,
	"uploader":     chainBadgerDs,
	"window":       chainBadgerDs,
	"winning":      chainBadgerDs,
	"poster":       chainBadgerDs,
	"osp-provider": chainBadgerDs,
	"osp-worker":   chainBadgerDs,
}

func chainBadgerDs(path string, readonly bool) (datastore.Batching, error) {
	opts := ChainBadgerOptions()
	opts.ReadOnly = readonly
	return badger.NewDatastore(path, &opts)
}

func badgerDs(path string, readonly bool) (datastore.Batching, error) {
	opts := badger.DefaultOptions
	opts.ReadOnly = readonly

	opts.Options = dgbadger.DefaultOptions("").WithTruncate(true).
		WithValueThreshold(1 << 10)
	return badger.NewDatastore(path, &opts)
}

func levelDs(path string, readonly bool) (datastore.Batching, error) {
	return levelds.NewDatastore(path, &levelds.Options{
		Compression: ldbopts.NoCompression,
		NoSync:      false,
		Strict:      ldbopts.StrictAll,
		ReadOnly:    readonly,
	})
}
