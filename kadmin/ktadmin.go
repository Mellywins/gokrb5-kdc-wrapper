package ktadmin

type KtadminExecutorSpec struct {
	local bool   // kadmin.local [-r realm] [-p principal] [-q query] [-d dbname] [-e enc:salt ...] [-m] [-x db_args]. Enabling this is the same as using kadmin.local instead of kadmin
	realm string // Use realm as the default database realm.

}

func (ex *KtadminExecutorSpec) UseLocal() {
	ex.local = true
}
