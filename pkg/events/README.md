# events

Events is a package that implements a websocket handler to accept events and put those events into Dynamo.  This package was created as a sample to illustrate using Golang interfaces for testing.  Specifically for testing network or 3rd party service errors.  In this case we use a "stub" dynamo package that implements the DynamoDBAPI interface that always returns an error.

Testing network or database responses can be challenging.  By using Go interfaces you can simulate whatever responses you would like.  The main method by which this is accomplished in events is by using the DynamoDBAPI interface to interact with vs using the dynamodb.DynamoDB to interact with.  

Specifically this is done in the line: 
```go
var Dynamo dynamodbinterface.DynamoDBAPI
```

By doing so we can use our stub implementation (that always returns an error) in testing and in production use an actual dynamodb.DynamoDB instance.  This allows us to simulate in testing any response we want from Dynamo (in this case an error) and ensure that the program acts accordingly.
