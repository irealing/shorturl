package shorturl

import (
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	"strconv"
)

type ShortedHandler struct {
	db *leveldb.DB
}

func NewHandler(fp string) (*ShortedHandler, error) {
	db, err := leveldb.OpenFile(fp, nil)
	if err != nil {
		return nil, err
	}
	return &ShortedHandler{db: db}, nil
}
func (sh *ShortedHandler) Create(url string) (*ShortedURL, error) {
	u := NewShortedURL(url)
	tran, err := sh.db.OpenTransaction()
	if err != nil {
		return nil, err
	}
	var idx = 0
	ck := sh.counterKey(u.Hash)
	if data, err := sh.db.Get(ck, nil); err != nil && data != nil {
		idx, err = strconv.Atoi(string(data))
		if err != nil {
			tran.Discard()
			return nil, err
		}
	}
	u.Hash.Index = uint32(idx)
	if idx > 0 {
		u.ReMake()
	}
	idx++
	cv := []byte(strconv.Itoa(idx))
	tran.Put(ck, cv, nil)
	tran.Put(sh.key(u.Hash), []byte(u.URL), nil)
	err = tran.Commit()
	return u, err
}
func (sh *ShortedHandler) counterKey(hash *URLHash) []byte {
	s := fmt.Sprintf("idx::%d", hash.Hash)
	return []byte(s)
}
func (sh *ShortedHandler) key(hash *URLHash) []byte {
	s := fmt.Sprintf("record::%d:%d", hash.Hash, hash.Index)
	return []byte(s)
}
func (sh *ShortedHandler) Find(s string) (*ShortedURL, error) {
	hash, err := Decode(s)
	if err != nil {
		return nil, err
	}
	key := sh.key(hash)
	data, err := sh.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return &ShortedURL{URL: string(data), Hash: hash, Shorted: s}, nil
}
