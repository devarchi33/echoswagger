package echoswagger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParamTypes(t *testing.T) {
	var pa interface{}
	var pb *int64
	var pc map[string]string
	var pd [][]float64
	tests := []struct {
		p     interface{}
		panic bool
		name  string
	}{
		{
			p:     pa,
			panic: true,
			name:  "Interface type",
		},
		{
			p:     &pa,
			panic: true,
			name:  "Interface pointer type",
		},
		{
			p:     &pb,
			panic: false,
			name:  "Int type",
		},
		{
			p:     &pc,
			panic: true,
			name:  "Map type",
		},
		{
			p:     nil,
			panic: true,
			name:  "Nil type",
		},
		{
			p:     0,
			panic: false,
			name:  "Int type",
		},
		{
			p:     &pd,
			panic: false,
			name:  "Array float64 type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := prepareApi()
			if tt.panic {
				assert.Panics(t, func() {
					a.AddParamPath(tt.p, tt.name, "")
				})
			} else {
				a.AddParamPath(tt.p, tt.name, "")
				sapi, ok := a.(*api)
				assert.Equal(t, ok, true)
				assert.Equal(t, len(sapi.operation.Parameters), 1)
				assert.Equal(t, tt.name, sapi.operation.Parameters[0].Name)
			}
		})
	}
}

func TestSchemaTypes(t *testing.T) {
	var pa interface{}
	var pb map[string]string
	type PT struct {
		Name      string
		ExpiredAt time.Time
	}
	var pc map[PT]string
	var pd PT
	var pe map[time.Time]string
	var pf map[*int]string
	type PU struct {
		Unknown interface{}
	}
	var pg PU
	tests := []struct {
		p     interface{}
		panic bool
		name  string
	}{
		{
			p:     pa,
			panic: true,
			name:  "Interface type",
		},
		{
			p:     nil,
			panic: true,
			name:  "Nil type",
		},
		{
			p:     "",
			panic: false,
			name:  "String type",
		},
		{
			p:     &pb,
			panic: false,
			name:  "Map type",
		},
		{
			p:     &pc,
			panic: true,
			name:  "Map struct type",
		},
		{
			p:     pd,
			panic: false,
			name:  "Struct type",
		},
		{
			p:     &pd,
			panic: false,
			name:  "Struct pointer type",
		},
		{
			p:     &pe,
			panic: false,
			name:  "Map time.Time key type",
		},
		{
			p:     &pf,
			panic: false,
			name:  "Map pointer key type",
		},
		{
			p:     &pg,
			panic: true,
			name:  "Struct inner invalid type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := prepareApi()
			if tt.panic {
				assert.Panics(t, func() {
					a.AddParamBody(tt.p, tt.name, "", true)
				})
			} else {
				a.AddParamBody(tt.p, tt.name, "", true)
				sapi, ok := a.(*api)
				assert.Equal(t, ok, true)
				assert.Equal(t, len(sapi.operation.Parameters), 1)
				assert.Equal(t, tt.name, sapi.operation.Parameters[0].Name)
			}
		})
	}
}