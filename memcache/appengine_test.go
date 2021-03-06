// +build appengine

package memcache

import (
	"bytes"
	"testing"

	"appengine/aetest"
)

func TestAppEngine(t *testing.T) {
	ctx, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	cache := New(ctx)

	key := "testKey"
	_, ok := cache.Get(key)
	if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	cache.Set(key, val)

	retVal, ok := cache.Get(key)
	if !ok {
		t.Fatal("could not retrieve an element we just added")
	}
	if !bytes.Equal(retVal, val) {
		t.Fatal("retrieved a different value than what we put in")
	}

	cache.Delete(key)

	_, ok = cache.Get(key)
	if ok {
		t.Fatal("deleted key still present")
	}

	// memcached has a limit of 250 characters for a key
	chars := make([]byte, 260)
	for i := range chars {
		chars[i] = "x"[0]
	}
	longKey := string(chars)

	_, ok = cache.Get(longKey)
	if ok {
		t.Fatal("retrieved long key before adding it")
	}

	cache.Set(longKey, val)

	retVal, ok = cache.Get(longKey)
	if !ok {
		t.Fatal("could not retrieve an element with a long key")
	}
	if !bytes.Equal(retVal, val) {
		t.Fatal("retrieved a different value than what we put in for an element with a long key")
	}

	cache.Delete(longKey)

	_, ok = cache.Get(longKey)
	if ok {
		t.Fatal("deleted long key still present")
	}
}
