package user

import "github.com/kronos1209/learn-grpc-with-go/internal/entities"

type UserRepository interface {
	// 指定した id を持つユーザを取得する
	Get(id int64) (*entities.User, error)
	// 指定した id を持つユーザリストを取得する
	GetAll(ids ...int64) ([]*entities.User, error)
	// リポジトリにユーザを追加する
	Add(user *entities.User) error
	// ユーザ情報を更新する
	Update(user *entities.User) error
	// リポジトリからユーザを削除する
	Del(id int64) (*entities.User, error)
}

func NewMemoryRepository() *MemoryRepository {
	r := MemoryRepository{
		currentSequentialId: 0,
		repository:          make(map[int64]*entities.User),
	}
	return &r
}
