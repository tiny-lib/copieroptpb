# copieroptpb
[copier](https://github.com/jinzhu/copier) option for protobuf wrappervalues
```go
import "github.com/tiny-lib/copieroptpb"
....
user:=&User{}
if err := copier.CopyWithOption(req.User, user, copieroptpb.Option()); err != nil {
    return nil, err
}
```
supported types:
+ [x] `wrappervalue*` to its related types
+ [x] `timestamppb` to `time.time`
