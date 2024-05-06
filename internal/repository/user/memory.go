package user

import (
	"fmt"
	"sync"

	"github.com/kronos1209/learn-grpc-with-go/internal/entities"
)

var ErrNotFoundUser = "User not found. 【id => %d】"
var ErrAlreadExistUser = "User already exist. 【id => %d】"

type MemoryRepository struct {
	repository          map[int64]*entities.User
	currentSequentialId int64
	sync.RWMutex
}

var _ UserRepository = (*MemoryRepository)(nil)

func (m *MemoryRepository) generateSequentialId() int64 {
	m.currentSequentialId++
	return m.currentSequentialId
}

func (m *MemoryRepository) Get(id int64) (*entities.User, error) {
	m.RLock()
	defer m.RUnlock()

	u, ok := m.repository[id]
	if !ok {
		return nil, fmt.Errorf(ErrNotFoundUser, id)
	}
	return u, nil
}

func (m *MemoryRepository) GetAll(ids ...int64) ([]*entities.User, error) {
	m.RLock()
	defer m.RUnlock()

	res := make([]*entities.User, 0, len(ids))
	for _, id := range ids {
		if u, ok := m.repository[id]; ok {
			res = append(res, u)
		}
	}
	return res, nil
}

func (m *MemoryRepository) Add(user *entities.User) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.repository[user.ID]; ok {
		return fmt.Errorf(ErrAlreadExistUser, user.ID)
	}
	sequentialId := m.generateSequentialId()
	user.ID = sequentialId
	m.repository[user.ID] = user
	return nil
}

func (m *MemoryRepository) Update(user *entities.User) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.repository[user.ID]; !ok {
		return fmt.Errorf(ErrNotFoundUser, user.ID)
	}
	m.repository[user.ID] = user
	return nil
}
func (m *MemoryRepository) Del(id int64) (*entities.User, error) {
	m.Lock()
	defer m.Unlock()

	u, ok := m.repository[id]
	if !ok {
		return nil, fmt.Errorf(ErrNotFoundUser, id)
	}

	delete(m.repository, id)
	return u, nil
}
