package log

import "testing"

func TestSanitizeForLog(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		want  string
	}{
		{
			name: "no_control_chars",
			in:   "GET /health",
			want: "GET /health",
		},
		{
			name: "newline_and_carriage_return",
			in:   "line1\nline2\rline3",
			want: "line1 line2 line3",
		},
		{
			name: "tab_and_null",
			in:   "a\tb\x00c",
			want: "a b c",
		},
		{
			name: "del_control",
			in:   "a\x7fb",
			want: "a b",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeForLog(tt.in); got != tt.want {
				t.Fatalf("sanitizeForLog(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
