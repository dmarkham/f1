package f1

import (
	"testing"
)

func TestF_Recall(t *testing.T) {
	type fields struct {
		truePositive float64
		type1Error   float64
		trueNegative float64
		type2Error   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			fields: fields{
				truePositive: 5,
				type2Error:   7,
			},
			want: float64(5) / float64(12),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := F1{
				TruePositive: tt.fields.truePositive,
				Type1Error:   tt.fields.type1Error,
				TrueNegative: tt.fields.trueNegative,
				Type2Error:   tt.fields.type2Error,
			}
			if got := f.Recall(); got != tt.want {
				t.Errorf("F.Recall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF_Precision(t *testing.T) {
	type fields struct {
		truePositive float64
		type1Error   float64
		trueNegative float64
		type2Error   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			fields: fields{
				truePositive: 5,
				type1Error:   3,
				type2Error:   7,
			},
			want: float64(5) / float64(8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := F1{
				TruePositive: tt.fields.truePositive,
				Type1Error:   tt.fields.type1Error,
				TrueNegative: tt.fields.trueNegative,
				Type2Error:   tt.fields.type2Error,
			}
			if got := f.Precision(); got != tt.want {
				t.Errorf("F.Precision() = %v, want %v", got, tt.want)
			}
		})
	}
}
