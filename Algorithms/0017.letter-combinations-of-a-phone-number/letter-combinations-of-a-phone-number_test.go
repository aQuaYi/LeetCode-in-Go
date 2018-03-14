package problem0017

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []string
}

func Test_Problem0017(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{""},
			ans{nil},
		},

		question{
			para{"23"},
			ans{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},

		question{
			para{"34567"},
			ans{[]string{"dgjmp", "dgjmq", "dgjmr", "dgjms", "dgjnp", "dgjnq", "dgjnr", "dgjns", "dgjop", "dgjoq", "dgjor", "dgjos", "dgkmp", "dgkmq", "dgkmr", "dgkms", "dgknp", "dgknq", "dgknr", "dgkns", "dgkop", "dgkoq", "dgkor", "dgkos", "dglmp", "dglmq", "dglmr", "dglms", "dglnp", "dglnq", "dglnr", "dglns", "dglop", "dgloq", "dglor", "dglos", "dhjmp", "dhjmq", "dhjmr", "dhjms", "dhjnp", "dhjnq", "dhjnr", "dhjns", "dhjop", "dhjoq", "dhjor", "dhjos", "dhkmp", "dhkmq", "dhkmr", "dhkms", "dhknp", "dhknq", "dhknr", "dhkns", "dhkop", "dhkoq", "dhkor", "dhkos", "dhlmp", "dhlmq", "dhlmr", "dhlms", "dhlnp", "dhlnq", "dhlnr", "dhlns", "dhlop", "dhloq", "dhlor", "dhlos", "dijmp", "dijmq", "dijmr", "dijms", "dijnp", "dijnq", "dijnr", "dijns", "dijop", "dijoq", "dijor", "dijos", "dikmp", "dikmq", "dikmr", "dikms", "diknp", "diknq", "diknr", "dikns", "dikop", "dikoq", "dikor", "dikos", "dilmp", "dilmq", "dilmr", "dilms", "dilnp", "dilnq", "dilnr", "dilns", "dilop", "diloq", "dilor", "dilos", "egjmp", "egjmq", "egjmr", "egjms", "egjnp", "egjnq", "egjnr", "egjns", "egjop", "egjoq", "egjor", "egjos", "egkmp", "egkmq", "egkmr", "egkms", "egknp", "egknq", "egknr", "egkns", "egkop", "egkoq", "egkor", "egkos", "eglmp", "eglmq", "eglmr", "eglms", "eglnp", "eglnq", "eglnr", "eglns", "eglop", "egloq", "eglor", "eglos", "ehjmp", "ehjmq", "ehjmr", "ehjms", "ehjnp", "ehjnq", "ehjnr", "ehjns", "ehjop", "ehjoq", "ehjor", "ehjos", "ehkmp", "ehkmq", "ehkmr", "ehkms", "ehknp", "ehknq", "ehknr", "ehkns", "ehkop", "ehkoq", "ehkor", "ehkos", "ehlmp", "ehlmq", "ehlmr", "ehlms", "ehlnp", "ehlnq", "ehlnr", "ehlns", "ehlop", "ehloq", "ehlor", "ehlos", "eijmp", "eijmq", "eijmr", "eijms", "eijnp", "eijnq", "eijnr", "eijns", "eijop", "eijoq", "eijor", "eijos", "eikmp", "eikmq", "eikmr", "eikms", "eiknp", "eiknq", "eiknr", "eikns", "eikop", "eikoq", "eikor", "eikos", "eilmp", "eilmq", "eilmr", "eilms", "eilnp", "eilnq", "eilnr", "eilns", "eilop", "eiloq", "eilor", "eilos", "fgjmp", "fgjmq", "fgjmr", "fgjms", "fgjnp", "fgjnq", "fgjnr", "fgjns", "fgjop", "fgjoq", "fgjor", "fgjos", "fgkmp", "fgkmq", "fgkmr", "fgkms", "fgknp", "fgknq", "fgknr", "fgkns", "fgkop", "fgkoq", "fgkor", "fgkos", "fglmp", "fglmq", "fglmr", "fglms", "fglnp", "fglnq", "fglnr", "fglns", "fglop", "fgloq", "fglor", "fglos", "fhjmp", "fhjmq", "fhjmr", "fhjms", "fhjnp", "fhjnq", "fhjnr", "fhjns", "fhjop", "fhjoq", "fhjor", "fhjos", "fhkmp", "fhkmq", "fhkmr", "fhkms", "fhknp", "fhknq", "fhknr", "fhkns", "fhkop", "fhkoq", "fhkor", "fhkos", "fhlmp", "fhlmq", "fhlmr", "fhlms", "fhlnp", "fhlnq", "fhlnr", "fhlns", "fhlop", "fhloq", "fhlor", "fhlos", "fijmp", "fijmq", "fijmr", "fijms", "fijnp", "fijnq", "fijnr", "fijns", "fijop", "fijoq", "fijor", "fijos", "fikmp", "fikmq", "fikmr", "fikms", "fiknp", "fiknq", "fiknr", "fikns", "fikop", "fikoq", "fikor", "fikos", "filmp", "filmq", "filmr", "films", "filnp", "filnq", "filnr", "filns", "filop", "filoq", "filor", "filos"}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, letterCombinations(p.one), "输入:%v", p)
	}
}
