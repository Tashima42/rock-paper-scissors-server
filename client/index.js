import WebSocket from "ws"
import axios from "axios"

function main() {
  const { player, token } = registerPlayer("pedro")

  try {
    const ws = new WebSocket("ws://127.0.0.1:3001/ws/1");
    ws.on('open', () => {
      ws.send("TEST DATA")
    })

    ws.on('message', function message(data) {
      console.log('received: %s', data);
    });

    ws.on('error', console.error);
  } catch (error) {
    console.error(error)
  }
}
main()


function registerPlayer(name) {
  axios.post('http://127.0.0.1:3001/player', { name })
    .then(res => {
      const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
      console.log('Status Code:', res.status);
      console.log('Date in Response header:', headerDate);
      return res.data
    })
    .catch(err => {
      console.log('Error: ', err.message);
    });

}