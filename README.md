Small script outputing song played in foobar2000 to file.

In `Tools > Control Server` set `Output Options > Fields` to `$if(%album artist%,%album artist%,$if(%artist%,%artist%))$if(%album%,|%album%|)%title%|%codec%|%bitrate%` and turn on UTF-8 Output/Input

Known issues:
 * When song doesn't have Artist and/or Album name, `undefined` will be return as song's name.
