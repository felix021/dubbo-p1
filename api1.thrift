namespace go echo1

include "java.thrift"

struct EchoRequest {
    1: required string message,
    2: optional java.Object obj,
}(JavaClassName="kitex.echo1.EchoRequest")

struct EchoResponse {
    1: required string message,
}(JavaClassName="kitex.echo1.EchoResponse")

service TestService1 {
    EchoResponse Echo(1: EchoRequest req)
} (JavaClassName="kitex.echo1.TestService1")
