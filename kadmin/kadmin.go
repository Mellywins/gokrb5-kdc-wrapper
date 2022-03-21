package kadmin

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/Mellywins/gokrb5-kdc-wrapper/internal/types"
)

type ExecutorSpec struct {
	Local              bool
	Realm              string // Defaults to the KDC server default realm.
	Principal          string
	UseKeytab          bool
	keytab             string
	CredentialsCache   string
	password           string
	query              types.Query
	Dbname             string
	AdminServer        string // admin_server[:port]
	Salt               string // TODO
	AuthGssapi         bool
	AuthGssapiFallback bool
	Verbose            bool
}

func (ex *ExecutorSpec) NewKadminExecutor() *ExecutorSpec {
	return &ExecutorSpec{
		Local:              false,
		Realm:              "",
		Principal:          "",
		UseKeytab:          false,
		keytab:             "",
		CredentialsCache:   "",
		password:           "",
		query:              nil,
		Dbname:             "",
		AdminServer:        "",
		Salt:               "",
		AuthGssapi:         false,
		AuthGssapiFallback: false,
	}
}

// ProbeLiveliness TODO
func (ex *ExecutorSpec) ProbeLiveliness() *ExecutorSpec {
	return ex
}

// ExecutorSpecBuilder  builder pattern code
type ExecutorSpecBuilder struct {
	executorSpec *ExecutorSpec
}

// NewExecutorSpecBuilder Incrementation creation of the kadmin Executor
func NewExecutorSpecBuilder() *ExecutorSpecBuilder {
	executorSpec := &ExecutorSpec{}
	b := &ExecutorSpecBuilder{executorSpec: executorSpec}
	return b
}

// MakeVerbose Enable the logging of the commands to the console.
// TODO: Hide sensitive DATA
func (b *ExecutorSpecBuilder) MakeVerbose(verb bool) *ExecutorSpecBuilder {
	b.executorSpec.Verbose = verb
	return b
}

// Local runs the Queries with the kadmin.local.
func (b *ExecutorSpecBuilder) Local(local bool) *ExecutorSpecBuilder {
	b.executorSpec.Local = local
	return b
}

func (b *ExecutorSpecBuilder) Realm(realm string) *ExecutorSpecBuilder {
	b.executorSpec.Realm = realm
	return b
}

func (b *ExecutorSpecBuilder) Principal(principal string) *ExecutorSpecBuilder {
	b.executorSpec.Principal = principal
	return b
}

func (b *ExecutorSpecBuilder) UseKeytab(useKeytab bool) *ExecutorSpecBuilder {
	b.executorSpec.UseKeytab = useKeytab
	return b
}

func (b *ExecutorSpecBuilder) keytab(keytab string) *ExecutorSpecBuilder {
	b.executorSpec.keytab = keytab
	return b
}

func (b *ExecutorSpecBuilder) CredentialsCache(credentialsCache string) *ExecutorSpecBuilder {
	b.executorSpec.CredentialsCache = credentialsCache
	return b
}

func (b *ExecutorSpecBuilder) password(password string) *ExecutorSpecBuilder {
	b.executorSpec.password = password
	return b
}

func (b *ExecutorSpecBuilder) query(query types.Query) *ExecutorSpecBuilder {
	b.executorSpec.query = query
	return b
}

func (b *ExecutorSpecBuilder) Dbname(dbname string) *ExecutorSpecBuilder {
	b.executorSpec.Dbname = dbname
	return b
}

func (b *ExecutorSpecBuilder) AdminServer(adminServer string) *ExecutorSpecBuilder {
	b.executorSpec.AdminServer = adminServer
	return b
}

func (b *ExecutorSpecBuilder) Salt(salt string) *ExecutorSpecBuilder {
	b.executorSpec.Salt = salt
	return b
}

func (b *ExecutorSpecBuilder) AuthGssapi(authGssapi bool) *ExecutorSpecBuilder {
	b.executorSpec.AuthGssapi = authGssapi
	return b
}

func (b *ExecutorSpecBuilder) AuthGssapiFallback(authGssapiFallback bool) *ExecutorSpecBuilder {
	b.executorSpec.AuthGssapiFallback = authGssapiFallback
	return b
}

// Build returns a reference to the ExecutorSpec, and error if it occurs
func (b *ExecutorSpecBuilder) Build() (*ExecutorSpec, error) {
	return b.executorSpec, nil
}

// Execute Accepts a type of interface Query. It will then formulate the shell script that will be run on the KDC and run it.
func (b *ExecutorSpec) Execute(command types.Query) *exec.Cmd {
	commandString := fmt.Sprintf("%s %s", "kadmin.local -q ", command.Exec())
	if b.Verbose {
		fmt.Println(commandString)
	}
	cmd := exec.Command("kadmin.local", "-q", command.Exec())
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Execution result: %s\n", out.String())
	return cmd
}
