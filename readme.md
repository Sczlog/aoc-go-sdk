# AOC Go SDK

Golang SDK for AOC

- [source code](https://github.com/arcfraoffical/aoc-go-sdk)
- [releases](https://github.com/arcfraoffical/aoc-go-sdk/releases)

## Installation

```shell
go get github.com/arcfraofficial/aoc-go-sdk
```

## How to

> In the example, we use two libs [pointy](https://github.com/openlyinc/pointy) and [go-funk](https://github.com/thoas/go-funk)，pointy is used to create pointer，go-funk provide some，like `Map`、`Filter`、`Reduce` and etc.

### Create instance

#### Create `ApiClient` instance

```go
import (
	apiclient "github.com/arcfraoffical/aoc-go-sdk/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
transport := httptransport.New("192.168.36.133", "/v2/api", []string{"http"})
client := apiclient.New(transport, strfmt.Default)
```

> add InsecureSkipVerify to skip tls verify for self-signed SSL certificate

```go
import (
	apiclient "github.com/arcfraoffical/aoc-go-sdk/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)
tlsClient, err := httptransport.TLSClient(httptransport.TLSClientOptions{
	InsecureSkipVerify: true,
})
if err != nil {
	fmt.Print(err)
	return
}
transport := httptransport.NewWithClient("192.168.29.157", "/v2/api", []string{"https"}, tlsClient)
client := apiclient.New(transport, strfmt.Default)
```

### Send request

#### import proper `client` package

```go
import (
  	vm "github.com/arcfraoffical/aoc-go-sdk/client/vm"
)
```

#### Authentication

Use `NewWithUserConfig` to create an `ApiClient` with authentication information

```go
import (
	apiclient "github.com/arcfraoffical/aoc-go-sdk/client"
	"github.com/arcfraoffical/aoc-go-sdk/models"
)

client, err := apiclient.NewWithUserConfig(apiclient.ClientConfig{
	Host:     "localhost:8090",
	BasePath: "v2/api",
	Schemes:  []string{"http"},
}, apiclient.UserConfig{
	Name:     "Name",
	Password: "Password",
	Source:   models.UserSourceLOCAL,
})
```

Or add authentication information to `ApiClient` already created

```go
import (
  User "github.com/arcfraoffical/aoc-go-sdk/client/user"
)
loginParams := User.NewLoginParams()
loginParams.RequestBody = &models.LoginInput{
	Username: pointy.String("username"),
	Password: pointy.String("password"),
	Source:   models.NewUserSource(models.UserSourceLOCAL),
}
logRes, err := client.User.Login(loginParams)
if err != nil {
	return err
}
transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", *logRes.Payload.Data.Token)
```

#### List resource

```go
getVmParams := vm.NewGetVmsParams();
getVmParams.RequestBody = &models.GetVmsRequestBody{
	Where: &models.VMWhereInput{
		ID: pointy.String("vm_id"),
	},
}
vmsRes, err := client.VM.GetVms(getVmParams)
if err != nil {
	return err
}
vms := vmsRes.Payload
```

#### Update resource

> Update resource will generate related asynchronous tasks, when the asynchronous task is completed, it means the resource operation is completed and the data is updated.

```go
target_vm := vmsRes.Payload[0]
vmStartParams := vm.NewStartVMParams()
vmStartParams.RequestBody = &models.VMStartParams{
	Where: &models.VMWhereInput{
		ID: target_vm.ID,
	},
}
startRes, err := client.VM.StartVM(vmStartParams)
if err != nil {
	return err
}
```

> `WaitTask` will wait for the asynchronous task to complete, if the task fails or times out, an exception will be returned.

```go
task := *startRes.Payload[0].TaskID
err = utils.WaitTask(context.TODO(), client, task, 1*time.Second)
if err != nil {
	return err
}
```

> For multiple tasks, you can use `WaitTasks`, which accepts multiple task ids, the rest is the same as `WaitTask`.

```go
tasks := funk.Map(startRes.Payload, func(tvm *models.WithTaskVM) string {
	return *tvm.TaskID
}).([]string)
err = utils.WaitTask(context.TODO(), client, tasks, 1*time.Second)
if err != nil {
	return err
}
```
