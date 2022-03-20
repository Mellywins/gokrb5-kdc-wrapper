package ktadmin
import (
	t "github.com/Mellywins/gokrb5-kdc-wrapper/internal/types"
)
type KadminExecutorSpec struct {
	Local                bool
	Realm                string
	Principal            string
	UseKeytab            bool
	keytab               string
	Credentials_cache    string
	password             string
	Query                t.Query
	Dbname               string
	Admin_server         string // admin_server[:port]
	Salt                 string // TODO
	AUTH_GSSAPI          bool
	AUTH_GSSAPI_FALLBACK bool
}

func (ex *KadminExecutorSpec) NewKadminExecutor() *KadminExecutorSpec{
	return &KadminExecutorSpec{
		Local: false,
		Realm: "",
		Principal: ""
		UserKeytab: false,
		keytab:"",
		Credentials_cache: "",
		password: "",
		Query: nil,
		Dbname: "",
		Admin_server: "",
		Salt: "",
		AUTH_GSSAPI: false,
		AUTH_GSSAPI_FALLBACK: false,

	}
}
func (ex *KadminExecutorSpec) UseLocal() *KadminExecutorSpec {
	ex.Local = true
	return ex
}
