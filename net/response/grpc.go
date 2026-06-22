package response

import (
	"time"

	convert "github.com/dpanel-cn/common/utils/convert"
	"google.golang.org/protobuf/types/known/structpb"
)

// GrpcResponse 定义主控/被控交互的统一 gRPC 响应结构。
type GrpcResponse struct {
	RequestID string
	Cmd       string
	Ok        bool
	Data      map[string]any
	Error     string
	Timestamp int64
}

// GrpcOk 返回统一成功 gRPC 响应帧。
func GrpcOk(requestID string, cmd string, data map[string]any) *structpb.Struct {
	return GrpcResponseToStruct(GrpcResponse{
		RequestID: requestID,
		Cmd:       cmd,
		Ok:        true,
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}

// GrpcErr 返回统一失败 gRPC 响应帧。
func GrpcErr(requestID string, cmd string, message string) *structpb.Struct {
	return GrpcResponseToStruct(GrpcResponse{
		RequestID: requestID,
		Cmd:       cmd,
		Ok:        false,
		Error:     message,
		Timestamp: time.Now().Unix(),
	})
}

// GrpcResponseToStruct 将统一响应结构转换为 gRPC Struct 帧。
func GrpcResponseToStruct(resp GrpcResponse) *structpb.Struct {
	payload := map[string]any{
		"ok": resp.Ok,
		"ts": resp.Timestamp,
	}
	if resp.RequestID != "" {
		payload["request_id"] = resp.RequestID
	}
	if resp.Cmd != "" {
		payload["cmd"] = resp.Cmd
	}
	if resp.Data != nil {
		payload["data"] = resp.Data
	}
	if resp.Error != "" {
		payload["error"] = resp.Error
	}
	return convert.ToStruct(payload)
}
