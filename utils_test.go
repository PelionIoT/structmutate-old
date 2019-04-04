package structmutate

// MIT License
//
// Copyright (c) 2018 WigWag Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
