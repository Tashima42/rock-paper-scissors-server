import WebSocket from "ws"

function main() {
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