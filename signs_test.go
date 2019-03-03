package main

import (
	"testing"
	"time"
)

func TestToDisplaySigns(t *testing.T) {
	ss := []sign{
		sign{
			Location: "123 prior to 456",
			Name:     "555",
			Message:  "11 MI<BR>12 MIN",
			Updated:  updatedTime(time.Now()),
			Beacon:   false,
		},
		sign{
			Location: "789 prior to 012",
			Name:     "777",
			Message:  "Crash before exit 5<BR/>Use alternate route<br>Drive safely",
			Updated:  updatedTime(time.Now()),
			Beacon:   true,
		},
	}

	expLines := []int{2, 3}

	ds := toDisplaySigns(ss)

	if len(ds) != len(ss) {
		t.Fatalf("Expected %d display signs, got %d", len(ss), len(ds))
	}

	for i := range ds {
		if len(ds[i].MessageLines) != expLines[i] {
			t.Fatalf("Wrong number of lines for sign %d. Expected %d, got %d", i, len(ds[i].MessageLines), expLines[i])
		}
	}
}

func TestFindByName(t *testing.T) {
	ss := []sign{
		sign{
			Location: "123 prior to 456",
			Name:     "555",
			Message:  "11 MI<BR>12 MIN",
			Updated:  updatedTime(time.Now()),
			Beacon:   false,
		},
		sign{
			Location: "789 prior to 012",
			Name:     "777",
			Message:  "Crash before exit 5<BR/>Use alternate route<br>Drive safely",
			Updated:  updatedTime(time.Now()),
			Beacon:   true,
		},
		sign{
			Location: "543 prior to 210",
			Name:     "888",
			Message:  "Icy conditions",
			Updated:  updatedTime(time.Now()),
			Beacon:   true,
		},
	}

	ms := messageSigns{ss}

	found := ms.FindByName("555", "888")

	if len(found) != 2 {
		t.Fatalf("Expected 2 results, Got %d", len(found))
	}

	if found[0].Name != "555" {
		t.Fatalf("Expected first result to be '555', got %s", found[0].Name)
	}

	if found[1].Name != "888" {
		t.Fatalf("Expected second result to be '555', got %s", found[1].Name)
	}
}
