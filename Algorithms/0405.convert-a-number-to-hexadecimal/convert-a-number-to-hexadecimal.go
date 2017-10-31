package Problem0405

var h = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
}

func toHex(num int) string {
	if num < 0 {
		return toHex(num + 1<<32)
	}

	if num == 0 {
		return "0"
	}

	res := ""

	for num > 0 {
		res = h[num%16] + res
		num /= 16
	}

	return res
}
