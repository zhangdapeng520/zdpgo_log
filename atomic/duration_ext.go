package atomic

import "time"

//go:generate bin/gen-atomicwrapper -name=Duration -type=time.Duration -wrapped=Int64 -pack=int64 -unpack=time.Duration -cas -swap -json -imports time -file=duration.go

// Add atomically adds to the wrapped time.Duration and returns the new value.
func (d *Duration) Add(n time.Duration) time.Duration {
	return time.Duration(d.v.Add(int64(n)))
}

// Sub atomically subtracts from the wrapped time.Duration and returns the new value.
func (d *Duration) Sub(n time.Duration) time.Duration {
	return time.Duration(d.v.Sub(int64(n)))
}

// String encodes the wrapped value as a string.
func (d *Duration) String() string {
	return d.Load().String()
}
