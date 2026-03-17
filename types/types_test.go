package types

import (
	"encoding/json"
	"testing"
)

func TestFlexibleInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected FlexibleInt
		wantErr  bool
	}{
		{
			name:     "valid integer",
			input:    `57`,
			expected: FlexibleInt{Value: 57, IsEmpty: false},
			wantErr:  false,
		},
		{
			name:     "valid 0 integer",
			input:    `0`,
			expected: FlexibleInt{Value: 0, IsEmpty: false},
			wantErr:  false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: FlexibleInt{Value: 0, IsEmpty: true},
			wantErr:  false,
		},
		{
			name:     "string number",
			input:    `"45"`,
			expected: FlexibleInt{Value: 45, IsEmpty: false},
			wantErr:  false,
		},
		{
			name:     "invalid string",
			input:    `"abc"`,
			expected: FlexibleInt{Value: 0, IsEmpty: true},
			wantErr:  true,
		},
		{
			name:     "pass null",
			input:    `null`,
			expected: FlexibleInt{Value: 0, IsEmpty: true},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fi FlexibleInt
			err := json.Unmarshal([]byte(tt.input), &fi)

			if (err != nil) != tt.wantErr {
				t.Errorf("FlexibleInt.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if fi.Value != tt.expected.Value || fi.IsEmpty != tt.expected.IsEmpty {
				t.Errorf("FlexibleInt.UnmarshalJSON() = %+v, want %+v", fi, tt.expected)
			}
		})
	}
}

func TestStringInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringInt
		wantErr  bool
	}{
		{
			name:     "integer value",
			input:    `42`,
			expected: StringInt(42),
			wantErr:  false,
		},
		{
			name:     "string number",
			input:    `"99"`,
			expected: StringInt(99),
			wantErr:  false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: StringInt(0),
			wantErr:  false,
		},
		{
			name:     "null value",
			input:    `null`,
			expected: StringInt(0),
			wantErr:  false,
		},
		{
			name:    "invalid string",
			input:   `"abc"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var si StringInt
			err := json.Unmarshal([]byte(tt.input), &si)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringInt.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && si != tt.expected {
				t.Errorf("StringInt.UnmarshalJSON() = %v, want %v", si, tt.expected)
			}
		})
	}
}

func TestStringFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringFloat64
		wantErr  bool
	}{
		{
			name:     "float value",
			input:    `3.14`,
			expected: StringFloat64(3.14),
			wantErr:  false,
		},
		{
			name:     "string float",
			input:    `"2.71"`,
			expected: StringFloat64(2.71),
			wantErr:  false,
		},
		{
			name:     "string int",
			input:    `"10"`,
			expected: StringFloat64(10),
			wantErr:  false,
		},
		{
			name:     "null value",
			input:    `null`,
			expected: StringFloat64(0),
			wantErr:  false,
		},
		{
			name:    "invalid string",
			input:   `"abc"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sf StringFloat64
			err := json.Unmarshal([]byte(tt.input), &sf)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringFloat64.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && sf != tt.expected {
				t.Errorf("StringFloat64.UnmarshalJSON() = %v, want %v", sf, tt.expected)
			}
		})
	}
}

func TestStringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected StringBool
		wantErr  bool
	}{
		{
			name:     "true bool",
			input:    `true`,
			expected: StringBool(true),
			wantErr:  false,
		},
		{
			name:     "false bool",
			input:    `false`,
			expected: StringBool(false),
			wantErr:  false,
		},
		{
			name:     "Y string",
			input:    `"Y"`,
			expected: StringBool(true),
			wantErr:  false,
		},
		{
			name:     "N string",
			input:    `"N"`,
			expected: StringBool(false),
			wantErr:  false,
		},
		{
			name:     "true string",
			input:    `"true"`,
			expected: StringBool(true),
			wantErr:  false,
		},
		{
			name:     "false string",
			input:    `"false"`,
			expected: StringBool(false),
			wantErr:  false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: StringBool(false),
			wantErr:  false,
		},
		{
			name:    "null value",
			input:   `null`,
			expected: StringBool(false),
			wantErr: false,
		},
		{
			name:    "invalid string",
			input:   `"maybe"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb StringBool
			err := json.Unmarshal([]byte(tt.input), &sb)

			if (err != nil) != tt.wantErr {
				t.Errorf("StringBool.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && sb != tt.expected {
				t.Errorf("StringBool.UnmarshalJSON() = %v, want %v", sb, tt.expected)
			}
		})
	}
}

func TestTourRelatedWinner(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		var w TourRelatedWinner
		err := json.Unmarshal([]byte(`""`), &w)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if w.PlayerId != nil || w.Name != nil || w.CountryName != nil || w.CountryCode != nil {
			t.Errorf("expected all fields nil, got %+v", w)
		}
	})

	t.Run("player object", func(t *testing.T) {
		input := `{"player_id": 7, "name": "Alice", "country_name": "USA", "country_code": "US"}`
		var w TourRelatedWinner
		err := json.Unmarshal([]byte(input), &w)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if w.PlayerId == nil || int(*w.PlayerId) != 7 {
			t.Errorf("expected player_id 7, got %v", w.PlayerId)
		}
		if w.Name == nil || *w.Name != "Alice" {
			t.Errorf("expected name Alice, got %v", w.Name)
		}
		if w.CountryName == nil || *w.CountryName != "USA" {
			t.Errorf("expected country_name USA, got %v", w.CountryName)
		}
		if w.CountryCode == nil || *w.CountryCode != "US" {
			t.Errorf("expected country_code US, got %v", w.CountryCode)
		}
	})
}

// Example of how to use these types in a struct
type PlayerExample struct {
	PlayerID    string         `json:"player_id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Age         FlexibleInt    `json:"age"`          // Can be number or empty string
}

func TestPlayerExampleUnmarshal(t *testing.T) {
	jsonData := `{
		"player_id": "123",
		"first_name": "John",
		"last_name": "Doe",
		"age": 45
	}`

	var player PlayerExample
	err := json.Unmarshal([]byte(jsonData), &player)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if player.Age.GetValue() != 45 {
		t.Errorf("expected age 45, got %d", player.Age.GetValue())
	}

	if !player.Age.HasValue() {
		t.Error("expected age to have value")
	}
}

func TestStringIntString(t *testing.T) {
	if StringInt(42).String() != "42" {
		t.Errorf("StringInt.String() = %s, want 42", StringInt(42).String())
	}
}

func TestStringFloat64String(t *testing.T) {
	sf := StringFloat64(1.5)
	if sf.String() != "1.500000" {
		t.Errorf("StringFloat64.String() = %s, want 1.500000", sf.String())
	}
}

func TestStringBoolString(t *testing.T) {
	if StringBool(true).String() != "true" {
		t.Errorf("StringBool(true).String() = %s, want true", StringBool(true).String())
	}
	if StringBool(false).String() != "false" {
		t.Errorf("StringBool(false).String() = %s, want false", StringBool(false).String())
	}
}

func TestFlexibleIntMarshalJSON(t *testing.T) {
	t.Run("with value", func(t *testing.T) {
		fi := FlexibleInt{Value: 57, IsEmpty: false}
		data, err := json.Marshal(fi)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(data) != "57" {
			t.Errorf("FlexibleInt.MarshalJSON() = %s, want 57", data)
		}
	})
	t.Run("empty", func(t *testing.T) {
		fi := FlexibleInt{IsEmpty: true}
		data, err := json.Marshal(fi)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(data) != `""` {
			t.Errorf(`FlexibleInt.MarshalJSON() = %s, want ""`, data)
		}
	})
}

func TestFlexibleIntUnmarshalJSONInvalid(t *testing.T) {
	var fi FlexibleInt
	if err := json.Unmarshal([]byte(`true`), &fi); err == nil {
		t.Error("expected error for boolean input, got nil")
	}
}

func TestPlayerExampleUnmarshalEmpty(t *testing.T) {
	jsonData := `{
		"player_id": "456",
		"first_name": "Jane",
		"last_name": "Smith",
		"age": ""
	}`

	var player PlayerExample
	err := json.Unmarshal([]byte(jsonData), &player)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	if player.Age.HasValue() {
		t.Error("expected age to be empty")
	}
}
