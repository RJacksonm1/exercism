package robotname

import (
	"regexp"
	"testing"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)

func New() *Robot { return new(Robot) }

func TestNameValid(t *testing.T) {
	n := New().Name()
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.Name()
	n2 := r.Name()
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().Name()
	if New().Name() == n1 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.Name()
	r.Reset()
	if r.Name() == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

func TestCollisions(t *testing.T) {
	m := map[string]bool{}
	// Test uniqueness for new robots.
	// 10k should be plenty to catch names generated
	// randomly without a uniqueness check.
	for i := 0; i < 10000; i++ {
		n := New().Name()
		if m[n] {
			t.Fatalf("Name %s reissued after %d robots.", n, i)
		}
		m[n] = true
	}
	// Test that names aren't recycled either.
	r := New()
	for i := 0; i < 10000; i++ {
		r.Reset()
		n := r.Name()
		if m[n] {
			t.Fatalf("Name %s reissued after Reset.", n)
		}
		m[n] = true
	}
}

// Note if you go for bonus points, this benchmark likely won't be
// meaningful.  Bonus thought exercise, why won't it be meaningful?
func BenchmarkName(b *testing.B) {
	// Benchmark combined time to create robot and name.
	for i := 0; i < b.N; i++ {
		New().Name()
	}
}
