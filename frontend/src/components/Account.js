import React, { useEffect, useState } from "react";

import "./Account.css";
import Cookies from "universal-cookie";

const cookies = new Cookies();
const URL = "http://localhost:9000"; // Cambia la URL a la del backend correspondiente
const URL2 = "http://localhost:8090";
const URL3 = "http://localhost:9001";

async function getUserById(id) {
  return await fetch(`${URL}/users/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "user/json",
    },
  }).then((response) => response.json());
}
async function deleteUser(userId) {
  return await fetch(`${URL}/user/${userId}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}
async function getUserItems(userId) {
  return await fetch(`${URL2}/users/${userId}/items`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}

async function deleteItem(itemId) {
  return await fetch(`${URL2}/item/${itemId}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}
async function getUserMessages(userId) {
  try {
    const response = await fetch(`${URL3}/users/${userId}/messages`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const messages = await response.json();
    return messages;
  } catch (error) {
    console.error("Error fetching user messages:", error);
    return [];
  }
}
function Account() {
  const [userData, setUserData] = useState(null);
  const [userItems, setUserItems] = useState(null);
  const [userMessages, setUserMessages] = useState([]);

  useEffect(() => {
    const userId = Number(cookies.get("user_id"));

    getUserData(userId);
    getUserItems(userId).then((items) => setUserItems(items));
    getUserMessages(userId).then((messages) => setUserMessages(messages));
  }, []);
  async function getUserData(userId) {
    try {
      const userData = await getUserById(userId);
      // Filtra los datos para excluir el user_id y la contraseña
      const { user_id, password, ...filteredData } = userData;
      setUserData(filteredData);
    } catch (error) {
      console.error("Error fetching user data:", error);
    }
  }

  async function handleDeleteItem(itemId) {
    try {
      await deleteItem(itemId);
      // Actualiza la lista de publicaciones después de eliminar la publicación
      const updatedItems = userItems.filter((item) => item._id !== itemId);
      setUserItems(updatedItems);
      alert("Publicacion eliminada con exito");
    } catch (error) {
      console.error("Error deleting item:", error);
    }
    window.location.reload();
  }

  function handleItemClick(itemId) {
    // Guardar el ID de la publicación en la cookie "user_item"
    cookies.set("user_item", itemId, { path: "/" });
    window.location.href = "/messages";
  }

  async function handleDeleteAccount() {
    try {
      // Redirigir a la página de inicio ("/")
      window.location.href = "/";
      
      const userId = cookies.get("user_id");
      await deleteUser(userId);
      // Borrar todas las cookies
      cookies.remove("user_id", { path: "/" });
      cookies.remove("user_item", { path: "/" });
      
    } catch (error) {
      console.error("Error deleting account:", error);
  
    }
  }
  return (
    <div>
      {userData ? (
        <div>
          <h2 id="titulo-info">Información del Usuario</h2>
          <img src="img/usuario.png" id="usuario-img" alt="Usuario"></img>
          <div id="user">@{userData.username}</div>
          <div id="nombre">
            {userData.first_name} {userData.last_name}
          </div>
          <div id="correo">{userData.email}</div>
        </div>
      ) : (
        <p>Cargando...</p>
      )}
  
      <div id="publicaciones">
        <h2 id="titulo-info">Mis publicaciones</h2>
        {userItems && userItems.length > 0 ? (
          userItems.map((item) => (
            <div key={item._id} id="item-info">
              <div id="texto">
                <p>Título: {item.title}</p>
                <p>Ubicación: {item.location}</p>
                <p>Descripción: {item.description}</p>
              </div>
              {/* Botón de eliminar */}
              <button id="boton-eliminar" onClick={() => handleDeleteItem(item._id)}>Eliminar</button>
              {/* Hacer clic en la publicación */}
              <button id="boton-detalle" onClick={() => handleItemClick(item._id)}>Mensajes</button>
            </div>
          ))
        ) : (
          <div id="cargando">
          <p>No hay publicaciones.</p>
          </div>
        )}
      </div>
            {/* Displaying user messages */}
            <div id="user-messages">
        <h2 id="titulo-info">Mis Mensajes</h2>
        {userMessages && userMessages.length > 0 ? (
          userMessages.map((message) => (
            <div key={message._id} id="message-info">
              <div id="message-body">{message.body}</div>
              <div id="item-id">Item ID: {message.item_id}</div>
              <div id="created-at">Fecha: {message.created_at}</div>
            </div>
          ))
        ) : (
          <div id="no-messages">
            <p>No hay mensajes.</p>
          </div>
        )}
      </div>
  
      {/* Botón para eliminar cuenta */}
      <button id="boton-eliminar-cuenta" onClick={handleDeleteAccount}>Eliminar cuenta</button>
    </div>
  );
  
        }
export default Account;