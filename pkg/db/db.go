package db

//Db - ..........
type Db interface {
	Merge(key string, value []byte) error
	Get(key string) ([]byte, error)
}


