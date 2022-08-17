// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2022 Sergey Matveev <stargrave@stargrave.org>
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
	"hash"
	"testing"

	"github.com/thefish/gogost/gost28147"
	"golang.org/x/crypto/pbkdf2"
)

func PBKDF2Hash() hash.Hash {
	return New(&gost28147.SboxIdGostR341194CryptoProParamSet)
}

// Test vectors for PBKDF2 taken from
// http://tc26.ru/methods/containers_v1/Addition_to_PKCS5_v1_0.pdf test vectors
func TestPBKDF2Vectors(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		if bytes.Compare(pbkdf2.Key(
			[]byte("password"),
			[]byte("salt"),
			1,
			32,
			PBKDF2Hash,
		), []byte{
			0x73, 0x14, 0xe7, 0xc0, 0x4f, 0xb2, 0xe6, 0x62,
			0xc5, 0x43, 0x67, 0x42, 0x53, 0xf6, 0x8b, 0xd0,
			0xb7, 0x34, 0x45, 0xd0, 0x7f, 0x24, 0x1b, 0xed,
			0x87, 0x28, 0x82, 0xda, 0x21, 0x66, 0x2d, 0x58,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("2", func(t *testing.T) {
		if bytes.Compare(pbkdf2.Key(
			[]byte("password"),
			[]byte("salt"),
			2,
			32,
			PBKDF2Hash,
		), []byte{
			0x99, 0x0d, 0xfa, 0x2b, 0xd9, 0x65, 0x63, 0x9b,
			0xa4, 0x8b, 0x07, 0xb7, 0x92, 0x77, 0x5d, 0xf7,
			0x9f, 0x2d, 0xb3, 0x4f, 0xef, 0x25, 0xf2, 0x74,
			0x37, 0x88, 0x72, 0xfe, 0xd7, 0xed, 0x1b, 0xb3,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("4096", func(t *testing.T) {
		if bytes.Compare(pbkdf2.Key(
			[]byte("password"),
			[]byte("salt"),
			4096,
			32,
			PBKDF2Hash,
		), []byte{
			0x1f, 0x18, 0x29, 0xa9, 0x4b, 0xdf, 0xf5, 0xbe,
			0x10, 0xd0, 0xae, 0xb3, 0x6a, 0xf4, 0x98, 0xe7,
			0xa9, 0x74, 0x67, 0xf3, 0xb3, 0x11, 0x16, 0xa5,
			0xa7, 0xc1, 0xaf, 0xff, 0x9d, 0xea, 0xda, 0xfe,
		}) != 0 {
			t.FailNow()
		}
	})

	/*
		t.Run("16777216", func(t *testing.T) {
			if bytes.Compare(pbkdf2.Key(
				[]byte("password"),
				[]byte("salt"),
				16777216,
				32,
				PBKDF2Hash,
			), []byte{
				0xa5, 0x7a, 0xe5, 0xa6, 0x08, 0x83, 0x96, 0xd1,
				0x20, 0x85, 0x0c, 0x5c, 0x09, 0xde, 0x0a, 0x52,
				0x51, 0x00, 0x93, 0x8a, 0x59, 0xb1, 0xb5, 0xc3,
				0xf7, 0x81, 0x09, 0x10, 0xd0, 0x5f, 0xcd, 0x97,
			}) != 0 {
				t.FailNow()
			}
		})
	*/

	t.Run("many", func(t *testing.T) {
		if bytes.Compare(pbkdf2.Key(
			[]byte("passwordPASSWORDpassword"),
			[]byte("saltSALTsaltSALTsaltSALTsaltSALTsalt"),
			4096,
			40,
			PBKDF2Hash,
		), []byte{
			0x78, 0x83, 0x58, 0xc6, 0x9c, 0xb2, 0xdb, 0xe2,
			0x51, 0xa7, 0xbb, 0x17, 0xd5, 0xf4, 0x24, 0x1f,
			0x26, 0x5a, 0x79, 0x2a, 0x35, 0xbe, 0xcd, 0xe8,
			0xd5, 0x6f, 0x32, 0x6b, 0x49, 0xc8, 0x50, 0x47,
			0xb7, 0x63, 0x8a, 0xcb, 0x47, 0x64, 0xb1, 0xfd,
		}) != 0 {
			t.FailNow()
		}
	})

	t.Run("zero byte", func(t *testing.T) {
		if bytes.Compare(pbkdf2.Key(
			[]byte("pass\x00word"),
			[]byte("sa\x00lt"),
			4096,
			20,
			PBKDF2Hash,
		), []byte{
			0x43, 0xe0, 0x6c, 0x55, 0x90, 0xb0, 0x8c, 0x02,
			0x25, 0x24, 0x23, 0x73, 0x12, 0x7e, 0xdf, 0x9c,
			0x8e, 0x9c, 0x32, 0x91,
		}) != 0 {
			t.FailNow()
		}
	})
}
