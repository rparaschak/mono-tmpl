package database

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
)

const (
	ewkbPointType = 1
	ewkbSRIDFlag  = 0x20000000
	ewkbZFlag     = 0x80000000
	ewkbMFlag     = 0x40000000
)

type Geolocation struct {
	Lng float64
	Lat float64
}

func NewGeolocation(lng float64, lat float64) Geolocation {
	return Geolocation{
		Lng: lng,
		Lat: lat,
	}
}

func (g Geolocation) Value() (driver.Value, error) {
	if g.Lat < -90 || g.Lat > 90 {
		return nil, errors.New("latitude must be between -90 and 90")
	}
	if g.Lng < -180 || g.Lng > 180 {
		return nil, errors.New("longitude must be between -180 and 180")
	}

	return fmt.Sprintf("SRID=4326;POINT(%f %f)", g.Lng, g.Lat), nil
}

func (g *Geolocation) Scan(value any) error {
	if value == nil {
		*g = Geolocation{}
		return nil
	}

	bytes, err := geolocationBytes(value)
	if err != nil {
		return err
	}

	lng, lat, err := scanEWKBPoint(bytes)
	if err != nil {
		return err
	}

	*g = Geolocation{
		Lng: lng,
		Lat: lat,
	}
	return nil
}

func geolocationBytes(value any) ([]byte, error) {
	switch typedValue := value.(type) {
	case string:
		bytes, err := hex.DecodeString(typedValue)
		if err != nil {
			return nil, fmt.Errorf("failed to decode Geolocation hex string: %w", err)
		}
		return bytes, nil
	case []byte:
		bytes, err := hex.DecodeString(string(typedValue))
		if err == nil {
			return bytes, nil
		}
		return typedValue, nil
	default:
		return nil, fmt.Errorf("failed to scan Geolocation: invalid data type %T", value)
	}
}

func scanEWKBPoint(bytes []byte) (float64, float64, error) {
	if len(bytes) < 1+4+16 {
		return 0, 0, errors.New("failed to scan Geolocation: invalid EWKB length")
	}

	var order binary.ByteOrder
	switch bytes[0] {
	case 0:
		order = binary.BigEndian
	case 1:
		order = binary.LittleEndian
	default:
		return 0, 0, errors.New("failed to scan Geolocation: invalid EWKB byte order")
	}

	offset := 1
	geometryType := order.Uint32(bytes[offset : offset+4])
	offset += 4

	baseType := geometryType & ^uint32(ewkbSRIDFlag|ewkbZFlag|ewkbMFlag)
	if baseType != ewkbPointType {
		return 0, 0, fmt.Errorf("failed to scan Geolocation: expected Point geometry, got type %d", baseType)
	}

	if geometryType&ewkbSRIDFlag != 0 {
		if len(bytes) < offset+4 {
			return 0, 0, errors.New("failed to scan Geolocation: invalid EWKB SRID length")
		}
		offset += 4
	}

	if len(bytes) < offset+16 {
		return 0, 0, errors.New("failed to scan Geolocation: invalid EWKB point length")
	}

	lng := math.Float64frombits(order.Uint64(bytes[offset : offset+8]))
	lat := math.Float64frombits(order.Uint64(bytes[offset+8 : offset+16]))

	return lng, lat, nil
}
