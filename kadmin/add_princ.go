package kadmin

import (
	"fmt"
	"log"
	"os/exec"
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

	password      string
	maxrenewlife  string                 // The maximum renewable life of tickets for the principal.
	principal     string                 // Principal name.
	CommandString string                 // String containing the commands progressively built with the available Builders
	attributes    AddPrincipalAttributes // Attribute list, visit https://web.mit.edu/kerberos/krb5-1.12/doc/admin/admin_commands/kadmin_local.html for more information.
	verbose       bool                   // Set to true to enable logging the command to the std.out, Defaults to false to omit the log traces.
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

	ok_to_auth_as_delegate int    // +ok_to_auth_as_delegate allows this principal to acquire forwardable tickets to itself from arbitrary users, for use with constrained delegation.
	no_auth_data_required  int    // +no_auth_data_required prevents PAC or AD-SIGNEDPATH data from being added to service tickets for the principal.
	CommandString          string // String containing the commands progressively built with the available Builders

}

var IntToSymbolMap = map[int]string{0: "-", 1: "+"}

// Type that indicates Wether we omit the flag (incase it wasn't specified by the user), set +flag or -flag.

const (
	Disallow int = iota // Value of 0 -> Translated to - in the kadmin flag
	Allow               // Value of 1 -> Translated to + in the kadmin flag
)

/*
Instantiate a new Add_Principal command runner.
Not setting a flag to a value of your chooseing implies that the KDC command will use its default value.
*/
func AddPrincipal(atts AddPrincipalAttributes) *AddPrincipalType {
	return &AddPrincipalType{
		randkey:       -1,
		nokey:         -1,
		expdate:       "",
		pwexpdate:     "",
		maxlife:       "",
		kvno:          -1,
		policy:        "",
		clearpolicy:   -1,
		password:      "",
		maxrenewlife:  "",
		principal:     "",
		CommandString: "",
		attributes:    atts,
		verbose:       false,
	}
}

func (apt *AddPrincipalType) ParseCommand() *AddPrincipalType {
	apt.CommandString += apt.attributes.CommandString
	fmt.Printf("Parsed Command: kadmin.local add_principal %s %s ", apt.CommandString, apt.principal)
	return apt
}
func (apt *AddPrincipalType) Exec() *exec.Cmd {

	fmt.Println("Execution ...")
	return exec.Command("ls")
}
func (apt *AddPrincipalType) WithAttributes(atts AddPrincipalAttributes) *AddPrincipalType {
	apt.attributes = atts
	return apt
}
func (aptt *AddPrincipalAttributes) appendFlag(n int, flag string) {
	aptt.CommandString += " " + IntToSymbolMap[n] + flag + " "
}

func (apt *AddPrincipalType) appendFlag(value interface{}, flag string) {
	if str, ok := value.(string); ok {
		apt.CommandString += " -" + flag + " " + str
	} else {
		apt.CommandString += " -" + flag + " " + fmt.Sprint(value.(int))
	}
}
func (aptt *AddPrincipalType) SetVerbosity(b bool) *AddPrincipalType {
	aptt.verbose = b
	return aptt
}

// (getdate time string) The expiration date of the principal.

func (apt *AddPrincipalType) WithExpDate(date string) *AddPrincipalType {
	apt.expdate = date
	apt.appendFlag(date, "expire")
	return apt
}

// (getdate time string) The password expiration date.

func (apt *AddPrincipalType) WithPwExpDate(date string) *AddPrincipalType {
	apt.pwexpdate = date
	apt.appendFlag(date, "pwexpire")
	return apt
}

// The initial key version number.

func (apt *AddPrincipalType) WithKvno(kvno int) *AddPrincipalType {
	apt.kvno = kvno
	apt.appendFlag(kvno, "kvno")

	return apt
}

// The password policy used by this principal. If not specified, the policy default is used if it exists (unless -clearpolicy is specified).
func (apt *AddPrincipalType) WithPolicy(policy string) *AddPrincipalType {
	apt.policy = policy
	apt.appendFlag(policy, "policy")
	return apt
}

// Sets the key of the principal to a random value.

func (apt *AddPrincipalType) WithRandKey() *AddPrincipalType {
	if apt.nokey == 1 {
		log.Fatal("Cannot Use RANDKEY while NOKEY flag is previously specified")
	}
	apt.randkey = 1
	apt.CommandString += " -randkey "
	return apt
}
func (apt *AddPrincipalType) WithNoKey() *AddPrincipalType {
	if apt.randkey == 1 {
		log.Fatal("Cannot use NoKEY flag while RANDKEY flag is previously specified")
	}
	apt.nokey = 1
	apt.CommandString += " -nokey "
	return apt
}

// (getdate time string) The maximum renewable life of tickets for the principal.

func (apt *AddPrincipalType) WithMaxLife(max_life_date string) *AddPrincipalType {
	apt.maxlife = max_life_date
	apt.appendFlag(max_life_date, "maxrenewlife")
	return apt
}

// Prevents any policy from being assigned when -policy is not specified.

func (apt *AddPrincipalType) WithClearPolicy() *AddPrincipalType {
	apt.clearpolicy = 1
	apt.CommandString += " -clearpolicy "
	return apt
}

// Sets the password of the principal to the specified string and does not prompt for a password. Note: using this option in a shell script may expose the password to other users on the system via the process list.
func (apt *AddPrincipalType) WithPassword(pw string) *AddPrincipalType {
	apt.password = pw
	apt.appendFlag(pw, "pw")
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
	apatts.appendFlag(n, "allow_postdated")
	return apatts
}

// -allow_forwardable(0) prohibits this principal from obtaining forwardable tickets. +allow_forwardable (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetForwardable(n int) *AddPrincipalAttributes {
	apatts.allow_forwardable = n
	apatts.appendFlag(n, "allow_forwardable")
	return apatts
}

// -allow_renewable(0) prohibits this principal from obtaining renewable tickets. +allow_renewable(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetRenewable(n int) *AddPrincipalAttributes {
	apatts.allow_renewable = n
	apatts.appendFlag(n, "allow_renewable")
	return apatts
}

// -allow_proxiable(0) prohibits this principal from obtaining proxiable tickets. +allow_proxiable(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetProxiable(n int) *AddPrincipalAttributes {
	apatts.allow_proxiable = n
	apatts.appendFlag(n, "allow_proxiable")

	return apatts
}

// -allow_dup_skey(0) disables user-to-user authentication for this principal by prohibiting this principal from obtaining a session key for another user. +allow_dup_skey (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetDupKey(n int) *AddPrincipalAttributes {
	apatts.allow_dup_key = n
	apatts.appendFlag(n, "allow_dup_skey")
	return apatts
}

// +requires_preauth(1) requires this principal to preauthenticate before being allowed to kinit. -requires_preauth(0) clears this flag. When +requires_preauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using preauthentication.
func (apatts *AddPrincipalAttributes) SetPreAuth(n int) *AddPrincipalAttributes {
	apatts.requires_preauth = n
	apatts.appendFlag(n, "requires_preauth")
	return apatts
}

// +requires_hwauth(1) requires this principal to preauthenticate using a hardware device before being allowed to kinit. -requires_hwauth(0) clears this flag. When +requires_hwauth is set on a service principal, the KDC will only issue service tickets for that service principal if the client’s initial authentication was performed using a hardware device to preauthenticate.
func (apatts *AddPrincipalAttributes) SetHwAuth(n int) *AddPrincipalAttributes {
	apatts.requires_hwauth = n
	apatts.appendFlag(n, "requires_hwauth")
	return apatts
}

// +ok_as_delegate(1) sets the okay as delegate flag on tickets issued with this principal as the service. Clients may use this flag as a hint that credentials should be delegated when authenticating to the service. -ok_as_delegate(0) clears this flag.
func (apatts *AddPrincipalAttributes) SetOkAsDelegate(n int) *AddPrincipalAttributes {
	apatts.ok_as_delegate = n
	apatts.appendFlag(n, "ok_as_delegate")
	return apatts
}

// -allow_svr(0) prohibits the issuance of service tickets for this principal. +allow_svr clears this flag(1).
func (apatts *AddPrincipalAttributes) SetSvr(n int) *AddPrincipalAttributes {
	apatts.allow_svr = n
	apatts.appendFlag(n, "allow_svr")
	return apatts
}

// -allow_tgs_req(0) specifies that a Ticket-Granting Service (TGS) request for a service ticket for this principal is not permitted. +allow_tgs_req(1) clears this flag.
func (apatts *AddPrincipalAttributes) SetTgsReq(n int) *AddPrincipalAttributes {
	apatts.allow_tgs_req = n
	apatts.appendFlag(n, "allow_tgs_req")

	return apatts
}

// -allow_tix(0) forbids the issuance of any tickets for this principal. +allow_tix (1) clears this flag.
func (apatts *AddPrincipalAttributes) SetTix(n int) *AddPrincipalAttributes {
	apatts.allow_tix = n
	apatts.appendFlag(n, "allow_tix")

	return apatts
}

// +needchange(1) forces a password change on the next initial authentication to this principal. -needchange(0) clears this flag.
func (apatts *AddPrincipalAttributes) SetNeedChange(n int) *AddPrincipalAttributes {
	apatts.needchange = n
	apatts.appendFlag(n, "needchange")

	return apatts
}

// +password_changing_service(1) marks this principal as a password change service principal.
func (apatts *AddPrincipalAttributes) SetPasswordChangingService() *AddPrincipalAttributes {
	apatts.password_changing_service = 1
	apatts.appendFlag(1, "password_changing_service")

	return apatts
}

// +ok_to_auth_as_delegate allows this principal to acquire forwardable tickets to itself from arbitrary users, for use with constrained delegation.
func (apatts *AddPrincipalAttributes) SetOkToAuthAsDelegate() *AddPrincipalAttributes {
	apatts.ok_to_auth_as_delegate = 1
	apatts.appendFlag(1, "ok_to_auth_as_delegate")

	return apatts
}

// +no_auth_data_required prevents PAC or AD-SIGNEDPATH data from being added to service tickets for the principal.
func (apatts *AddPrincipalAttributes) SetNoAuthDataRequired() *AddPrincipalAttributes {
	apatts.no_auth_data_required = 1
	apatts.appendFlag(1, "no_auth_data_required")

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
