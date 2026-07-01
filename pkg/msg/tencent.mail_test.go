package msg

import "testing"

func TestTencentConfig_FromEmailAddress(t *testing.T) {
	tests := []struct {
		name    string
		cfg     TencentConfig
		want    string
		wantErr bool
	}{
		{
			name: "email only",
			cfg:  TencentConfig{SenderEmail: "noreply@mail.example.com"},
			want: "noreply@mail.example.com",
		},
		{
			name: "alias and email",
			cfg:  TencentConfig{SenderEmail: "noreply@mail.example.com", SenderAlias: "通知"},
			want: "通知 <noreply@mail.example.com>",
		},
		{
			name: "legacy full address",
			cfg:  TencentConfig{Sender: "Team <noreply@mail.example.com>"},
			want: "Team <noreply@mail.example.com>",
		},
		{
			name: "legacy email only",
			cfg:  TencentConfig{Sender: "noreply@mail.example.com"},
			want: "noreply@mail.example.com",
		},
		{
			name:    "missing address",
			cfg:     TencentConfig{AppId: "x"},
			wantErr: true,
		},
		{
			name:    "colon in alias",
			cfg:     TencentConfig{SenderEmail: "a@b.com", SenderAlias: "bad:name"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cfg.FromEmailAddress()
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q want %q", got, tt.want)
			}
		})
	}
}
