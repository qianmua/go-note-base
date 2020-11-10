package go_error

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
)

/**
 * @author
 * @date 2020/11/10 16:42
 * create by Goland
 * @version 1.0
 */ 
 
// error
func TestErrorM1(t *testing.T) {
	f, err := t1(-1)
	if err !=nil {
		fmt.Println("err " , err)
		os.Exit(-1)
	}
	fmt.Println(f)
}

func t1(f float64) (float64 , error){
	if f < 0 {
		return 0 , errors.New("test this value " + strconv.FormatFloat(f , 'E', -1 , 64)  +  "is le zore")
	}
	// do nothing
	return 0 , nil
}