package kadmin

import (
	"os/exec"

	"github.com/Mellywins/gokrb5-kdc-wrapper/internal/types"
)

type KadminExecutorSpec struct {
	Local                bool
	Realm                string
	Principal            string
	UseKeytab            bool
	keytab               string
	Credentials_cache    string
	password             string
	query                types.Query
	Dbname               string
	Admin_server         string // admin_server[:port]
	Salt                 string // TODO
	AUTH_GSSAPI          bool
	AUTH_GSSAPI_FALLBACK bool
}

func (ex *KadminExecutorSpec) NewKadminExecutor() *KadminExecutorSpec {
	return &KadminExecutorSpec{
		Local:                false,
		Realm:                "",
		Principal:            "",
		UseKeytab:            false,
		keytab:               "",
		Credentials_cache:    "",
		password:             "",
		query:                nil,
		Dbname:               "",
		Admin_server:         "",
		Salt:                 "",
		AUTH_GSSAPI:          false,
		AUTH_GSSAPI_FALLBACK: false,
	}
}
func (ex *KadminExecutorSpec) UseLocal() *KadminExecutorSpec {
	ex.Local = true
	return ex
}

// TODO
func (ex *KadminExecutorSpec) ProbeLiveness() *KadminExecutorSpec {
	return ex
}

// TODO
func (ex *KadminExecutorSpec) Execute(query types.Query) *exec.Cmd {
	return exec.Command("test")
}
