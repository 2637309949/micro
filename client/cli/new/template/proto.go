package template

var (
	ProtoServiceSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

import "proto/{{dehyphen .Alias}}/{{dehyphen .Alias}}.proto";

service {{title .Alias}} {
	rpc Call(Request) returns (Response) {}
}
`
	ProtoModelSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

message Request {
	string name = 1 [json_name = "name"];
    uint64 operator_id = 2 [json_name = "operator_id"];
    string operator_name = 3 [json_name = "operator_name"];
}

message Response {
	string msg = 1;
}
`
)
