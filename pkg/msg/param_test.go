package msg

import "testing"

func TestParam001(t *testing.T) {
	var mapping = `
{
	"name":{
		"description":"姓名",
		"attr":["name", "nick"]
	}, 
	"dayNum":{
		"description":"姓名",
		"attr":["dayNum"]
	}
}
`
	var input = `
{"name":"张三", "dayNum":"2023-09-10"}
`

	res, err := HandleMessageParams(input, mapping)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
