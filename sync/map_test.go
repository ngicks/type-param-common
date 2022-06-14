package sync_test

import (
	"bytes"
	"reflect"
	sync_ "sync"
	"testing"
	"time"

	"github.com/ngicks/type-param-common/sync"
)

type kv[K, V any] struct {
	Key   K
	Value V
}

func mapTestSet[K, V any](t *testing.T, m sync.Map[K, V], keyValues ...kv[K, V]) {
	if len(keyValues) < 3 {
		panic("keyValues must be length of 3 or more")
	}

	lastEle := keyValues[len(keyValues)-1]
	keyValues = keyValues[:len(keyValues)-1]

	// Store
	for _, kv := range keyValues {
		m.Store(kv.Key, kv.Value)
	}

	// Load
	for _, kv := range keyValues {
		v, ok := m.Load(kv.Key)
		if !ok {
			t.Errorf("cloud not load %v", kv.Key)
		}
		if !reflect.DeepEqual(v, kv.Value) {
			t.Fatalf("mismatched stored value, stored: %v, loaded: %v", kv.Value, v)
		}
	}
	v, ok := m.Load(lastEle.Key)
	if ok {
		t.Errorf("loaded unstored key: %v", lastEle.Key)
	}
	if !reflect.ValueOf(v).IsZero() {
		t.Fatalf("mismatched value, must be zero value, but loaded: %v", v)
	}

	// Range
	m.Range(func(key K, value V) bool {
		for _, kv := range keyValues {
			if reflect.DeepEqual(key, kv.Key) && reflect.DeepEqual(value, kv.Value) {
				return true
			}
		}
		t.Fatalf("Range passing incorrect value")
		return true
	})

	// LoadOrStore
	if v, loaded := m.LoadOrStore(keyValues[0].Key, keyValues[0].Value); !reflect.DeepEqual(v, keyValues[0].Value) || !loaded {
		t.Fatalf("must be stored but could not load")
	}
	if v, loaded := m.LoadOrStore(keyValues[0].Key, keyValues[1].Value); !reflect.DeepEqual(v, keyValues[0].Value) || !loaded {
		t.Fatalf("must be stored but could not load")
	}
	if v, _ := m.LoadOrStore(lastEle.Key, lastEle.Value); !reflect.DeepEqual(v, lastEle.Value) {
		t.Fatalf("mismatched stored value, stored: %v, loaded: %v", lastEle.Value, v)
	}

	// LoadAndDelete
	if v, loaded := m.LoadAndDelete(keyValues[0].Key); !reflect.DeepEqual(v, keyValues[0].Value) || !loaded {
		t.Fatalf("mismatched stored value, stored: %v, loaded: %v", keyValues[0].Value, v)
	}
	if v, loaded := m.LoadAndDelete(keyValues[0].Key); !reflect.ValueOf(v).IsZero() || loaded {
		t.Fatalf("must be deleted already and loaded value must be zero. loaded: %v", v)
	}
	if _, loaded := m.Load(keyValues[0].Key); loaded {
		t.Fatalf("must be deleted")
	}

	// Delete
	if _, loaded := m.Load(keyValues[1].Key); !loaded {
		t.Fatalf("must be stored")
	}
	m.Delete(keyValues[1].Key)
	if _, loaded := m.Load(keyValues[1].Key); loaded {
		t.Fatalf("must be deleted")
	}
}

func TestMap(t *testing.T) {
	m := sync.Map[string, time.Time]{}
	now := time.Now()
	mapTestSet(t, m,
		[]kv[string, time.Time]{
			{"foo", now},
			{"bar", now.Add(time.Hour)},
			{"baz", now.Add(2 * time.Hour)},
		}...,
	)
}

func TestMapRace(t *testing.T) {
	m := sync.Map[time.Time, *bytes.Buffer]{}
	now := time.Now()

	randomMethodCall := func() {
		keyOrder := map[int]any{
			0: nil,
			1: nil,
			2: nil,
			3: nil,
			4: nil,
			5: nil,
		}
		for key := range keyOrder {
			switch key {
			case 0:
				m.Delete(now)
			case 1:
				m.Load(now)
			case 2:
				m.LoadAndDelete(now)
			case 3:
				m.LoadOrStore(now, bytes.NewBuffer([]byte{}))
			case 4:
				m.Range(func(key time.Time, value *bytes.Buffer) bool {
					return true
				})
			case 5:
				m.Store(now, bytes.NewBuffer([]byte{}))
			}
		}
	}

	wg := sync_.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			randomMethodCall()
			wg.Done()
		}()
	}
	wg.Wait()
}
