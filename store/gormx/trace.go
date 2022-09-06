package gormx

import "gorm.io/gorm"

type DBTextMapCarrier struct {
	db *gorm.DB
}

// Get returns the value associated with the passed key.
func (carrier *DBTextMapCarrier) Get(key string) string {
	v, ok := carrier.db.InstanceGet(key)
	vv := v.(string)
	if ok && len(vv) > 0 {
		return vv
	}
	return ""
}

// Set stores the key-value pair.
func (carrier *DBTextMapCarrier) Set(key string, value string) {
	carrier.db.InstanceSet(key, value)
}

// Keys lists the keys stored in this carrier.
func (carrier *DBTextMapCarrier) Keys() []string {
	return nil
}
