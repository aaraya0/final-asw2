import {MESSAGESHOST, USERSHOST, MESSAGESPORT, USERSPORT} from "./config/config";

const URL = MESSAGESHOST + ":" + MESSAGESPORT
const URLUSERS = USERSHOST + ":" + USERSPORT

async function getUserById(id){
  return await fetch(`${URLUSERS}/users/` + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  }).then(response => response.json())
}

async function getMessages() {
  return await fetch(URL + "/messages", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getUsersFirstNames(comments) {
  let user_first_name = {}
  for(let i = 0; i < comments.length; i++){
    let user = await getUserById(comments[i].user_id)
    user_first_name[user.user_id] = user.first_name;
  }
  return user_first_name
}

export const getComments = async () => {
  return await getMessages().then( async (response) => {
    let users = await getUsersFirstNames(response)
    for(let i = 0; i < response.length; i++){
      response[i].first_name = users[response[i].user_id]
      response[i].created_at = parseDate(response[i].created_at)
    }
    return response
  });
}


function parseDate(date){

  let dateParts = date.split("-");
  return `${dateParts[0]} ${dateParts[1] - 1} ${dateParts[2].substring(0, 2)}`;

}

export const createComment = async (text, uid, itemid) => { //cambiar por un POST

  return await fetch(URL + "/message", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body:JSON.stringify({
      user_id: uid,
      body: text,
      item_id: itemid,
      system: false,
    })
  }).then(response => response.json()).then(
      async (response) => {
        let user = await getUserById(response.user_id)
        response.first_name = user.first_name
        response.created_at = parseDate(response.created_at)
        return response
      }
  )
  
};