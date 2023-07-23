import React, { useEffect, useState } from "react";
import Cookies from "universal-cookie";

const cookies = new Cookies();
const URL = "http://localhost:9000"; // Cambia la URL a la del backend correspondiente
const URL2 = "http://localhost:8090";

async function getUserById(id) {
  return await fetch(`${URL}/users/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "user/json",
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

function Account() {
  const [userData, setUserData] = useState(null);
  const [userItems, setUserItems] = useState(null);

  useEffect(() => {
    // Obtén el user_id del usuario desde las cookies
    const userId = Number(cookies.get("user_id"));

    // Recupera los datos del usuario desde el backend
    getUserData(userId);

    // Recupera las publicaciones del usuario desde el backend
    getUserItems(userId).then((items) => setUserItems(items));
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
    } catch (error) {
      console.error("Error deleting item:", error);
    }
    window.location.reload()
    alert("Publicacion eliminada con exito")
    
  }

  return (
    <div>
      {userData ? (
        <div>
          <h2>Información del Usuario</h2>
          <p>Nombre de Usuario: {userData.username}</p>
          <p>Nombre: {userData.first_name}</p>
          <p>Apellido: {userData.last_name}</p>
          <p>Email: {userData.email}</p>
        </div>
      ) : (
        <p>Cargando...</p>
      )}

      {userItems ? (
        <div>
          <h2>Mis publicaciones</h2>
          {userItems.map((item) => (
            <div key={item._id}>
              <p>Título: {item.title}</p>
              <p>Ubicación: {item.location}</p>
              <p>Descripción: {item.description}</p>
              {/* Botón de eliminar */}
              <button onClick={() => handleDeleteItem(item._id)}>Eliminar</button>
            </div>
          ))}
        </div>
      ) : (
        <p>Cargando publicaciones...</p>
      )}
    </div>
  );
}

export default Account;
