# learn-grpc
On the path of learning gRPC.

Part-1 - Implementing Unary Greet API Using gRPC

Part-2 - Implementation of  a Sum RPC Unary API in a CalculatorService:

The function takes a Request message that has two integers, and returns a Response that represents the sum of them.

Example:

The client will send two numbers (8 and 22) and the server will respond with (30)

Part-3 - Implementing Server Streaming

Part-4 - Implementation of a  PrimeNumberDecomposition RPC Server Streaming API in a CalculatorService:

The function takes a Request message that has one integer, and returns a stream of Responses that represent the prime number decomposition of that number.

Example:

The client will send one number (120) and the server will respond with a stream of (2,2,2,3,5), because 120=2*2*2*3*5

Part-5 - Implementing Client Streaming

Part-6 - Implementation of a ComputeAverage RPC Client Streaming API in a CalculatorService:

The function takes a stream of Request message that has one integer, and returns a Response with a double that represents the computed average

Example:

The client will send a stream of number (1,2,3,4) and the server will respond with (2.5), because (1+2+3+4)/4 = 2.5 

Part-7 - Implementing BiDi rpc 
