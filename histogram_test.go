package metrics

import (
	"testing"
)

func TestEmptyHistogram(t *testing.T) {
	h := NewHistogram(NewUniformSample(100))
	if count := h.Count(); 0 != count {
		t.Errorf("h.Count(): 0 != %v\n", count)
	}
	if min := h.Min(); 0 != min {
		t.Errorf("h.Min(): 0 != %v\n", min)
	}
	if max := h.Max(); 0 != max {
		t.Errorf("h.Max(): 0 != %v\n", max)
	}
	if mean := h.Mean(); 0.0 != mean {
		t.Errorf("h.Mean(): 0.0 != %v\n", mean)
	}
	if stdDev := h.StdDev(); 0.0 != stdDev {
		t.Errorf("h.StdDev(): 0.0 != %v\n", stdDev)
	}
	ps := h.Percentiles([]float64{0.5, 0.75, 0.99})
	if 0.0 != ps[0] {
		t.Errorf("median: 0.0 != %v\n", ps[0])
	}
	if 0.0 != ps[1] {
		t.Errorf("75th percentile: 0.0 != %v\n", ps[1])
	}
	if 0.0 != ps[2] {
		t.Errorf("99th percentile: 0.0 != %v\n", ps[2])
	}
}

func TestHistogram10000(t *testing.T) {
	h := NewHistogram(NewUniformSample(100000))
	for i := 1; i <= 10000; i++ { h.Update(int64(i)) }
	if count := h.Count(); 10000 != count {
		t.Errorf("h.Count(): 10000 != %v\n", count)
	}
	if min := h.Min(); 1 != min {
		t.Errorf("h.Min(): 1 != %v\n", min)
	}
	if max := h.Max(); 10000 != max {
		t.Errorf("h.Max(): 10000 != %v\n", max)
	}
	if mean := h.Mean(); 5000.5 != mean {
		t.Errorf("h.Mean(): 5000.5 != %v\n", mean)
	}
	if stdDev := h.StdDev(); 2886.8956799071675 != stdDev {
		t.Errorf("h.StdDev(): 2886.8956799071675 != %v\n", stdDev)
	}
	ps := h.Percentiles([]float64{0.5, 0.75, 0.99})
	if 5000.5 != ps[0] {
		t.Errorf("median: 5000.5 != %v\n", ps[0])
	}
	if 7500.75 != ps[1] {
		t.Errorf("75th percentile: 7500.75 != %v\n", ps[1])
	}
	if 9900.99 != ps[2] {
		t.Errorf("99th percentile: 9900.99 != %v\n", ps[2])
	}
}
