const net    = require('net');
const fs     = require('fs');
let   socket = new net.Socket();

socket.connect(
  3333,
  () => {
    console.log('Connected...');
  }
)
socket.on(
  'data',
  (data) => {
    getSongAndSave(String(data).split('|'));
  }
)

function getSongAndSave(data) {
  if(data[0] == '111') {
    let song = data.slice(4, data.length-3);
    let _song = `${song[0]}${((' / ' + song[1]) || ',')}${((' - ' + song[2]) || ',')}`;
    fs.writeFileSync('./np.txt', _song);
    console.log('Now Playing: ', _song);
  }
}
