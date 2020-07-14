package pgconn_test

import (
	"testing"

	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestConfigError(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		expectedMsg string
	}{
		{
			name:        "url with password",
			err:         pgconn.NewParseConfigError("postgresql://foo:password@host", "msg", nil),
			expectedMsg: "cannot parse `postgresql://foo:xxxxx@host`: msg",
		},
		{
			name:        "dsn with password unquoted",
			err:         pgconn.NewParseConfigError("host=host password=password user=user", "msg", nil),
			expectedMsg: "cannot parse `host=host password=xxxxx user=user`: msg",
		},
		{
			name:        "dsn with password quoted",
			err:         pgconn.NewParseConfigError("host=host password='pass word' user=user", "msg", nil),
			expectedMsg: "cannot parse `host=host password=xxxxx user=user`: msg",
		},
		{
			name:        "weird url",
			err:         pgconn.NewParseConfigError("postgresql://foo::pasword@host:1:", "msg", nil),
			expectedMsg: "cannot parse `postgresql://foo:xxxxx@host:1:`: msg",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.EqualError(t, tt.err, tt.expectedMsg)
		})
	}
}