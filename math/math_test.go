package math

import "testing"

func TestFloatCompare(t *testing.T) {
	eq := FloatCompare(1.01, 1.02, 0.02)
	lt := FloatCompare(1.01, 1.02, 0.0001)
	gt := FloatCompare(1.01, 1.0, 0.001)

	if eq != 0 || lt != -1 || gt != 1 {
		t.FailNow()
	}
}

func TestDegreesToRadians(t *testing.T) {
	a, b, c, d := 0.0, 90.0, 180.0, 200.0
	aaa, bbb, ccc, ddd := 0.0, 1.5708, 3.14159, 3.49066

	aa := DegreesToRadians(a)
	bb := DegreesToRadians(b)
	cc := DegreesToRadians(c)
	dd := DegreesToRadians(d)

	th := 0.001
	ar := FloatCompare(aa, aaa, th)
	br := FloatCompare(bb, bbb, th)
	cr := FloatCompare(cc, ccc, th)
	dr := FloatCompare(dd, ddd, th)

	if ar != 0 || br != 0 || cr != 0 || dr != 0 {
		t.FailNow()
	}

}

func TestRadiansToDegrees(t *testing.T) {
	aaa, bbb, ccc, ddd := 0.0, 90.0, 180.0, 200.0
	a, b, c, d := 0.0, 1.5708, 3.14159, 3.49066

	aa := RadiansToDegrees(a)
	bb := RadiansToDegrees(b)
	cc := RadiansToDegrees(c)
	dd := RadiansToDegrees(d)

	th := 0.001
	ar := FloatCompare(aa, aaa, th)
	br := FloatCompare(bb, bbb, th)
	cr := FloatCompare(cc, ccc, th)
	dr := FloatCompare(dd, ddd, th)

	if ar != 0 || br != 0 || cr != 0 || dr != 0 {
		t.FailNow()
	}

}
