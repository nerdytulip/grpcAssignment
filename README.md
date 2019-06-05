# grpcAssignment
I have added the functionality to check whether a number is Armstrong or not , in the basic calculator grpc code

HOW TO RUN THE CODE :
1)run the .proto file that is there in the proto folder
![1](https://user-images.githubusercontent.com/40175918/58931445-577d4500-877d-11e9-9607-83343badca0a.png)

To run this .proto file the command used is:
protoc -I go/src/grpcAssignment/proto/ go/src/grpcAssignment/proto/*.proto -- go_out=plugins=grpc:go/src/grpcAssignment/calculatorpb/proto/proto

this creates the calclator.pb.go file 
2)Then we navigate to the grpcAssignment/ and Run the main.go file 

OUTPUT:



