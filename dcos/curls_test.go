package dcos

import (
	"fmt"
	"testing"
)

func TestCurlIt(t *testing.T) {

	location := "http://gdmst001v.gsil.rri-usa.org/" //os.Args[1]
	username := "ggaspar"
	password := "0J151nn0c3nt!"

	var entry InjectorEntry
	entry.keyName = "A_KEYFILE_NAME"
	entry.value = "bazz"
	entry.path = "/gsiltest/autofix"
	entry.secret = false
	entry.function = "DELETE"

	token := Login(location, username, password)
	fmt.Printf("\n%v\n\n", string(token))

	respURL := IntentManagerURL(location, string(token))
	fmt.Printf("\n%v\n\n", string(respURL))

	token2 := RegisterKey(location, string(token), entry.keyName)
	fmt.Printf("\n%v\n\n", string(token2))

	// SetInjectorEntry(masterURL string, token string, keyName string, value string, secret bool, path string) []byte
	token3 := SetInjectorEntry(location, string(token), entry)
	fmt.Printf("\n%v\n\n", string(token3))

}
