Small script outputing song played in foobar2000 to file.

Dependency:
[foo_controlserver](https://github.com/audiohead/foo_controlserver)

In `Tools > Control Server` set `Output Options > Fields` to

`{"lengthSecound": "%length_seconds%","codec": "%codec%","bitrate": "%bitrate%","artist": "$if(%album artist%,%album artist%,%artist%)","album": "%album%","date": "%date%","genre": "%genre%","trackNummber": "%tracknumber%","title": "%title%"}`

and turn on UTF-8 Output/Input.

Node.js version: run `node index.js`

Python 3 version: WIP
