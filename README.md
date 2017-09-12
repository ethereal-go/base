


```


import (
	"github.com/ethereal-go/ethereal"
	"github.com/ethereal-go/base"
)


	func main() {

	ethereal.Queries().Add("users", &base.UserField).Add("roles", &base.RoleField)
	ethereal.Mutations().Add("createUsers", &base.CreateUser)
	}	
```
