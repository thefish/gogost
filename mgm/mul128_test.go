// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2024 Sergey Matveev <stargrave@stargrave.org>
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

package mgm

import (
	"crypto/rand"
	"testing"

	"github.com/thefish/gogost/gost3412128"
)

func BenchmarkMul128(b *testing.B) {
	x := make([]byte, gost3412128.BlockSize)
	y := make([]byte, gost3412128.BlockSize)
	rand.Read(x)
	rand.Read(y)
	mul := newMul128()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul.Mul(x, y)
	}
}
