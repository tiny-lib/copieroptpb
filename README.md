# copieroptpb
[copier](https://github.com/jinzhu/copier) option for protobuf wrappervalues
```go
user:=&User{}
if err := copier.CopyWithOption(req.User, user, copieroptpb.Option()); err != nil {
		return nil, err
	}
```
