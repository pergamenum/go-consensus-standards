package reflection

import (
	"fmt"
	"testing"
)

func Test_AutoMap_Field_Value_Value(t *testing.T) {

	type Source struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Info string `automap:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}

func Test_AutoMap_Field_Value_Pointer(t *testing.T) {

	type Source struct {
		Info string `automap:"info"`
	}

	type Target struct {
		Info *string `automap:"info"`
	}

	source := Source{
		Info: "Source.Info",
	}

	target, err := AutoMap[Target](source)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if source.Info != *target.Info {
		fmt.Println("AutoMap failed to set value while returning no error.")
		t.Fail()
	}
}
