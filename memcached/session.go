package memcached

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"github.com/syseditor/libeloula/utils"
)

var CacheManager *SessionManager

var cachePrefix = "player_session:" //can be anything

type SessionManager struct {
	cache *memcache.Client
}

type PlayerSession struct { //test player session settings
	Username     string
	UUID         string
	JoinedAt     int64
	BlocksBroken int64
}

func NewSessionManager(addr ...string) *SessionManager {
	gob.Register(PlayerSession{}) //register the PlayerSession type to gob

	return &SessionManager{
		cache: memcache.New(addr...),
	}
}

func (s *SessionManager) LoadOrCreate(uuid uuid.UUID, username string) (*PlayerSession, error) {
	key := cachePrefix + uuid.String()
	data, err := s.cache.Get(key)

	switch err {
	case nil: //no errors found, meaning no cache miss
		var session PlayerSession
		buffer := bytes.NewBuffer(data.Value)
		if err := gob.NewDecoder(buffer).Decode(&session); err != nil {
			return nil, err
		}

		return &session, nil
	case memcache.ErrCacheMiss: //cache miss, no session found
		newSession := &PlayerSession{ //new profile, basically wrong because here we need to make a db query and retreive the player data
			Username:     username,
			UUID:         uuid.String(),
			JoinedAt:     time.Now().Unix(),
			BlocksBroken: 0,
		}

		if err := s.Save(uuid, newSession); err != nil {
			return nil, err
		}

		return newSession, nil
	}

	return nil, err //in case something else is wrong
}

func (s *SessionManager) Load(uuid uuid.UUID) (*PlayerSession, error) { //only if we're sure the session has been created in cache before
	var session PlayerSession
	key := cachePrefix + uuid.String()
	data, _ := s.cache.Get(key)
	buffer := bytes.NewBuffer(data.Value)

	if err := gob.NewDecoder(buffer).Decode(&session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *SessionManager) Save(uuid uuid.UUID, ps *PlayerSession) error {
	key := cachePrefix + uuid.String()

	var buffer bytes.Buffer

	if err := gob.NewEncoder(&buffer).Encode(ps); err != nil {
		return err
	}

	return s.cache.Set(&memcache.Item{
		Key:        key,
		Value:      buffer.Bytes(),
		Expiration: 3600, //1 hour
	})
}

func (s *SessionManager) Delete(uuid uuid.UUID) error {
	return s.cache.Delete(cachePrefix + uuid.String())
}

func (s *SessionManager) AddBlocksBroken(uuid uuid.UUID) error {
	session, err := s.Load(uuid)
	utils.Check(err)

	if session != nil {
		session.BlocksBroken++
		return s.Save(uuid, session)
	} else {
		return errors.New("Call AddBlocksBroken() for non-existant session while loaded correctly")
	}
}
