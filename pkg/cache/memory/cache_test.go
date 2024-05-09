package memcache

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func BenchmarkMemoryCacheSetWithCleanUp(b *testing.B) {
	const cleanupInterval = time.Millisecond * 300
	const defaultExpiration = time.Millisecond * 500

	cache := New[int](context.Background(), 0, 0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cacheWithCleanUp := New[int](ctx, defaultExpiration, cleanupInterval)

	runtime.GC()
	b.ResetTimer()
	b.Run("Set without clean up", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				value := rand.Int()
				key := strconv.Itoa(value)
				cache.Set(key, value, 0)
			}
		})
	})

	b.Run("Set with clean up", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				value := rand.Int()
				key := strconv.Itoa(value)
				cacheWithCleanUp.Set(
					key,
					value,
					time.Millisecond*time.Duration(rand.Intn(int(defaultExpiration))),
				)
			}
		})
	})
}

func BenchmarkMemoryCacheGet(b *testing.B) {
	const N = 1000000

	testData := make([]int, 0, N)
	for i := 0; i < N; i++ {
		testData = append(testData, rand.Int())
	}

	cache := New[int](context.Background(), 0, 0)
	for i, value := range testData {
		cache.Set(strconv.Itoa(i), value, 0)
	}

	runtime.GC()
	b.ResetTimer()

	b.Run("Get", func(sb *testing.B) {
		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := rand.Intn(N)
				want := testData[i]

				got, _ := cache.Get(strconv.Itoa(i))
				assert.Equal(sb, want, got, fmt.Sprintf("want %d, got %d", want, got))
			}
		})

		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := rand.Intn(N)
				want := testData[i]

				got, _ := cache.Get(strconv.Itoa(i))
				assert.Equal(sb, want, got, fmt.Sprintf("want %d, got %d", want, got))
			}
		})

		sb.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := rand.Intn(N)
				want := testData[i]

				got, _ := cache.Get(strconv.Itoa(i))
				assert.Equal(sb, want, got, fmt.Sprintf("want %d, got %d", want, got))
			}
		})
	})
}

func TestBasics(t *testing.T) {
	type testValue struct {
		key        string
		value      string
		expiration time.Duration
	}

	cleanupInterval := time.Millisecond * 15
	testData := []testValue{
		{
			key:        "1",
			value:      "1",
			expiration: time.Millisecond * 50,
		},
		{
			key:        "2",
			value:      "2",
			expiration: time.Millisecond * 50,
		},
		{
			key:        "3",
			value:      "3",
			expiration: 0,
		},
		{
			key:        "4",
			value:      "4",
			expiration: 0,
		},
		{
			key:        "5",
			value:      "5",
			expiration: 0,
		},
	}

	now := time.Now()
	cache := New[string](context.Background(), 0, cleanupInterval)
	for _, value := range testData {
		cache.Set(value.key, value.value, value.expiration)
	}

	t.Run("check cache storage data", func(t *testing.T) {
		for _, testValue := range testData {
			value, err := cache.Get(testValue.key)
			assert.Equal(t, nil, err, "should not be error to get not expired data")
			assert.Equal(
				t,
				testValue.value,
				value,
				fmt.Sprintf("want %s, got %s", testValue.value, value),
			)
		}
	})

	time.Sleep(time.Millisecond * 100)

	t.Run("check expired data", func(t *testing.T) {
		for _, testValue := range testData {
			value, err := cache.Get(testValue.key)
			if testValue.expiration > 0 &&
				now.Add(testValue.expiration).UnixNano() < time.Now().UnixNano() {
				if assert.Error(t, err, "should be ErrCacheMissed error") {
					assert.Equal(
						t,
						ErrCacheMissed,
						err,
						fmt.Sprintf("want %s, got %s", ErrCacheMissed.Error(), err.Error()),
					)
				}
			} else {
				assert.Equal(
					t,
					testValue.value,
					value,
					fmt.Sprintf("want %s, got %s", testValue.value, value),
				)
			}
		}
	})

	t.Run("check expired data", func(t *testing.T) {
		for _, testValue := range testData {
			if !(testValue.expiration > 0 &&
				now.Add(testValue.expiration).UnixNano() < time.Now().UnixNano()) {

				cache.Delete(testValue.key)
				_, err := cache.Get(testValue.key)
				if assert.Error(t, err, "should be ErrCacheMissed error") {
					assert.Equal(
						t,
						ErrCacheMissed,
						err,
						fmt.Sprintf("want %s, got %s", ErrCacheMissed.Error(), err.Error()),
					)
				}
			}
		}
	})
}

func Test_expiredItems(t *testing.T) {
	testData := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}

	cache := New[int](context.Background(), time.Millisecond*50, 0)
	for key, value := range testData {
		cache.Set(key, value, 0)
	}

	t.Run("check len of cache storage", func(t *testing.T) {
		want := len(testData)
		got := len(cache.storage)

		assert.Equal(
			t,
			want,
			got,
			fmt.Sprintf("want %d, got %d", want, got),
		)
	})

	time.Sleep(time.Millisecond * 75)

	t.Run("check expired keys", func(t *testing.T) {
		want := maps.Keys(cache.storage)
		got := cache.expiredItems()

		sort.Strings(want)
		sort.Strings(got)

		assert.Equal(
			t,
			want,
			got,
			fmt.Sprintf("want %s, got %s", want, got),
		)
	})
}

func Test_cleanItems(t *testing.T) {
	testData := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}

	cache := New[int](context.Background(), 0, 0)
	for key, value := range testData {
		cache.Set(key, value, 0)
	}

	t.Run("check len of cache storage", func(t *testing.T) {
		want := len(testData)
		got := len(cache.storage)

		assert.Equal(
			t,
			want,
			got,
			fmt.Sprintf("want %d, got %d", want, got),
		)
	})

	t.Run("check len of cache storage after cleaning", func(t *testing.T) {
		cache.cleanItems(maps.Keys(testData))

		want := 0
		got := len(cache.storage)

		assert.Equal(
			t,
			want,
			got,
			fmt.Sprintf("want %d, got %d", want, got),
		)
	})
}
