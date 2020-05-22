package main

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		res:=0
		for i := 1; i <=n; i = i +4 {
			res +=read4(buf)
			if len(buf)>=4{
				buf=buf[4:]
			}
		}
		if n< res {
			return n
		}else {
			return res
		}
	}
}
