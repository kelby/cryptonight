// +build !amd64

package cryptonight

func byteAddMul(ret *[2]uint64, x, y uint64) {
	var low, high uint64
	mul128(&low, &high, x, y)
	ret[0] += high
	ret[1] += low
}

func mul128(low, high *uint64, x, y uint64) {
	xhi, yhi := x>>32, y>>32
	xlo, ylo := x&0xffffffff, y&0xffffffff

	hihi := xhi * yhi
	lolo := xlo * ylo
	lohi := xlo * yhi
	hilo := xhi * ylo

	mid := lolo>>32 + lohi&0xffffffff + hilo&0xffffffff
	*low = mid<<32 | (lolo & 0xffffffff)
	*high = hihi + lohi>>32 + hilo>>32 + mid>>32
}
