package persistence

import "time"
import "github.com/go-gorp/gorp"

type typeConverter struct{}

func (t typeConverter) ToDb(val interface{}) (converted interface{}, err error) {
	switch v := val.(type) {
	case time.Time:
		converted = uint64(v.Unix())
	case *time.Time:
		converted = uint64(v.Unix())
	default:
		converted = val
	}

	return
}

func (t typeConverter) FromDb(target interface{}) (scanner gorp.CustomScanner, useScanner bool) {
	switch target.(type) {
	case *time.Time:
		holder := uint64(0)

		scanner.Target = target
		scanner.Holder = &holder
		scanner.Binder = func(holder, target interface{}) (err error) {
			*target.(*time.Time) = time.Unix(int64(*holder.(*uint64)), 0)
			return
		}

		useScanner = true
	default:
		useScanner = false
	}

	return
}
