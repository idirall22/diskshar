package userAccount

import (
	"testing"
)

// TestValidatePassword
func TestValidatePassword(t *testing.T) {
	passwords := []string{
		// Password length should be minimum >= minPassLength
		"apM47*",
		// Password length should be maximum <= maxPassLength
		"apM47*dùmlj,fdqsg^OJOPojkfd$pfldskpsg45894*/*--dslfkfdsùsldfapM47*dùmlj,fdqsg^OJOPojkfd$pfldskpsg45894*/*--dslfkfdsùsldfpfldskpsg45894*/*--",
		// strong password
		"djzoPOdhjlde6548/*",
	}

	for i, pass := range passwords {
		err := validatePassword(pass)
		switch i {
		case 0:
			if err != ErrorMinPass {
				t.Errorf("Error should be %s but got %s", ErrorMinPass, err)
			}
			break
		case 1:
			if err != ErrorMaxPass {
				t.Errorf("Error should be %s but got %s", ErrorMaxPass, err)
			}
			break
		case 2:
			if err != nil {
				t.Errorf("Error should be nil but got %s", err)
			}
			break
		}

	}
}
