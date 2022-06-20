package set

import (
	"testing"
	"time"

	"github.com/ngicks/type-param-common/slice"
)

var utc, jst *time.Location

func init() {
	var err error
	utc, err = time.LoadLocation("")
	if err != nil {
		panic(err)
	}
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func TestSet(t *testing.T) {
	set := Set[time.Time]{}

	now := time.Now()
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	nanosecond := now.Nanosecond()

	tgt := time.Date(year, month, day, hour, minute, second, nanosecond, utc)
	tgtInOtherLocation := time.Date(year, month, day, hour, minute, second, nanosecond, jst)

	if set.Len() != 0 {
		t.Fatalf("wrong len")
	}
	set.Add(tgt)
	if set.Len() != 1 {
		t.Fatalf("wrong len")
	}
	set.Add(tgtInOtherLocation)
	if set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	set.Add(tgt)
	if set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	if !set.Has(tgt) {
		t.Fatalf("must have")
	}
	if set.Has(now) {
		t.Fatalf("must not have")
	}

	d := make([]time.Time, 0)
	iter := set.Values()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		d = append(d, next)
	}

	if !(slice.Has(d, tgt) && slice.Has(d, tgtInOtherLocation)) {
		t.Fatalf("must contain")
	}

	d = d[:]
	set.ForEach(func(v time.Time) {
		d = append(d, v)
	})
	if !(slice.Has(d, tgt) && slice.Has(d, tgtInOtherLocation)) {
		t.Fatalf("must contain")
	}

	if set.Delete(now); set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	if set.Delete(tgt); set.Len() != 1 {
		t.Fatalf("wrong len")
	}

	set.Add(now.Add(time.Second))
	set.Add(now.Add(2 * time.Second))
	set.Add(now.Add(3 * time.Second))
	set.Clear()
	if set.Len() != 0 {
		t.Fatalf("wrong len")
	}
}
