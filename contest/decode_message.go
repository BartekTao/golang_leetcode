package contest

import "bytes"

func decodeMessage(key string, message string) string {
	mp := make(map[byte]byte)
	c := "abcdefghijklmnopqrstuvwxyz"
	e := make(map[int]byte)
	for i := 0; i < len(c); i++ {
		e[i+1] = c[i]
	}
	for i := 0; i < len(key); i++ {
		v := key[i]
		if _, ok := mp[v]; !ok && v != ' ' {
			mp[v] = e[len(mp)+1]
		}
		if len(mp) == 26 {
			break
		}
	}
	var byte_buf bytes.Buffer
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			byte_buf.WriteByte(' ')
		} else {
			byte_buf.WriteByte(mp[message[i]])
		}
	}
	return byte_buf.String()
}
