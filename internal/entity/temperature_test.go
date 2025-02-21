package entity

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTemperature(t *testing.T) {
	tests := []Temperature{
		{-32, -25.6, 241},
		{-10.1, 13.82, 262.9},
		{0, 32, 273},
		{12.35, 54.23, 285.35},
		{28, 82.4, 301},
		{49.5, 121.1, 322.5},
	}

	for _, test := range tests {
		result := NewTemperature(test.TempC)
		assert.Equal(t, test.TempF, math.Round(result.TempF*100)/100)
		assert.Equal(t, test.TempK, math.Round(result.TempK*100)/100)
	}
}
