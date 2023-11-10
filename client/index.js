import WebSocket from "ws"
import axios from "axios"

async function main() {
  const registerPlayerRes = await registerPlayer("player 1")
  console.log(registerPlayerRes)
  const registerMatchRes = await registerMatch(registerPlayerRes.sessionToken, 3)
  console.log(registerMatchRes)
  const registerPlayer2Res = await registerPlayer("player 2")
  console.log(registerPlayer2Res)
  const joinMatchRes = await joinMatch(registerPlayer2Res.sessionToken, registerMatchRes.id)
  console.log(joinMatchRes)
  const startMatchRes = await startMatch(registerPlayerRes.sessionToken, registerMatchRes.id)
  console.log(startMatchRes)
  // const registerPlayerRes = { sessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQbGF5ZXIiOnsiaWQiOiJtMmVcdTAwMjZ6IiwibmFtZSI6InBsYXllciAxIn0sImlzcyI6InJvY2stcGFwZXItc2Npc3NvcnMtc2VydmVyIiwiZXhwIjoxNzAwODQ3MDQ4LCJuYmYiOjE2OTk2Mzc0NDgsImlhdCI6MTY5OTYzNzQ0OH0.cZ_FclzjqoHuedLBCJol6wbcNGuwwOeFhPsjBAnsZ6w" }
  // const registerPlayer2Res = { sessionToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQbGF5ZXIiOnsiaWQiOiJ4NEBnKiIsIm5hbWUiOiJwbGF5ZXIgMiJ9LCJpc3MiOiJyb2NrLXBhcGVyLXNjaXNzb3JzLXNlcnZlciIsImV4cCI6MTcwMDg0NzA0OCwibmJmIjoxNjk5NjM3NDQ4LCJpYXQiOjE2OTk2Mzc0NDh9._rmDYPpsCJLiFmTBKy-lq9a3wcrDaXOWRsOjUNv__5A" }
  // const registerMatchRes = { id: "kto*w" }

  try {
    const ws = new WebSocket(`ws://127.0.0.1:3001/ws/${registerMatchRes.id}`, "rockpaperscissors", { headers: { "Authorization": registerPlayerRes.sessionToken } });
    ws.on('open', () => {
      ws.send(JSON.stringify({ "move": 0 }))
    })
    ws.on('message', function message(data) {
      console.log('received p1: %s', data);
    });

    ws.on('error', console.error);
  } catch (error) {
    console.error(error)
  }

  try {
    const ws = new WebSocket(`ws://127.0.0.1:3001/ws/${registerMatchRes.id}`, "rockpaperscissors", { headers: { "Authorization": registerPlayer2Res.sessionToken } });
    ws.on('open', () => {
      ws.send(JSON.stringify({ "move": 1 }))
    })
    ws.on('message', function message(data) {
      console.log('received p2: %s', data);
    });

    ws.on('error', console.error);
  } catch (error) {
    console.error(error)
  }
}
main()


async function registerPlayer(name) {
  const res = await axios.post('http://127.0.0.1:3001/player', { name })
  const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
  console.log('Status Code:', res.status);
  console.log('Date in Response header:', headerDate);
  return res.data
}

async function registerMatch(token, maxScore) {
  const res = await axios.post('http://127.0.0.1:3001/match', { maxScore }, { headers: { "Authorization": token } })
  const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
  console.log('Status Code:', res.status);
  console.log('Date in Response header:', headerDate);
  return res.data
}

async function joinMatch(token, matchId) {
  const res = await axios.post(`http://127.0.0.1:3001/match/${matchId}/join`, {}, { headers: { "Authorization": token } })
  const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
  console.log('Status Code:', res.status);
  console.log('Date in Response header:', headerDate);
  return res.data
}

async function startMatch(token, matchId) {
  const res = await axios.post(`http://127.0.0.1:3001/match/${matchId}/start`, {}, { headers: { "Authorization": token } })
  const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
  console.log('Status Code:', res.status);
  console.log('Date in Response header:', headerDate);
  return res.data
}
