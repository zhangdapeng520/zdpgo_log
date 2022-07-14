package atomic

import "strconv"

//go:generate bin/gen-atomicwrapper -name=Float64 -type=float64 -wrapped=Uint64 -pack=math.Float64bits -unpack=math.Float64frombits -cas -json -imports math -file=float64.go

// Add atomically adds to the wrapped float64 and returns the new value.
func (f *Float64) Add(s float64) float64 {
	for {
		old := f.Load()
		new := old + s
		if f.CAS(old, new) {
			return new
		}
	}
}

// Sub atomically subtracts from the wrapped float64 and returns the new value.
func (f *Float64) Sub(s float64) float64 {
	return f.Add(-s)
}

// String encodes the wrapped value as a string.
func (f *Float64) String() string {
	// 'g' is the behavior for floats with %v.
	return strconv.FormatFloat(f.Load(), 'g', -1, 64)
}
