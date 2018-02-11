package hmac

import "testing"

func TestAddAnd(t *testing.T) {
	if s := addAnd("a"); s != "&" {
		t.Errorf("Function addAnd is wrong for have string")
	}

	if s := addAnd(""); s != "" {
		t.Errorf("Function addAnd is wrong for empty string")
	}
}

func TestHashTagForType(t *testing.T) {
	hash := new(HashTag)
	hash.SetType("node")

	if s := hash.String(); s != "type=node" {
		t.Errorf("Test type string fail")
	}
}

func TestHashTagForID(t *testing.T) {
	hash := new(HashTag)
	hash.SetID("id")

	if s := hash.String(); s != "id=id" {
		t.Errorf("Test id string fail")
	}
}

func TestHashTagForOption(t *testing.T) {
	hash := new(HashTag)
	hash.AddOption(Option{"k1", "v1"})

	if s := hash.String(); s != "option[k1]=v1" {
		t.Errorf("Test One option fail")
	}

	hash.AddOption(Option{"k2", "v2"})
	if s := hash.String(); s != "option[k1]=v1&option[k2]=v2" {
		t.Errorf("Test Two option fail")
	}
}

func TestGetFullQueryString(t *testing.T) {
	hash := new(HashTag)
	hash.SetKey("1231232131232131231212312312")
	hash.SetType("node")
	hash.SetID("1232132")
	hash.AddOption(Option{"nocache", "1"})
	hash.AddOption(Option{"app_id", "2223432"})

	if s := hash.GetHmacWithKey(); s != "4e7e79539aa7cf06546a06b67e7d1193a0fbe9e48adac0c171a3314ad36a49a2" {
		t.Errorf("GetHmacWithKey fail")
	}

	if s := hash.GetFullQueryString(); s != "type=node&id=1232132&option[nocache]=1&option[app_id]=2223432&hash=4e7e79539aa7cf06546a06b67e7d1193a0fbe9e48adac0c171a3314ad36a49a2" {
		t.Errorf("GetFullQueryString fail")
	}
}
