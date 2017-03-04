Small now playing output to file in node.js

In `Tools > Control Server` set `Output Options > Fields` to `$if(%album artist%,%album artist%,$if(%artist%,%artist%))$if(%album%,|%album%|)%title%|%codec%|%bitrate%` and turn on UTF-8 Output/Input

Know issues:
 * When song doesn't heve Artist and/or Album name will return `undefined` in song's name.
