package ktadmin

type AddPrincipalType struct {
	randkey      bool
	nokey        bool
	expdate      string // Check the formats accepted by the ktadmin from MIT.
	pwexpdate    string // Password expiery date
	maxlife      string //
	kvno         int    //
	policy       string // policy to include in the principal
	clearpolicy  bool
	password     string
	maxrenewlife string
}

func AddPrincipal() *AddPrincipalType {
	return &AddPrincipalType{}
}

func (apt *AddPrincipalType) WithExpDate(date string) *AddPrincipalType {
	apt.expdate = date
	return apt
}
func (apt *AddPrincipalType) WithPwExpDate(date string) *AddPrincipalType {
	apt.pwexpdate = date
	return apt
}
func (apt *AddPrincipalType) WithKvno(kvno int) *AddPrincipalType {
	apt.kvno = kvno
	return apt
}
func (apt *AddPrincipalType) WithPolicy(policy string) *AddPrincipalType {
	apt.policy = policy
	return apt
}
