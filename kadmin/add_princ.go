package ktadmin

import "log"

type AddPrincipalType struct {
	randkey     int    // Sets the key of the principal to a random value.
	nokey       int    // Note that you cant use randkey and no key. Using them both will result in an error during the command parsing
	expdate     string // Check the formats accepted by the ktadmin from MIT.
	pwexpdate   string // Password expiery date
	maxlife     string // The maximum ticket life for the principal.
	kvno        int    // Initial key version number
	policy      string // The password policy used by this principal. If not specified, the policy default is used if it exists (unless -clearpolicy is specified).
	clearpolicy int    // Prevents any policy from being assigned when -policy is not specified.

	password     string
	maxrenewlife string // The maximum renewable life of tickets for the principal.
	principal    string // Principal name.
	attributes   string // Attribute list, visit https://web.mit.edu/kerberos/krb5-1.12/doc/admin/admin_commands/kadmin_local.html for more information.
}

/*
Instantiate a new Add_Principal command runner
*/
func AddPrincipal() *AddPrincipalType {
	return &AddPrincipalType{
		randkey:      0,
		nokey:        0,
		expdate:      "",
		pwexpdate:    "",
		maxlife:      "",
		kvno:         -1,
		policy:       "",
		clearpolicy:  0,
		password:     "",
		maxrenewlife: "",
		principal:    "",
		attributes:   "",
	}
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
func (apt *AddPrincipalType) WithRandKey() *AddPrincipalType {
	if apt.nokey == 1 {
		log.Fatal("Cannot Use RANDKEY while NOKEY flag is previously specified")
	}
	apt.randkey = 1
	return apt
}
func (apt *AddPrincipalType) WithNoKey() *AddPrincipalType {
	if apt.randkey == 1 {
		log.Fatal("Cannot use NoKEY flag while RANDKEY flag is previously specified")
	}
	apt.nokey = 1
	return apt
}
func (apt *AddPrincipalType) WithMaxLife(max_life_date string) *AddPrincipalType {
	apt.maxlife = max_life_date
	return apt
}
func (apt *AddPrincipalType) WithClearPolicy() *AddPrincipalType {
	apt.clearpolicy = 1
	return apt
}
func (apt *AddPrincipalType) WithPassword(pw string) *AddPrincipalType {
	apt.password = pw
	return apt
}
func (apt *AddPrincipalType) WithPrincipal(name string) *AddPrincipalType {
	apt.principal = name
	return apt
}
func (apt *AddPrincipalType) AllowPostdated() *AddPrincipalType {

}

func (apt *AddPrincipalType) sanitizeAttributes() *AddPrincipalType {

	return apt
}
