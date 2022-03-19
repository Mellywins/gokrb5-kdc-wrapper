package ktadmin

import (
	"fmt"
	"log"
)

type Executable interface {
	ParseCommand()
	Exec()
}

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
	maxrenewlife string                 // The maximum renewable life of tickets for the principal.
	principal    string                 // Principal name.
	attributes   AddPrincipalAttributes // Attribute list, visit https://web.mit.edu/kerberos/krb5-1.12/doc/admin/admin_commands/kadmin_local.html for more information.
}

// Attributes holder for the different options provided by Kerberos MIT.
type AddPrincipalAttributes struct {
	allow_postdated   int // -allow_postdated prohibits this principal from obtaining postdated tickets. +allow_postdated clears this flag.
	allow_forwardable int // -allow_forwardable prohibits this principal from obtaining forwardable tickets. +allow_forwardable clears this flag.
	allow_renewable   int // -allow_renewable prohibits this principal from obtaining renewable tickets. +allow_renewable clears this flag.
	allow_proxiable   int // -allow_proxiable prohibits this principal from obtaining proxiable tickets. +allow_proxiable clears this flag.
	allow_dup_key     int // -allow_dup_skey disables user-to-user authentication for this principal by prohibiting this principal from obtaining a session key for another user. +allow_dup_skey clears this flag.
	requires_preauth  int // +requires_preauth requires this principal to preauthenticate before being allowed to kinit. -requires_preauth clears this flag. When +requires_preauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using preauthentication.
	requires_hwauth   int // +requires_hwauth requires this principal to preauthenticate using a hardware device before being allowed to kinit. -requires_hwauth clears this flag. When +requires_hwauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using a hardware device to preauthenticate.
	ok_as_delegate    int // +ok_as_delegate sets the okay as delegate flag on tickets issued with this principal as the service. Clients may use this flag as a hint that credentials should be delegated when authenticating to the service. -ok_as_delegate clears this flag.
	allow_svr         int // -allow_svr prohibits the issuance of service tickets for this principal. +allow_svr clears this flag.

	allow_tgs_req int // -allow_tgs_req specifies that a Ticket-Granting Service (TGS) request for a service ticket for this principal is not permitted. +allow_tgs_req clears this flag.
	allow_tix     int // -allow_tix forbids the issuance of any tickets for this principal. +allow_tix clears this flag.

	needchange                int // +needchange forces a password change on the next initial authentication to this principal. -needchange clears this flag.
	password_changing_service int // +password_changing_service marks this principal as a password change service principal.

	ok_to_auth_as_delegate int // +ok_to_auth_as_delegate allows this principal to acquire forwardable tickets to itself from arbitrary users, for use with constrained delegation.
	no_auth_data_required  int // +no_auth_data_required prevents PAC or AD-SIGNEDPATH data from being added to service tickets for the principal.
}

var IntToSymbolMap = map[int]string{0: "-", 1: "+"}

/*
Instantiate a new Add_Principal command runner.
Not setting a flag to a value of your chooseing implies that the KDC command will use its default value.
*/
func AddPrincipal(atts AddPrincipalAttributes) *AddPrincipalType {
	return &AddPrincipalType{
		randkey:      -1,
		nokey:        -1,
		expdate:      "",
		pwexpdate:    "",
		maxlife:      "",
		kvno:         -1,
		policy:       "",
		clearpolicy:  -1,
		password:     "",
		maxrenewlife: "",
		principal:    "",
		attributes:   atts,
	}
}

func (apt *AddPrincipalType) ParseCommand() {
	fmt.Println("Parsed Command")
}
func (apt *AddPrincipalType) Exec() {
	fmt.Println("Execution ...")
}
func (apt *AddPrincipalType) WithAttributes(atts AddPrincipalAttributes) *AddPrincipalType {
	apt.attributes = atts
	return apt
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

func (apt *AddPrincipalType) sanitizeAttributes() *AddPrincipalType {

	return apt
}

// -allow_postdated(0) prohibits this principal from obtaining postdated tickets. +allow_postdated(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetPostdated(n int) *AddPrincipalAttributes {
	apatts.allow_postdated = n
	return apatts
}

// -allow_forwardable(0) prohibits this principal from obtaining forwardable tickets. +allow_forwardable (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetForwardable(n int) *AddPrincipalAttributes {
	apatts.allow_forwardable = n
	return apatts
}

// -allow_renewable(0) prohibits this principal from obtaining renewable tickets. +allow_renewable(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetRenewable(n int) *AddPrincipalAttributes {
	apatts.allow_renewable = n
	return apatts
}

// -allow_proxiable(0) prohibits this principal from obtaining proxiable tickets. +allow_proxiable(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetProxiable(n int) *AddPrincipalAttributes {
	apatts.allow_proxiable = n
	return apatts
}

// -allow_dup_skey(0) disables user-to-user authentication for this principal by prohibiting this principal from obtaining a session key for another user. +allow_dup_skey (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetDupKey(n int) *AddPrincipalAttributes {
	apatts.allow_dup_key = n
	return apatts
}

// +requires_preauth(1) requires this principal to preauthenticate before being allowed to kinit. -requires_preauth(0) clears this flag. When +requires_preauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using preauthentication.
func (apatts *AddPrincipalAttributes) SetPreAuth(n int) *AddPrincipalAttributes {
	apatts.requires_preauth = n
	return apatts
}

// +requires_hwauth(1) requires this principal to preauthenticate using a hardware device before being allowed to kinit. -requires_hwauth(0) clears this flag. When +requires_hwauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using a hardware device to preauthenticate.
func (apatts *AddPrincipalAttributes) SetHwAuth(n int) *AddPrincipalAttributes {
	apatts.requires_hwauth = n
	return apatts
}

// +ok_as_delegate(1) sets the okay as delegate flag on tickets issued with this principal as the service. Clients may use this flag as a hint that credentials should be delegated when authenticating to the service. -ok_as_delegate(0) clears this flag.
func (apatts *AddPrincipalAttributes) SetOkAsDelegate(n int) *AddPrincipalAttributes {
	apatts.ok_as_delegate = n
	return apatts
}

// -allow_svr(0) prohibits the issuance of service tickets for this principal. +allow_svr clears this flag(1).
func (apatts *AddPrincipalAttributes) SetSvr(n int) *AddPrincipalAttributes {
	apatts.allow_svr = n
	return apatts
}

// -allow_tgs_req(0) specifies that a Ticket-Granting Service (TGS) request for a service ticket for this principal is not permitted. +allow_tgs_req(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetTgsReq(n int) *AddPrincipalAttributes {
	apatts.allow_tgs_req = n
	return apatts
}

// -allow_tix(0) forbids the issuance of any tickets for this principal. +allow_tix (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetTix(n int) *AddPrincipalAttributes {
	apatts.allow_tix = n
	return apatts
}

// +needchange(1) forces a password change on the next initial authentication to this principal. -needchange(0) clears this flag.
func (apatts *AddPrincipalAttributes) SetNeedChange(n int) *AddPrincipalAttributes {
	apatts.needchange = n
	return apatts
}

// +password_changing_service(1) marks this principal as a password change service principal.
func (apatts *AddPrincipalAttributes) SetPasswordChangingService() *AddPrincipalAttributes {
	apatts.password_changing_service = 1
	return apatts
}

// +ok_to_auth_as_delegate allows this principal to acquire forwardable tickets to itself from arbitrary users, for use with constrained delegation.
func (apatts *AddPrincipalAttributes) SetOkToAuthAsDelegate() *AddPrincipalAttributes {
	apatts.ok_to_auth_as_delegate = 1
	return apatts
}

// +no_auth_data_required prevents PAC or AD-SIGNEDPATH data from being added to service tickets for the principal.
func (apatts *AddPrincipalAttributes) SetNoAuthDataRequired() *AddPrincipalAttributes {
	apatts.no_auth_data_required = 1
	return apatts
}
func CreateAddPrincipalAttributes() *AddPrincipalAttributes {
	return &AddPrincipalAttributes{
		allow_postdated:           -1,
		allow_forwardable:         -1,
		allow_renewable:           -1,
		allow_proxiable:           -1,
		allow_dup_key:             -1,
		requires_preauth:          -1,
		requires_hwauth:           -1,
		ok_as_delegate:            -1,
		allow_svr:                 -1,
		allow_tgs_req:             -1,
		allow_tix:                 -1,
		needchange:                -1,
		password_changing_service: -1,
		ok_to_auth_as_delegate:    -1,
		no_auth_data_required:     -1,
	}
}
