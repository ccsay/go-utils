package encodeutil

import "testing"

func TestCompressAndEncode(t *testing.T) {
	type args struct {
		plain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "compress hello world",
			args:    args{plain: "hello world"},
			want:    "H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQQAAP//hRFKDQsAAAA=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompressAndEncode(tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompressAndEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CompressAndEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUncompressAndDecode(t *testing.T) {
	type args struct {
		cypher string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "uncompress hello world",
			args:    args{cypher: "H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQQAAP//hRFKDQsAAAA="},
			want:    "hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UncompressAndDecode(tt.args.cypher)
			if (err != nil) != tt.wantErr {
				t.Errorf("UncompressAndDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UncompressAndDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
