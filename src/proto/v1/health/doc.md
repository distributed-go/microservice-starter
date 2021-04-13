# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [health.proto](#health.proto)
    - [Health](#v1.healthpb.Health)
    - [InboundConnection](#v1.healthpb.InboundConnection)
    - [OutboundConnection](#v1.healthpb.OutboundConnection)
  
    - [ConnectionStatus](#v1.healthpb.ConnectionStatus)
    - [ServiceStatus](#v1.healthpb.ServiceStatus)
  
- [Scalar Value Types](#scalar-value-types)



<a name="health.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## health.proto



<a name="v1.healthpb.Health"></a>

### Health



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timestamp_utc | [string](#string) |  | current timestamp in utc |
| service_name | [string](#string) |  | service name |
| service_provider | [string](#string) |  | service provider name |
| service_version | [string](#string) |  | service version |
| service_status | [ServiceStatus](#v1.healthpb.ServiceStatus) |  | service status |
| service_start_time_utc | [string](#string) |  | service start time in utc |
| uptime | [double](#double) |  | service uptime in utc |
| inbound_connections | [InboundConnection](#v1.healthpb.InboundConnection) | repeated | inbound connections list |
| outbound_connections | [OutboundConnection](#v1.healthpb.OutboundConnection) | repeated | outbound connections list |






<a name="v1.healthpb.InboundConnection"></a>

### InboundConnection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_name | [string](#string) |  | name of the application |
| connection_status | [ConnectionStatus](#v1.healthpb.ConnectionStatus) |  | connectins status of the application |
| timestamp_utc | [string](#string) |  | current timestamp in utc |
| hostname | [string](#string) |  | hostname |
| address | [string](#string) |  | ip address of the application |
| os | [string](#string) |  | OS |






<a name="v1.healthpb.OutboundConnection"></a>

### OutboundConnection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| application_name | [string](#string) |  | applcition name |
| timestamp_utc | [string](#string) |  | current timestamp in utc |
| urls | [string](#string) | repeated | connection urls |
| connection_status | [ConnectionStatus](#v1.healthpb.ConnectionStatus) |  | connection status of application |





 


<a name="v1.healthpb.ConnectionStatus"></a>

### ConnectionStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| Disconnected | 0 | internet is disconneced |
| Active | 1 | internet is connected |



<a name="v1.healthpb.ServiceStatus"></a>

### ServiceStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| Stopped | 0 | service stopped |
| Running | 1 | service running |
| Degraded | 2 | service heath is degraded |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

