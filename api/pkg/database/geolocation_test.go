package database

import (
	"math"
	"testing"
)

func TestGeolocationValue(t *testing.T) {
	value, err := NewGeolocation(21.0122, 52.2297).Value()
	if err != nil {
		t.Fatalf("Value() error = %v", err)
	}

	want := "SRID=4326;POINT(21.012200 52.229700)"
	if value != want {
		t.Fatalf("Value() = %q, want %q", value, want)
	}
}

func TestGeolocationValueValidatesLatitude(t *testing.T) {
	_, err := NewGeolocation(21.0122, 91).Value()
	if err == nil {
		t.Fatal("Value() error = nil, want validation error")
	}
}

func TestGeolocationValueValidatesLongitude(t *testing.T) {
	_, err := NewGeolocation(181, 52.2297).Value()
	if err == nil {
		t.Fatal("Value() error = nil, want validation error")
	}
}

func TestGeolocationScan(t *testing.T) {
	var geolocation Geolocation

	err := geolocation.Scan("0101000020E6100000F46C567DAE033540A835CD3B4E1D4A40")
	if err != nil {
		t.Fatalf("Scan() error = %v", err)
	}

	if math.Abs(geolocation.Lng-21.01438125) > 0.00000001 {
		t.Fatalf("Lng = %.13f, want %.13f", geolocation.Lng, 21.01438125)
	}
	if math.Abs(geolocation.Lat-52.22895) > 0.00000001 {
		t.Fatalf("Lat = %.13f, want %.13f", geolocation.Lat, 52.22895)
	}
}
