package pkg

import "testing"

func Test_GetIps(t *testing.T) {
	rs := GetIPs()
	target := "192.168.0.113"
	t.Log("rs", rs)
	if ans := In(target, rs); !ans {
		t.Fatalf("expected %s in %v,but got %v", target, true, ans)
	}
}

func Test_GetAllFiles(t *testing.T) {
	expected := []string{"/static/common.go", "/static/common_test.go", "/static/ember.go", "/static/httpstatic.go", "/static/main.go", "/static/method.go", "/static/template.go"}
	rs, err := GetAllFiles(".", []string{".go"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("GetAllFiles ", rs)
	for _, x := range rs {
		if !In(x, expected) {
			t.Errorf("%s expected in %v,but got false", x, rs)
		}
	}
}

var InData []string = []string{"this", "is", "test", "data", "for", "in", "function"}

func Test_In(t *testing.T) {
	t.Run("bingo", func(t *testing.T) {
		if In("this", InData) == false {
			t.Fatalf("this expected in InData [%v], but got fasle", InData)
		}
	})

	cases := []struct {
		Name    string
		A       string
		Expectd bool
	}{
		{"yes", "this", true},
		{"no", "ok", false},
		{"error", "err", false},
	}

	for _, x := range cases {
		t.Run(x.Name, func(t *testing.T) {
			if ans := In(x.A, InData); ans != x.Expectd {
				t.Fatalf("%s expected %v,but got %v", x.A, InData, ans)
			}
		})
	}
}
