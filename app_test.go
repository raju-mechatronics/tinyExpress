package te

import (
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	app := App()

	if app == nil {
		t.Error("Expected app to be created")
	}

	err := app.Listen()
	fmt.Println("err=>", err)

}
