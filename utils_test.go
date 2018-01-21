package structmutate

import (
    "testing"
    "fmt"
    "log"
)


var debugout = func (format string, a ...interface{}) {
    s := fmt.Sprintf(format, a...)
    fmt.Printf("[debug-typeutils]  %s\n", s)
}
// func TestMain(m *testing.M) {
//     m.Run()
// }

func TestMergeIn(t *testing.T) {
    SetupUtilsLogs(debugout,nil)

    type foo struct {
        One int
        Two string
        Three []byte
        Abool bool
        Abool2 bool
        Abool3 bool
    }


    type bar struct {
        One int
        Notthere int
        Three []byte
        Abool bool
        Abool2 bool
        Abool3 bool
    }

    foo1 := &foo{
        One: 1,
        Two: "foo!!",
        Three: []byte{0,1,2},
        Abool: true,
        Abool2: true,
        Abool3: true,
    }

    bar1 := &bar{
        One: 101,
        Notthere: 2,
        Three: []byte{99,100},
        Abool: true,
        Abool3: false,
    }

    err := MergeInStruct(foo1,bar1)
    if err != nil {
        log.Fatal("Error: ",err)
    }
    if foo1.One != 101 || len(foo1.Three) != 2 || foo1.Two != "foo!!" || foo1.Abool != true || foo1.Abool2 != false || foo1.Abool3 != false {
        log.Fatalf("Did not get expect results. foo1 = %+v\n",foo1)
    }


    fmt.Printf("foo1: %+v\n",foo1)
}
