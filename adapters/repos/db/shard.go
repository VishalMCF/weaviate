package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Shard is the smallest completely-contained index unit. A shard mananages
// database files for all the objects it owns. How a shard is determined for a
// target object (e.g. Murmur hash, etc.) is still open at this point
type Shard struct {
	index *Index // a reference to the underlying index, which in turn contains schema information
	name  string
	db    *bolt.DB // one db file per shard, uses buckets for separation between data storage, index storage, etc.
}

var (
	ObjectsBucket []byte = []byte("objects")
)

func NewShard(shardName string, index *Index) (*Shard, error) {
	s := &Shard{
		index: index,
		name:  shardName,
	}

	err := s.initDBFile()
	if err != nil {
		return nil, errors.Wrapf(err, "init shard %s", s.ID())
	}

	return s, nil
}

func (s *Shard) ID() string {
	return fmt.Sprintf("%s_%s", s.index.ID(), s.name)
}

func (s *Shard) DBPath() string {
	return fmt.Sprintf("%s/%s.db", s.index.Config.RootPath, s.ID())
}

func (s *Shard) initDBFile() error {
	boltdb, err := bolt.Open(s.DBPath(), 0600, nil)
	if err != nil {
		return errors.Wrapf(err, "open bolt at %s", s.DBPath())
	}

	err = boltdb.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(ObjectsBucket); err != nil {
			return errors.Wrapf(err, "create objects bucket '%s'", string(ObjectsBucket))
		}

		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "create bolt buckets")
	}

	s.db = boltdb
	return nil
}

func (s *Shard) putObject(ctx context.Context, object *KindObject) error {
	buf := bytes.NewBuffer(nil)
	// TODO: optimize storage. For example it makes no sense to store the vector
	// as json (which is a string of floats), instead we could just store the raw
	// vector and safe on space. Similarly we will want to add other relevant
	// fields, such as the unique shard id, etc.
	json.NewEncoder(buf).Encode(object)
	idBytes, err := uuid.MustParse(object.ID().String()).MarshalBinary()
	if err != nil {
		return err
	}

	err = s.db.Batch(func(tx *bolt.Tx) error {
		return tx.Bucket(ObjectsBucket).Put([]byte(idBytes), buf.Bytes())
	})
	if err != nil {
		return errors.Wrap(err, "bolt batch tx")
	}

	return nil
}
