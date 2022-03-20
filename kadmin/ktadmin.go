package ktadmin

type KtadminExecutorSpec struct {
	local                bool
	realm                string
	principal            string
	keytab               string
	credentials_cache    string
	password             string
	query                string
	dbname               string
	admin_server         string // admin_server[:port]
	salt                 string // TODO
	AUTH_GSSAPI          bool
	AUTH_GSSAPI_FALLBACK bool
}

func (ex *KtadminExecutorSpec) UseLocal() {
	ex.local = true
}
