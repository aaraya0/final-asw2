import React, { useState, useEffect } from "react";
import "./Message.css";
import Cookies from "universal-cookie";



const cookies = new Cookies();
const URL = "http://localhost:9001"
const defaultMessage = "Hola, estoy interesado en obtener más información sobre esta propiedad.";
async function getItem() {
    var id = cookies.get("item_id")
    return await fetch("http://localhost:8090/items/" + id).then((response) => response.json());
  }

  
  function parseField(field) {
    if (field !== undefined) {
      return field;
    }
    return "Not available";
  }
  function showItem(item) {
  
    return (
      <div
      obj={item}
      key={item._id}
      className="item2" >
        <img id="imagen" src={`img/${item._id}.png`} alt={`${item._id}`}/>
        <a id="tituloitem">{parseField(item.title)}</a>
  
        <div id="info">
          <div id="location">{parseField(item.location)}</div>
          <div id="description">{parseField(item.description)}</div>
          <div id="price">$ {parseField(item.price)}</div>
          <div id="mts"><img src="img/metro-cuadrado.png"></img>{parseField(item.mts2)} mts²</div>
          <div id="class">{parseField(item.class)}</div>
        </div>
      </div>
    );
  }
async function sendMessage(user_id, body, item_id,system) {
  return await fetch(`${URL}/message`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({
      "user_id": user_id,
      "body":body,
      "item_id":item_id,
      "system":system
    })
  })
  
}

function goto(path){
  window.location = window.location.origin + path
}

function Message() {

    const [item, setItem] = useState([]);
    const [needItem, setNeedItem] = useState(true);
  
    useEffect(() => {
      if (!item.length && needItem) {
        getItem().then((response) => {
            
          setItem(response);

          setNeedItem(false);
        });
      }
    }, [item, needItem]);

  const handleSubmit = (event) => {
    //Prevent page reload
    event.preventDefault();

    var {body} = document.forms[0];
    var user_id = parseInt(cookies.get("user_id"))
    var item_id = cookies.get("item_id")
    sendMessage(user_id, body.value, item_id, false)
    .then((response) => {
      if (response.status === 201) {
        // Si el código de respuesta es 201, significa que la solicitud fue exitosa
        alert("Mensaje enviado");
        goto("/results");
      } else {
        // Si el código de respuesta no es 201, muestra una alerta con el mensaje de error
        alert("Error al enviar el mensaje: " + response.statusText);
      }
    })
    .catch((error) => {
      // Si ocurre un error en la solicitud, también muestra una alerta con el mensaje de error
      alert("Error al enviar el mensaje: " + error.message);
    });
  
  };
  // JSX code for login form
  const renderForm = (
    <div className="message-form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label className="label">Contactate por la propiedad enviando un mensaje:</label>
          <input type="text"  name="body" className="text-area"  defaultValue={defaultMessage}/>
        </div>
        <div className="button-container">
          <input type="submit" value="Enviar mensaje"/>
        </div>
      </form>
      
    </div>
  );


  return (
    <div>
    <div className="home">
      <div className="topnavHOME">
    
      </div>
    </div>

    <div className="app">
        <div className="item-info">
            {showItem(item)}
        </div>
      <div className="message-form">
      
        {renderForm}
      </div>
    </div>
    </div>
  );
}

export default Message;