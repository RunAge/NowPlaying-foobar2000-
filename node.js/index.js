const net = require('net');
const fs = require('fs');
const socket = new net.Socket();

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
  if (data[0] == '111') {
    let json = JSON.parse(data[4]);
    let song = `${json.artist != '?' ? json.artist : ''}${json.album != '?' ? " [" + json.album + "]" : ''} - ${json.title != '?' ? json.title : ''}`
    fs.writeFileSync('./np.txt', song);
    console.log('Now Playing: ', song);
  }
}
