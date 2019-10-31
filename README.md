<p align="center">
  <img alt="MoDB" src="https://user-images.githubusercontent.com/1941100/67701630-31647580-f9a8-11e9-9617-22f0c9053fde.png?style=centerme">
</p>

In memory database built with Go

## Methods
    SET
        TBA
        
    GET
        TBA
        
    DELETE
        TBA
        
    EXPIRE
        TBA
        
    Close
        TBA

## REST API
    Url: http://127.0.0.1:34443/db
    
    SET
        HTTP Method: POST
        Payload: { "Method": "SET", "Key": "myKey", "Value": "myValue"}
        
    GET
        HTTP Method: POST
        Payload: { "Method": "GET", "Key": "myKey"}
        
    DELETE
        HTTP Method: POST
        Payload: { "Method": "DELETE", "Key": "myKey"}
        
## Performance

moDB was tested in a single machine. And the performance was executed to 
insert a certain amount of data and deleting those using the builtin methods. 
Results are displayed below:

<p align="center">
  <img alt="MoDB" src="https://user-images.githubusercontent.com/1941100/67943511-d2d70b80-fbe2-11e9-9f91-ff1a88fe73a8.png">
</p>
