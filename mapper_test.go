package mapper

import (
	"os"
	"reflect"
	"testing"
)

//declare struct
type Src struct {
	A             string   `mapper:"A"`
	B             int      `mapper:"B"`
	C             bool     `mapper:"C"`
	D             []string `mapper:"D"`
	AnonymousTest `mapper:"E"`
}

type Dest struct {
	A             string   `mapper:"A"`
	B             int      `mapper:"B"`
	C             bool     `mapper:"C"`
	D             []string `mapper:"D"`
	AnonymousTest `mapper:"E"`
}

type AnonymousTest struct {
	AA int
	bb int
}

var Test_Src Src
var Test_Dest Dest

//设置及拆卸
//setup and teardown
func TestMain(m *testing.M) {
	Test_Src = Src{
		A:             "test",
		B:             123,
		C:             true,
		D:             []string{"test1", "test2", "test3"},
		AnonymousTest: AnonymousTest{AA: 1, bb: 2},
	}
	retCode := m.Run() //run test
	os.Exit(retCode)
}

//逻辑行为测试
//act test
func TestStructMapByFieldName(t *testing.T) {
	//arrange
	expect_dest := Dest{
		A:             Test_Src.A,
		B:             Test_Src.B,
		C:             Test_Src.C,
		D:             Test_Src.D,
		AnonymousTest: Test_Src.AnonymousTest,
	}

	//act
	err := StructMapByFieldName(&Test_Src, &Test_Dest)
	if err != nil {
		t.Errorf("type error:%s", err)
	}

	//assert
	if !reflect.DeepEqual(expect_dest, Test_Dest) {
		t.Errorf("expect:%v,actual:%v", expect_dest, Test_Dest)
	}
}

func TestStructMapByFieldTag(t *testing.T) {
	//arrange
	expect_dest := Dest{
		A:             Test_Src.A,
		B:             Test_Src.B,
		C:             Test_Src.C,
		D:             Test_Src.D,
		AnonymousTest: Test_Src.AnonymousTest,
	}

	//act
	err := StructMapByFieldTag(&Test_Src, &Test_Dest)
	if err != nil {
		t.Errorf("type error:%s", err)
	}

	//assert
	if !reflect.DeepEqual(expect_dest, Test_Dest) {
		t.Errorf("expect:%v,actual:%v", expect_dest, Test_Dest)
	}
}

//性能测试
//benchmark test
func BenchmarkStructMapByFieldName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StructMapByFieldName(&Test_Src, &Test_Dest)
	}
}

func BenchmarkStructMapByFieldTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StructMapByFieldTag(&Test_Src, &Test_Dest)
	}
}
