package config

import "testing"

func TestValidation(t *testing.T) {
	ted := &Member{
		Name:    "Ted Hahn",
		Id:      7,
		IdCards: []string{"Foo", "Bar", "baz"},
	}
	jp := &Member{
		Name:    "JP Sugarbroad",
		Id:      2,
		IdCards: []string{"jpsugar"},
	}
	dave := &Member{
		Name:    "David Stansel-Garner",
		Id:      3,
		IdCards: []string{"qux"},
	}
	sverre := &Member{
		Name:    "Sverre Rabbelier",
		Id:      17,
		IdCards: []string{"xyzzy"},
	}
	err := &Member{
		Name:    "ERROR",
		Id:      0,
		IdCards: []string{"ERROR"},
	}
	if !validateCardlist(&Cardlist{"Ted Hahn": ted,
		"JP Sugarbroad":        jp,
		"David Stansel-Garner": dave,
		"Sverre Rabbelier":     sverre}) {
		t.Fail()
	}
	if validateCardlist(&Cardlist{"Ted Hahn": ted,
		"JP Sugarbroad":        jp,
		"David Stansel-Garner": dave,
		"ERROR":                err}) {
		t.Fail()
	}
	if validateCardlist(&Cardlist{"Ted Hahn": ted,
		"David Stansel-Garner": dave,
		"ERROR":                err}) {
		t.Fail()
	}
}

func TestEmptyValidation(t *testing.T) {
	if validateCardlist(&Cardlist{}) {
		t.Fail()
	}
}
