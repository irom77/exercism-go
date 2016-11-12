package secret

//borrowed from https://github.com/ThomasZumsteg/exercism-go/blob/master/secret-handshake/secret_handshake.go
const testVersion = 1
var secrets = []string{"wink", "double blink", "close your eyes", "jump"}
func Handshake (code uint) []string {
	var result []string

	for s, secret := range secrets {
		if (1<<uint(s))&code > 0 {
			result = append(result, secret)
		}
	}
	if (code & 16) > 0 {
		reverse(result)
	}
	return result
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}