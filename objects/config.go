package objects

type Config struct {
	Key   string `gorm:"primaryKey"`
	Value string
}
