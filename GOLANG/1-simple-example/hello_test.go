package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Bara", "Khadija", "Salim"},
			result: "Hello, Bara, Khadija, Salim!",
		},
		{
			items:  []string{"Mahmoud"},
			result: "Hello, Mahmoud!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted [[ %v ]], but got %v \n", st.result, s)
		}
	}

}
