package ktadmin

type KtadminExecutorSpec struct {
	local bool
}

func (ex *KtadminExecutorSpec) UseLocal() {
	ex.local = true
}
