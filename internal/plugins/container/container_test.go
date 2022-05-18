package container

import "testing"

type sampleA struct {
	v string
}

type sampleB struct {
	v string
}

func Test001(t *testing.T) {
	var a = sampleA{v: "a"}
	var b = sampleB{v: "b"}

	Add(&a, "sample-a")
	Add(&b)

	var err error
	var aa *sampleA
	var bb *sampleB
	t.Log("get a")
	aa, err = Get(new(sampleA), "sample-a")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(aa.v)

	t.Log("get b")
	bb, err = Get(bb)

	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(bb.v)

	t.Log("get b2")
	bb, err = Get(bb, unique)

	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(bb.v)
}
