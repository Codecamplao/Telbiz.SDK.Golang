# TelBiz.SDK.Golang

## Installation

Begin by installing this package through Go. Just run following command to terminal-

```shell script
go get github.com/Codecamplao/Telbiz.SDK.Golang
```

Next, Add more config on environment file. Open `.env`, and add a new three config like below.

```dotenv
TELBIZ_BASE_URI=""
TELBIZ_CLIENT_ID=""
TELBIZ_CLIENT_SECRET=""
```

## Example to use
```go
import (
	"fmt"
	"github.com/Codecamplao/Telbiz.SDK.Golang/telbiz"
)

func main() {
	sms := telbiz.SMSService(telbiz.Default, "20xxxxxxxx", "Hello Telbiz")
	fmt.Printf("SMS Response :: %+v\n", sms)
}
```


