package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// HashKey key
//const HashKey = "1231232131232131231212312312"

// HashTag use for create something
type HashTag struct {
	key    string
	t      string   //type , cannot use type cos it is golang special
	id     string   // id for the node
	option []Option // map type, when for loop , the order is random...fuk, now use slice
}

//Option for option
type Option struct {
	k string
	v string
}

// format struct to string, url.value.add seem will change the order >.<
func (t HashTag) String() string {

	s := ""

	if t.t != "" {
		s = "type=" + t.t
	}

	if t.id != "" {
		s += addAnd(s) + "id=" + t.id
	}

	if len(t.option) != 0 {
		for _, v := range t.option {
			s += addAnd(s) + "option[" + v.k + "]=" + v.v
		}
	}

	return s
}

// SetKey for the Key
func (t *HashTag) SetKey(s string) {
	t.key = s
}

// SetType for the T
func (t *HashTag) SetType(s string) {
	t.t = s
}

// SetID for the id
func (t *HashTag) SetID(s string) {
	t.id = s
}

// AddOption for add option
func (t *HashTag) AddOption(o Option) {
	t.option = append(t.option, o)
}

/*
func main() {
	//todo try builder patteren
	t := new(HashTag)

	t.SetKey(HashKey)
	t.SetType("node")
	t.SetID("1232132")
	t.AddOption(Option{"nocache", "1"})
	t.AddOption(Option{"app_id", "2223432"})
	//t.Option = make(map[string]string)
	//t.Option["nocache"] = "1"
	//t.Option["app_id"] = "2223432"

	fmt.Println(t)

	//t2 := t.GetHmacWithKey()
	for i := 1; i <= 20; i++ {
		fmt.Println(t.GetHmacWithKey())
	}

	for i := 1; i <= 20; i++ {
		fmt.Println(t.GetFullQueryString())
	}

}*/

//GetFullQueryString get full url string
func (t *HashTag) GetFullQueryString() string {
	//t2 := t.GetHmacWithKey()
	//fmt.Println(t2)
	return t.String() + "&hash=" + t.GetHmacWithKey()
}

// GetHmacWithKey Get Hashed string with key
func (t *HashTag) GetHmacWithKey() string {

	key := []byte(t.key)
	message := []byte(t.String())

	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	// now can directly use fmt.Printin(hash)
	return fmt.Sprintf("%s", hex.EncodeToString(hash.Sum(nil)))
	// below code need use fmt.Printf("%s",hash)
	// so i cannot use the fmt.Printin later
	//return hex.EncodeToString(hash.Sum(nil))
}

// helper format string
func addAnd(s string) string {
	if s != "" {
		return "&"
	}
	return ""
}
