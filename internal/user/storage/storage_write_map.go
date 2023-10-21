package storage

import (
	"fmt"
	"sync"

	"github.com/LNMMusic/msauth/internal/user"
	"github.com/LNMMusic/optional"
)

// NewStorageWriteMap creates a new instance of StorageWriteMap
func NewStorageWriteMap(db *sync.Map, lastId int) *StorageWriteMap {
	// default config
	defaultDb := &sync.Map{}
	defaultLastId := 0
	if db != nil {
		defaultDb = db
	}
	if lastId > 0 {
		defaultLastId = lastId
	}

	return &StorageWriteMap{
		db: defaultDb,
		lastId: defaultLastId,
		mu: &sync.Mutex{},
	}
}

// StorageWriteMap is the storage in map format
type StorageWriteMap struct {
	// db is the database in map format
	// - key: user id
	// - value: user data
	db *sync.Map

	// lastId is the last id of the user
	lastId int

	// mu is the mutex
	mu *sync.Mutex
}

// Create a new user
func (m *StorageWriteMap) Create(u *user.User) (err error) {
	// lock the mutex
	m.mu.Lock()
	defer m.mu.Unlock()
	
	// set the id
	m.lastId++
	(*u).Id = m.lastId

	// save the user
	m.db.Store(m.lastId, u)

	return
}

// Update an existing user
func (m *StorageWriteMap) Update(u *user.User) (err error) {
	// update the user
	_, ok := m.db.Swap(u.Id, u)
	if !ok {
		err = fmt.Errorf("%w - %d", ErrStorageNotFound, u.Id)
		return
	}

	return
}

// Delete an existing user
func (m *StorageWriteMap) Delete(id int) (err error) {
	// delete the user
	_, ok := m.db.LoadAndDelete(id)
	if !ok {
		err = fmt.Errorf("%w - %d", ErrStorageNotFound, id)
		return
	}

	return
}

// Activate an existing user
func (m *StorageWriteMap) Activate(id int) (err error) {
	// lock the mutex
	m.mu.Lock()
	defer m.mu.Unlock()

	// get the user
	v, ok := m.db.Load(id)
	if !ok {
		err = fmt.Errorf("%w - %d", ErrStorageNotFound, id)
		return
	}

	// type assertion
	u := v.(user.User)

	// activate the user
	u.IsActive = optional.Some[bool](true)

	// update the user
	m.db.Swap(u.Id, u)

	return
}

