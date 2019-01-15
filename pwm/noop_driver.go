package pwm

import (
	"os"
	"sync"
)

type recorder struct {
	values map[string][]byte
}

func (r *recorder) Get(s string) []byte {
	return r.values[s]
}

func Noop() (Driver, *recorder) {
	rec := &recorder{values: make(map[string][]byte)}
	mu := new(sync.Mutex)
	d := &driver{
		sysfs: SysFS,
		writeFile: func(f string, c []byte, p os.FileMode) error {
			mu.Lock()
			defer mu.Unlock()
			rec.values[f] = c
			return nil
		},
	}
	return d, rec
}
