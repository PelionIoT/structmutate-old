package structmutate

import (
    "reflect"
    "errors"
    "fmt"
)


type logOut func (format string, a ...interface{})

var errOut = func (format string, a ...interface{}) {
    s := fmt.Sprintf(format, a...)
    fmt.Printf("[error-structmutate]  %s\n", s)
}

var dbgOut = func (format string, a ...interface{}) {

}


func SetupUtilsLogs (debugout logOut, errout logOut) {
    if debugout != nil {
        dbgOut = debugout
    }
    if errout != nil {
        errOut = errout
    }
}

// merges in any fields in 'b' which are not empty, where
// empty is considered an empty string, a nil pointer, or a 0 number
// onto 'a' - if 'a' has the same field name and type
// Takes two pointers to structs
func MergeInStruct(a interface{}, b interface{}) (err error) {
    kind := reflect.ValueOf(a).Kind()
    kind_b := reflect.ValueOf(b).Kind()

    dbgOut("debugging is on! %+v",kind)

    if kind != reflect.Ptr {
        return errors.New("Not a pointer (a).")
    } 
    if kind_b != reflect.Ptr {
        return errors.New("Not a pointer (b).")
    }
    // if kind != reflect.Struct {
    //     return errors.New("Not a struct (a).")
    // }
    // if kind_b != reflect.Struct {
    //     return errors.New("Not a struct (b).")
    // }

    structType := reflect.ValueOf(a).Elem().Type()
    structTypeB := reflect.ValueOf(b).Elem().Type()

    if structType.Kind() != reflect.Struct {
        return errors.New("Not a pointer (a) to struct.")
    }
    if structTypeB.Kind() != reflect.Struct {
        return errors.New("Not a pointer (b) to struct.")
    }

    structVal := reflect.ValueOf(a).Elem()
    structValType := structVal.Type()
    // structTypeB := reflect.TypeOf(b)
    structValB := reflect.ValueOf(b).Elem()
    // structValTypeB := structValB.Type()

    for i := 0; i < structType.NumField(); i++ {
//        typeName := reflectType.Field(i).Name
        if structVal.Field(i).IsValid() {
            fieldname := structValType.Field(i).Name
            fieldvala := structVal.FieldByName(fieldname)
            fieldvalb := structValB.FieldByName(fieldname)
            if !fieldvala.CanSet() {
                errOut("MergeInStruct(): Field <%s> of struct can not be set, skipping.\n",fieldname)
                continue
            }

            if fieldvalb.Kind() == reflect.Bool {
                fieldvala.SetBool(fieldvalb.Bool())
                dbgOut("copied field <%s> (bool)\n",fieldname)                    
                continue
            }
            if fieldvalb.IsValid() {
                    if fieldvalb.Type().Kind() == fieldvala.Type().Kind() {
                        fieldvala.Set(fieldvalb)
                        dbgOut("copied field <%s>\n",fieldname)                    
                    } else {
                        dbgOut("skipping field <%s> - is zero/not valid\n",fieldname)
                    }                        
            } else {
                dbgOut("fields <%s> not of same Kind",fieldname)
            }

        } else {
            dbgOut("field <%s> in recieving struct is not valid",structValType.Field(i).Name)
        }
    }

    return
}
