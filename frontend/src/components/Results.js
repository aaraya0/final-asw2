import React, { useState, useEffect } from "react";
import "./Results.css";
import Cookies from "universal-cookie";

const cookies = new Cookies();

async function getItems() {
  let query = cookies.get("busqueda_limpia");
  return await fetch("http://localhost:8060/searchAll=" + query).then((response) => response.json());
}

function goto(path) {
  window.location = window.location.origin + path;
}

function retry() {
  goto("/home");
}

function parseField(field) {
  if (field !== undefined) {
    return field;
  }
  return "Not available";
}

function showItems(items) {
  
  return items.map((item) => (
    <div obj={item} key={item._id} className="item" onClick={() => goto("/info")}>
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
  ));
}

function Results() {
  const [items, setItems] = useState([]);
  const [needItems, setNeedItems] = useState(true);

  useEffect(() => {
    if (!items.length && needItems) {
      getItems().then((response) => {
        setItems(response);
        setNeedItems(false);
      });
    }
  }, [items, needItems]);

  return (
    <div className="home">
      <div className="topnavHOME">
        <h1 id="titulogrande">Resultados para la búsqueda "{cookies.get("busqueda_limpia")}"</h1>
        {items.length === 0 && !needItems && (
          <div className="no_results">
            <p id="no_results">No hay resultados disponibles.</p>
            <div ><button onClick={retry} id="search_button" >Otra búsqueda</button></div>
          </div>
        )}
        {showItems(items)}
      </div>
    </div>
  );
}

export default Results;
