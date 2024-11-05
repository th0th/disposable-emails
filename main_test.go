package disposableemail

import (
	"reflect"
	"testing"
)

func Test_service_Check(t *testing.T) {
	type fields struct {
		domainsMap map[string]bool
	}
	type args struct {
		emailOrDomain string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CheckResult
	}{
		{
			name: "disposable domain",
			fields: fields{
				domainsMap: map[string]bool{
					"domain.tld": true,
				},
			},
			args: args{
				emailOrDomain: "domain.tld",
			},
			want: &CheckResult{
				IsDisposable: true,
			},
		},
		{
			name: "non-disposable domain",
			fields: fields{
				domainsMap: map[string]bool{
					"domain.tld": true,
				},
			},
			args: args{
				emailOrDomain: "domain2.tld",
			},
			want: &CheckResult{
				IsDisposable: false,
			},
		},
		{
			name: "e-mail address with disposable domain",
			fields: fields{
				domainsMap: map[string]bool{
					"domain.tld": true,
				},
			},
			args: args{
				emailOrDomain: "user@domain.tld",
			},
			want: &CheckResult{
				IsDisposable: true,
			},
		},
		{
			name: "e-mail address non-disposable domain",
			fields: fields{
				domainsMap: map[string]bool{
					"domain.tld": true,
				},
			},
			args: args{
				emailOrDomain: "user@domain2.tld",
			},
			want: &CheckResult{
				IsDisposable: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				domainsMap: tt.fields.domainsMap,
			}
			if got := s.Check(tt.args.emailOrDomain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
