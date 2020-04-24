// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gflock_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gflock"
)

func Test_GFlock_Base(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		fileName := "test"
		lock := gflock.New(fileName)
		t.Assert(lock.Path(), filepath.Join(os.TempDir(), "gflock", fileName))
		t.Assert(lock.IsLocked(), false)
		lock.Lock()
		t.Assert(lock.IsLocked(), true)
		lock.Unlock()
		t.Assert(lock.IsLocked(), false)
	})

	gtest.C(t, func(t *gtest.T) {
		fileName := "test"
		lock := gflock.New(fileName)
		t.Assert(lock.Path(), filepath.Join(os.TempDir(), "gflock", fileName))
		t.Assert(lock.IsRLocked(), false)
		lock.RLock()
		t.Assert(lock.IsRLocked(), true)
		lock.RUnlock()
		t.Assert(lock.IsRLocked(), false)
	})
}

func Test_GFlock_Lock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			fileName = "testLock"
			array    = garray.New(true)
			lock     = gflock.New(fileName)
			lock2    = gflock.New(fileName)
		)
		go func() {
			lock.Lock()
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			lock.Unlock()
		}()

		go func() {
			time.Sleep(100 * time.Millisecond)
			lock2.Lock()
			array.Append(1)
			lock2.Unlock()
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_GFlock_RLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			fileName = "testRLock"
			array    = garray.New(true)
			lock     = gflock.New(fileName)
			lock2    = gflock.New(fileName)
		)
		go func() {
			lock.RLock()
			array.Append(1)
			time.Sleep(400 * time.Millisecond)
			lock.RUnlock()
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			lock2.RLock()
			array.Append(1)
			lock2.RUnlock()
		}()

		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_GFlock_TryLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			fileName = "testTryLock"
			array    = garray.New(true)
			lock     = gflock.New(fileName)
			lock2    = gflock.New(fileName)
		)
		go func() {
			lock.TryLock()
			array.Append(1)
			time.Sleep(200 * time.Millisecond)
			lock.Unlock()
		}()

		go func() {
			time.Sleep(100 * time.Millisecond)
			if lock2.TryLock() {
				array.Append(1)
				lock2.Unlock()
			}
		}()

		go func() {
			time.Sleep(300 * time.Millisecond)
			if lock2.TryLock() {
				array.Append(1)
				lock2.Unlock()
			}
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(200 * time.Millisecond)
		t.Assert(array.Len(), 2)
	})
}

func Test_GFlock_TryRLock(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			fileName = "testTryRLock"
			array    = garray.New(true)
			lock     = gflock.New(fileName)
			lock2    = gflock.New(fileName)
		)
		go func() {
			lock.TryRLock()
			array.Append(1)
			time.Sleep(300 * time.Millisecond)
			lock.Unlock()
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			if lock2.TryRLock() {
				array.Append(1)
				lock2.Unlock()
			}
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			if lock2.TryRLock() {
				array.Append(1)
				lock2.Unlock()
			}
		}()

		go func() {
			time.Sleep(200 * time.Millisecond)
			if lock2.TryRLock() {
				array.Append(1)
				lock2.Unlock()
			}
		}()
		time.Sleep(100 * time.Millisecond)
		t.Assert(array.Len(), 1)
		time.Sleep(300 * time.Millisecond)
		t.Assert(array.Len(), 4)
	})
}
