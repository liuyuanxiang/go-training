package address

type RepoInterface interface {
	CreateAddr(addr *Address) (string, error)

	UpdateAddr(id string, addr *Address) error

	DeleteAddr(id string) error

	GetAddr(id string) (*Address, error)

	ListAddr(pageNum, pageSize int64) ([]*Address, error)
}
