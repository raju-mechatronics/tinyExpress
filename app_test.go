package te

import (
	"fmt"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	app := App()

	if app == nil {
		t.Error("Expected app to be created")
	}

	err := app.Listen()
	fmt.Println("err=>", err)

	time.Sleep(5000 * time.Second)
}
