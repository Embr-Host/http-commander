# HTTP COMMANDER
 Webapp that can monitor and run a background process

 ## Setup

 ### Create a settings.json

 **`settings.json` file must be in the same location as application**

```
{
    "apiKey": "******************************************************",
    "commanderPort": "5000",
    "startCommand": "go run src/dummy-cmd.go",
    "noAuth": false
}
```

variable | description | required
--- | --- | ---
apiKey | use for authorization header | yes
commanderPort | port that webapp starts on | yes
startCommand | command will run in the background of the web app | yes
noAuth  | when set to `true` it will ignore the `apiKey` | no (default is false)

## API Routes

### GET `/status`

Checks if the command is still running in the background. If the command has stopped running it will send back a 500 status.

**Does not require authorization**

### POST `/send-command`

Takes the `command` param that is sent and sends it to the input buffer of the background task

parameter | description | required 
--- | --- | ---
command | Text that is sent to the input buffer of the background task | yes

### GET `/stop`

Stops the background task

### GET `/restart`

Stops the background task if not stopped already and then starts it again

### GET `/logging`

Creates a http event stream and will send the background command output buffer back

**Since this route requires authorization it may not work with the `EventSource` object**

## License

React is [MIT licensed](./LICENSE).