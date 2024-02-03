package database

type Store struct {
}

func New() Store {
	return Store{}
}

func (s *Store) Close() {

}
