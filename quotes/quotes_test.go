package quotes

import "testing"

func basic_test(t *testing.T) {

	favs := Favs()

	if len(favs) != 2 {
		t.Errorf("favs does not have the correct length")
	}
}
