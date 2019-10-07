// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2019 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost341194

import (
	"bytes"
	"crypto/rand"
	"hash"
	"testing"
	"testing/quick"

	"go.cypherpunks.ru/gogost/v4/gost28147"
)

func TestHashInterface(t *testing.T) {
	h := New(SboxDefault)
	var _ hash.Hash = h
}

func TestVectors(t *testing.T) {
	h := New(SboxDefault)

	t.Run("empty", func(t *testing.T) {
		if bytes.Compare(h.Sum(nil), []byte{
			0xce, 0x85, 0xb9, 0x9c, 0xc4, 0x67, 0x52, 0xff,
			0xfe, 0xe3, 0x5c, 0xab, 0x9a, 0x7b, 0x02, 0x78,
			0xab, 0xb4, 0xc2, 0xd2, 0x05, 0x5c, 0xff, 0x68,
			0x5a, 0xf4, 0x91, 0x2c, 0x49, 0x49, 0x0f, 0x8d,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("a", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("a"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xd4, 0x2c, 0x53, 0x9e, 0x36, 0x7c, 0x66, 0xe9,
			0xc8, 0x8a, 0x80, 0x1f, 0x66, 0x49, 0x34, 0x9c,
			0x21, 0x87, 0x1b, 0x43, 0x44, 0xc6, 0xa5, 0x73,
			0xf8, 0x49, 0xfd, 0xce, 0x62, 0xf3, 0x14, 0xdd,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("abc", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("abc"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xf3, 0x13, 0x43, 0x48, 0xc4, 0x4f, 0xb1, 0xb2,
			0xa2, 0x77, 0x72, 0x9e, 0x22, 0x85, 0xeb, 0xb5,
			0xcb, 0x5e, 0x0f, 0x29, 0xc9, 0x75, 0xbc, 0x75,
			0x3b, 0x70, 0x49, 0x7c, 0x06, 0xa4, 0xd5, 0x1d,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("message digest", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("message digest"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xad, 0x44, 0x34, 0xec, 0xb1, 0x8f, 0x2c, 0x99,
			0xb6, 0x0c, 0xbe, 0x59, 0xec, 0x3d, 0x24, 0x69,
			0x58, 0x2b, 0x65, 0x27, 0x3f, 0x48, 0xde, 0x72,
			0xdb, 0x2f, 0xde, 0x16, 0xa4, 0x88, 0x9a, 0x4d,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("128U", func(t *testing.T) {
		h.Reset()
		for i := 0; i < 128; i++ {
			h.Write([]byte("U"))
		}
		if bytes.Compare(h.Sum(nil), []byte{
			0x53, 0xa3, 0xa3, 0xed, 0x25, 0x18, 0x0c, 0xef,
			0x0c, 0x1d, 0x85, 0xa0, 0x74, 0x27, 0x3e, 0x55,
			0x1c, 0x25, 0x66, 0x0a, 0x87, 0x06, 0x2a, 0x52,
			0xd9, 0x26, 0xa9, 0xe8, 0xfe, 0x57, 0x33, 0xa4,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("lazy dog", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("The quick brown fox jumps over the lazy dog"))
		if bytes.Compare(h.Sum(nil), []byte{
			0x77, 0xb7, 0xfa, 0x41, 0x0c, 0x9a, 0xc5, 0x8a,
			0x25, 0xf4, 0x9b, 0xca, 0x7d, 0x04, 0x68, 0xc9,
			0x29, 0x65, 0x29, 0x31, 0x5e, 0xac, 0xa7, 0x6b,
			0xd1, 0xa1, 0x0f, 0x37, 0x6d, 0x1f, 0x42, 0x94,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("lazy cog", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("The quick brown fox jumps over the lazy cog"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xa3, 0xeb, 0xc4, 0xda, 0xaa, 0xb7, 0x8b, 0x0b,
			0xe1, 0x31, 0xda, 0xb5, 0x73, 0x7a, 0x7f, 0x67,
			0xe6, 0x02, 0x67, 0x0d, 0x54, 0x35, 0x21, 0x31,
			0x91, 0x50, 0xd2, 0xe1, 0x4e, 0xee, 0xc4, 0x45,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("32", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("This is message, length=32 bytes"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xb1, 0xc4, 0x66, 0xd3, 0x75, 0x19, 0xb8, 0x2e,
			0x83, 0x19, 0x81, 0x9f, 0xf3, 0x25, 0x95, 0xe0,
			0x47, 0xa2, 0x8c, 0xb6, 0xf8, 0x3e, 0xff, 0x1c,
			0x69, 0x16, 0xa8, 0x15, 0xa6, 0x37, 0xff, 0xfa,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("50", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("Suppose the original message has length = 50 bytes"))
		if bytes.Compare(h.Sum(nil), []byte{
			0x47, 0x1a, 0xba, 0x57, 0xa6, 0x0a, 0x77, 0x0d,
			0x3a, 0x76, 0x13, 0x06, 0x35, 0xc1, 0xfb, 0xea,
			0x4e, 0xf1, 0x4d, 0xe5, 0x1f, 0x78, 0xb4, 0xae,
			0x57, 0xdd, 0x89, 0x3b, 0x62, 0xf5, 0x52, 0x08,
		}) != 0 {
			t.FailNow()
		}
	})
}

func TestVectorsCryptoPro(t *testing.T) {
	h := New(&gost28147.SboxIdGostR341194CryptoProParamSet)

	t.Run("empty", func(t *testing.T) {
		if bytes.Compare(h.Sum(nil), []byte{
			0x98, 0x1e, 0x5f, 0x3c, 0xa3, 0x0c, 0x84, 0x14,
			0x87, 0x83, 0x0f, 0x84, 0xfb, 0x43, 0x3e, 0x13,
			0xac, 0x11, 0x01, 0x56, 0x9b, 0x9c, 0x13, 0x58,
			0x4a, 0xc4, 0x83, 0x23, 0x4c, 0xd6, 0x56, 0xc0,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("a", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("a"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xe7, 0x4c, 0x52, 0xdd, 0x28, 0x21, 0x83, 0xbf,
			0x37, 0xaf, 0x00, 0x79, 0xc9, 0xf7, 0x80, 0x55,
			0x71, 0x5a, 0x10, 0x3f, 0x17, 0xe3, 0x13, 0x3c,
			0xef, 0xf1, 0xaa, 0xcf, 0x2f, 0x40, 0x30, 0x11,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("abc", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("abc"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xb2, 0x85, 0x05, 0x6d, 0xbf, 0x18, 0xd7, 0x39,
			0x2d, 0x76, 0x77, 0x36, 0x95, 0x24, 0xdd, 0x14,
			0x74, 0x74, 0x59, 0xed, 0x81, 0x43, 0x99, 0x7e,
			0x16, 0x3b, 0x29, 0x86, 0xf9, 0x2f, 0xd4, 0x2c,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("message digest", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("message digest"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xbc, 0x60, 0x41, 0xdd, 0x2a, 0xa4, 0x01, 0xeb,
			0xfa, 0x6e, 0x98, 0x86, 0x73, 0x41, 0x74, 0xfe,
			0xbd, 0xb4, 0x72, 0x9a, 0xa9, 0x72, 0xd6, 0x0f,
			0x54, 0x9a, 0xc3, 0x9b, 0x29, 0x72, 0x1b, 0xa0,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("lazy dog", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("The quick brown fox jumps over the lazy dog"))
		if bytes.Compare(h.Sum(nil), []byte{
			0x90, 0x04, 0x29, 0x4a, 0x36, 0x1a, 0x50, 0x8c,
			0x58, 0x6f, 0xe5, 0x3d, 0x1f, 0x1b, 0x02, 0x74,
			0x67, 0x65, 0xe7, 0x1b, 0x76, 0x54, 0x72, 0x78,
			0x6e, 0x47, 0x70, 0xd5, 0x65, 0x83, 0x0a, 0x76,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("32", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("This is message, length=32 bytes"))
		if bytes.Compare(h.Sum(nil), []byte{
			0x2c, 0xef, 0xc2, 0xf7, 0xb7, 0xbd, 0xc5, 0x14,
			0xe1, 0x8e, 0xa5, 0x7f, 0xa7, 0x4f, 0xf3, 0x57,
			0xe7, 0xfa, 0x17, 0xd6, 0x52, 0xc7, 0x5f, 0x69,
			0xcb, 0x1b, 0xe7, 0x89, 0x3e, 0xde, 0x48, 0xeb,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("50", func(t *testing.T) {
		h.Reset()
		h.Write([]byte("Suppose the original message has length = 50 bytes"))
		if bytes.Compare(h.Sum(nil), []byte{
			0xc3, 0x73, 0x0c, 0x5c, 0xbc, 0xca, 0xcf, 0x91,
			0x5a, 0xc2, 0x92, 0x67, 0x6f, 0x21, 0xe8, 0xbd,
			0x4e, 0xf7, 0x53, 0x31, 0xd9, 0x40, 0x5e, 0x5f,
			0x1a, 0x61, 0xdc, 0x31, 0x30, 0xa6, 0x50, 0x11,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("128U", func(t *testing.T) {
		h.Reset()
		for i := 0; i < 128; i++ {
			h.Write([]byte{'U'})
		}
		if bytes.Compare(h.Sum(nil), []byte{
			0x1c, 0x4a, 0xc7, 0x61, 0x46, 0x91, 0xbb, 0xf4,
			0x27, 0xfa, 0x23, 0x16, 0x21, 0x6b, 0xe8, 0xf1,
			0x0d, 0x92, 0xed, 0xfd, 0x37, 0xcd, 0x10, 0x27,
			0x51, 0x4c, 0x10, 0x08, 0xf6, 0x49, 0xc4, 0xe8,
		}) != 0 {
			t.FailNow()
		}
	})
}

func TestRandom(t *testing.T) {
	h := New(SboxDefault)
	f := func(data []byte) bool {
		h.Reset()
		h.Write(data)
		d1 := h.Sum(nil)
		h.Reset()
		for _, c := range data {
			h.Write([]byte{c})
		}
		d2 := h.Sum(nil)
		return bytes.Compare(d1, d2) == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkHash(b *testing.B) {
	h := New(SboxDefault)
	src := make([]byte, BlockSize+1)
	rand.Read(src)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Write(src)
		h.Sum(nil)
	}
}
