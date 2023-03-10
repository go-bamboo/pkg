package maths

import (
	"testing"
)

func TestMaxMin(t *testing.T) {
	f32Slice := []float32{1.1, 2.1, 3.1}
	f32 := Float32Var.Max(f32Slice...)
	if f32 != f32Slice[len(f32Slice)-1] {
		t.Error("float32 Max error", f32)
	}
	f32 = Float32Var.Min(f32Slice...)
	if f32 != f32Slice[0] {
		t.Error("float32 Min error", f32)
	}

	f64Slice := []float64{1.1, 2.1, 3.1}
	f64 := Float64Var.Max(f64Slice...)
	if f64 != f64Slice[len(f64Slice)-1] {
		t.Error("float64 Max error", f64)
	}
	f64 = Float64Var.Min(f64Slice...)
	if f64 != f64Slice[0] {
		t.Error("float64 Min error", f64)
	}

	iSlice := []int{1, 2, 3}
	i := IntVar.Max(iSlice...)
	if i != iSlice[len(iSlice)-1] {
		t.Error("int Max error", i)
	}
	i = IntVar.Min(iSlice...)
	if i != iSlice[0] {
		t.Error("int Min error", i)
	}

	uiSlice := []uint{1, 2, 3}
	ui := UintVar.Max(uiSlice...)
	if ui != uiSlice[len(uiSlice)-1] {
		t.Error("uint Max error", ui)
	}
	ui = UintVar.Min(uiSlice...)
	if ui != uiSlice[0] {
		t.Error("uint Min error", ui)
	}

	strSlice := []string{"a", "b", "c"}
	s := StringVar.Max(strSlice...)
	if s != strSlice[len(strSlice)-1] {
		t.Error("string Max error", s)
	}
	s = StringVar.Min(strSlice...)
	if s != strSlice[0] {
		t.Error("string Min error", s)
	}

}

func TestRange(t *testing.T) {
	sum, size := 0, 10
	for range Range(size) {
		sum++
	}
	if sum != size {
		t.Error("Range error")
	}
}

func TestTernary(t *testing.T) {
	r := Ternary(2 > 1, 2, 1)
	if r.(int) != 2 {
		t.Error("Ternary error")
	}

	s := Ternary("is andy2046 handsome?" == "yes", "no", "yes")
	if s.(string) != "yes" {
		t.Error("Ternary error")
	}
}
