# TEAM 4

## Tasks
1. Consruct a Sphero client
2. Decrypt data
3. Convert the data into Sphero commands
4. Deliver the code to the release party
 
## Decrypting the data
The data received by pusher is encrypted to send it to sphero you must first decrypt it.  The encruption algorythim is ququite simple and listed below.  The format of the message is ...

```
([char number][char number]...)*[mesage id]
```

Therfore if the orginal value is `ABC` and the message_id is 2 the encrypted value would be:
`20406`

To decrypt this value we need to first divide it by the message id:
`20406/2 = 10203`

If the number is odd in length then we need to convert it to a string and pre-pend a `0` to the beginning.
`010203`

Next we break this into 2 character blocks convert them to a letter where 01 is A and 26 is Z
`A,B,C`

And reassemble them to form our decrypted data
`ABC`

## Robot commands
```go
// speed 0-100
// direction 0-359
driver.Roll(speed, direction)
driver.Stop()
```
