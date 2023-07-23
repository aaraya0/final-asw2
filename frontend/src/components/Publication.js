import React, { useState } from "react";
import Cookies from "universal-cookie";
import "./Publication.css";

const cookies = new Cookies();
const URLITEMS = "http://localhost:8090";
const URLUSERS = "http://localhost:9000";

async function getUserById(id) {
  return fetch(`${URLUSERS}/users/` + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}

async function getItemById(id) {
  return fetch(`${URLITEMS}/item/` + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((response) => response.json());
}
function Publication() {
    const [jsonData, setJsonData] = useState(null);
    const [isPosted, setIsPosted] = useState(false);
  
  
    const handleFileChange = async (event) => {
      const file = event.target.files[0];
      const reader = new FileReader();
  
      reader.onload = (event) => {
        const content = event.target.result;
        try {
          let parsedJson = JSON.parse(content);
  
          // If parsedJson is not an array, wrap it in an array
          if (!Array.isArray(parsedJson)) {
            parsedJson = [parsedJson];
          }
  
          // Add usuario_id to each item in the array
          const userId = cookies.get("user_id");
          parsedJson.forEach((item) => {
            item.usuario_id = Number(userId);
          });
  
          setJsonData(parsedJson);
        } catch (error) {
          console.error("Error parsing JSON:", error);
          setJsonData(null);
        }
      };
  
      reader.readAsText(file);
    };
  
    const handlePublish = async () => {
      try {
        const response = await fetch(`${URLITEMS}/items`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(jsonData),
        });
  
        if (!response.ok) {
          throw new Error("Failed to publish items");
        }
        if (response.ok) {
            setIsPosted(true);
          }
      } catch (error) {
        console.error("Error publishing items:", error);
      }
    
    };
    const handleNewPublication = () => {
        setJsonData(null);
        setIsPosted(false);
      };
    
      return (
        <div className>
          <input type="file" className="input-publ" onChange={handleFileChange} />
          {/* Display JSON data or error message */}
          {jsonData ? (
            <>
              <pre>{JSON.stringify(jsonData, null, 2)}</pre>
              <button onClick={handlePublish}>Publicar</button>
            </>
          ) : (
            <p className="msj">Selecciona un archivo JSON válido</p>
          )}
    
          {isPosted && (
            <div>
              <p className="msj">¡Publicación realizada con éxito!</p>
            </div>
          )}
        </div>
      );
    }
    
    export default Publication;