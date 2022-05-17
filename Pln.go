package mine

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func Pln(a ...any) {
	_, file, line, _ := runtime.Caller(1)
	marshal, _ := json.Marshal(a)
	fmt.Println(Now(), file, line, ":", string(marshal))
}
