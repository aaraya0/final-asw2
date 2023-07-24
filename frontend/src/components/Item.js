import React, { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import "./Item.css" 

const cookies = new Cookies();
const URL_MESSAGES = "http://localhost:9001";
const URL_USERS = "http://localhost:9000";

async function getMessagesByItemId(itemId) {
  return await fetch(`${URL_MESSAGES}/items/${itemId}/messages`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}

async function getUserById(userId) {
  return await fetch(`${URL_USERS}/users/${userId}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}

function Item() {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    // Obtén el item_id desde la cookie "user_item"
    const itemId = cookies.get("user_item");

    // Recupera los mensajes para el item desde el backend
    getMessagesByItemId(itemId).then((messages) => {
      // Obtén los datos del usuario para cada mensaje y agrega el nombre, apellido y email
      const updatedMessages = messages.map((message) => getUserDataForMessage(message));
      
      Promise.all(updatedMessages).then((resolvedMessages) => {
        console.log("Mensajes enriquecidos con datos del usuario:", resolvedMessages);
        setMessages(resolvedMessages);
      });
    });
  }, []);

  async function getUserDataForMessage(message) {
    try {
      const userData = await getUserById(message.user_id);
      // Agrega los datos del usuario (nombre, apellido y email) al mensaje
      message.username = userData.username;
      message.last_name = userData.last_name;
      message.first_name = userData.first_name;
      message.email = userData.email;
      return message;
    } catch (error) {
      console.error("Error fetching user data for message:", error);
      return message;
    }
  }

  return (
    <div >
      <h2 id="titulo-msj">Mensajes para la publicación</h2>
      {messages.map((message) => (
        <div key={message.message_id} id="information">
         <div id="usuario"><img src="img/usuario.png"/>@{message.username} - {message.first_name} {message.last_name}</div>
          <div id="mail">{message.email}</div>
         <div id="mensaje"> <img src="img/mail.png"/>"{message.body}"</div>
          <div id="hora"> <img src="img/reloj.png"/>{message.created_at}</div>
        </div>
      ))}
    </div>
  );
}

export default Item;
