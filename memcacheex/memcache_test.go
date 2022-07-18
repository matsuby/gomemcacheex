package memcacheex

import (
	"errors"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/mock/gomock"
)

var (
	testErr  = errors.New("test")
	testKey  = "testKey"
	testItem = &memcache.Item{
		Key:   testKey,
		Value: []byte("test"),
	}
	testDelta uint64 = 1
)

func TestClientWrapper(t *testing.T) {
	mc := NewMockClient(gomock.NewController(t))
	cw := NewClientWrapper(mc)

	t.Run("FlushAll", func(t *testing.T) {
		mc.EXPECT().FlushAll().Return(testErr)
		bn := "FlushAll:Before"
		an := "FlushAll:After"

		cw.Callback().FlushAll().Before().Register(bn, func(args, results []any) {})
		cw.Callback().FlushAll().After().Register(an, func(args, results []any) {})
		cw.FlushAll()

		if l := len(*cw.Callback().FlushAll().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().FlushAll().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().FlushAll().Before().Unregister(bn)
		cw.Callback().FlushAll().After().Unregister(an)

		if l := len(*cw.Callback().FlushAll().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().FlushAll().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Get", func(t *testing.T) {
		mc.EXPECT().Get(gomock.Eq(testKey)).Return(nil, testErr)
		bn := "Get:Before"
		an := "Get:After"

		cw.Callback().Get().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Get().After().Register(an, func(args, results []any) {})
		cw.Get(testKey)

		if l := len(*cw.Callback().Get().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Get().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Get().Before().Unregister(bn)
		cw.Callback().Get().After().Unregister(an)

		if l := len(*cw.Callback().Get().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Get().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Touch", func(t *testing.T) {
		testSec := int32(1)
		mc.EXPECT().Touch(gomock.Eq(testKey), gomock.Eq(testSec)).Return(testErr)
		bn := "Touch:Before"
		an := "Touch:After"

		cw.Callback().Touch().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Touch().After().Register(an, func(args, results []any) {})
		cw.Touch(testKey, testSec)

		if l := len(*cw.Callback().Touch().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Touch().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Touch().Before().Unregister(bn)
		cw.Callback().Touch().After().Unregister(an)

		if l := len(*cw.Callback().Touch().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Touch().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("GetMulti", func(t *testing.T) {
		testKeys := []string{testKey}
		mc.EXPECT().GetMulti(gomock.Eq(testKeys)).Return(nil, testErr)
		bn := "GetMulti:Before"
		an := "GetMulti:After"

		cw.Callback().GetMulti().Before().Register(bn, func(args, results []any) {})
		cw.Callback().GetMulti().After().Register(an, func(args, results []any) {})
		cw.GetMulti(testKeys)

		if l := len(*cw.Callback().GetMulti().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().GetMulti().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().GetMulti().Before().Unregister(bn)
		cw.Callback().GetMulti().After().Unregister(an)

		if l := len(*cw.Callback().GetMulti().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().GetMulti().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Set", func(t *testing.T) {
		mc.EXPECT().Set(gomock.Eq(testItem)).Return(testErr)
		bn := "Set:Before"
		an := "Set:After"

		cw.Callback().Set().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Set().After().Register(an, func(args, results []any) {})
		cw.Set(testItem)

		if l := len(*cw.Callback().Set().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Set().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Set().Before().Unregister(bn)
		cw.Callback().Set().After().Unregister(an)

		if l := len(*cw.Callback().Set().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Set().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Add", func(t *testing.T) {
		mc.EXPECT().Add(gomock.Eq(testItem)).Return(testErr)
		bn := "Add:Before"
		an := "Add:After"

		cw.Callback().Add().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Add().After().Register(an, func(args, results []any) {})
		cw.Add(testItem)

		if l := len(*cw.Callback().Add().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Add().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Add().Before().Unregister(bn)
		cw.Callback().Add().After().Unregister(an)

		if l := len(*cw.Callback().Add().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Add().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Replace", func(t *testing.T) {
		mc.EXPECT().Replace(gomock.Eq(testItem)).Return(testErr)
		bn := "Replace:Before"
		an := "Replace:After"

		cw.Callback().Replace().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Replace().After().Register(an, func(args, results []any) {})
		cw.Replace(testItem)

		if l := len(*cw.Callback().Replace().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Replace().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Replace().Before().Unregister(bn)
		cw.Callback().Replace().After().Unregister(an)

		if l := len(*cw.Callback().Replace().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Replace().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("CompareAndSwap", func(t *testing.T) {
		mc.EXPECT().CompareAndSwap(gomock.Eq(testItem)).Return(testErr)
		bn := "CompareAndSwap:Before"
		an := "CompareAndSwap:After"

		cw.Callback().CompareAndSwap().Before().Register(bn, func(args, results []any) {})
		cw.Callback().CompareAndSwap().After().Register(an, func(args, results []any) {})
		cw.CompareAndSwap(testItem)

		if l := len(*cw.Callback().CompareAndSwap().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().CompareAndSwap().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().CompareAndSwap().Before().Unregister(bn)
		cw.Callback().CompareAndSwap().After().Unregister(an)

		if l := len(*cw.Callback().CompareAndSwap().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().CompareAndSwap().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		mc.EXPECT().Delete(gomock.Eq(testKey)).Return(testErr)
		bn := "Delete:Before"
		an := "Delete:After"

		cw.Callback().Delete().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Delete().After().Register(an, func(args, results []any) {})
		cw.Delete(testKey)

		if l := len(*cw.Callback().Delete().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Delete().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Delete().Before().Unregister(bn)
		cw.Callback().Delete().After().Unregister(an)

		if l := len(*cw.Callback().Delete().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Delete().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("DeleteAll", func(t *testing.T) {
		mc.EXPECT().DeleteAll().Return(testErr)
		bn := "DeleteAll:Before"
		an := "DeleteAll:After"

		cw.Callback().DeleteAll().Before().Register(bn, func(args, results []any) {})
		cw.Callback().DeleteAll().After().Register(an, func(args, results []any) {})
		cw.DeleteAll()

		if l := len(*cw.Callback().DeleteAll().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().DeleteAll().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().DeleteAll().Before().Unregister(bn)
		cw.Callback().DeleteAll().After().Unregister(an)

		if l := len(*cw.Callback().DeleteAll().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().DeleteAll().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Ping", func(t *testing.T) {
		mc.EXPECT().Ping().Return(testErr)
		bn := "Ping:Before"
		an := "Ping:After"

		cw.Callback().Ping().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Ping().After().Register(an, func(args, results []any) {})
		cw.Ping()

		if l := len(*cw.Callback().Ping().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Ping().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Ping().Before().Unregister(bn)
		cw.Callback().Ping().After().Unregister(an)

		if l := len(*cw.Callback().Ping().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Ping().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Increment", func(t *testing.T) {
		mc.EXPECT().Increment(gomock.Eq(testKey), gomock.Eq(testDelta)).Return(10+testDelta, testErr)
		bn := "Increment:Before"
		an := "Increment:After"

		cw.Callback().Increment().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Increment().After().Register(an, func(args, results []any) {})
		cw.Increment(testKey, testDelta)

		if l := len(*cw.Callback().Increment().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Increment().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Increment().Before().Unregister(bn)
		cw.Callback().Increment().After().Unregister(an)

		if l := len(*cw.Callback().Increment().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Increment().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})

	t.Run("Decrement", func(t *testing.T) {
		mc.EXPECT().Decrement(gomock.Eq(testKey), gomock.Eq(testDelta)).Return(10+testDelta, testErr)
		bn := "Decrement:Before"
		an := "Decrement:After"

		cw.Callback().Decrement().Before().Register(bn, func(args, results []any) {})
		cw.Callback().Decrement().After().Register(an, func(args, results []any) {})
		cw.Decrement(testKey, testDelta)

		if l := len(*cw.Callback().Decrement().Before()); l != 1 {
			t.Errorf("%s was not registered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Decrement().After()); l != 1 {
			t.Errorf("%s was not registered callback: %d", an, l)
		}

		cw.Callback().Decrement().Before().Unregister(bn)
		cw.Callback().Decrement().After().Unregister(an)

		if l := len(*cw.Callback().Decrement().Before()); l != 0 {
			t.Errorf("%s was not unregistered callback: %d", bn, l)
		}
		if l := len(*cw.Callback().Decrement().After()); l != 0 {
			t.Errorf("%s was not unregistered callback, %d", an, l)
		}
	})
}
