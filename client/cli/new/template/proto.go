package template

var (
	ProtoServiceSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

import "proto/{{dehyphen .Alias}}/{{dehyphen .Alias}}.proto";

service {{title .Alias}} {
    rpc InsertInfo (InsertInfoRequest) returns (InsertInfoResponse) {}
	rpc DeleteInfo (DeleteInfoRequest) returns (DeleteInfoResponse) {}
    rpc UpdateInfo (UpdateInfoRequest) returns (UpdateInfoResponse) {}
    rpc QueryInfo (QueryInfoRequest) returns (QueryInfoResponse) {}
    rpc QueryInfoDetail (QueryInfoDetailRequest) returns (QueryInfoDetailResponse) {}
}
`
	ProtoModelSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

message InsertInfoRequest {
	string name = 1 [json_name = "name"];
}

message InsertInfoResponse {
	uint32 id = 1 [json_name = "id"];
}

message DeleteInfoRequest {
	uint32 id = 1 [json_name = "id"];
}

message DeleteInfoResponse {
	uint32 id = 1 [json_name = "id"];
}

message UpdateInfoRequest {
	uint32 id = 1 [json_name = "id"];
	string name = 2 [json_name = "name"];
}

message UpdateInfoResponse {
	uint32 id = 1 [json_name = "id"];
}

message QueryInfoRequest {
	uint32 id = 1 [json_name = "id"];
	string name = 2 [json_name = "name"];
	int32 page = 3 [json_name = "page"];
	int32 size = 4 [json_name = "size"];
	int32 order_type = 5 [json_name = "order_type"];
	string order_col = 6 [json_name = "order_col"];
}

message QueryInfoResponseItem {
	uint32 id = 1 [json_name = "id"];
	string name = 2 [json_name = "name"];
}

message QueryInfoResponse {
    repeated QueryInfoResponseItem data = 1 [json_name = "data"];

	int32 page = 2 [json_name = "page"];
	int32 size = 3 [json_name = "size"];
	int32 total_count = 4 [json_name = "total_count"];
}

message QueryInfoDetailRequest {
	uint32 id = 1 [json_name = "id"];
}

message QueryInfoDetailResponse {
	uint32 id = 1 [json_name = "id"];
	string name = 2 [json_name = "name"];
}
`

	EnumSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

enum InfoType {
	None = 0;
}
`
)
