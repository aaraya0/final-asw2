import React, { useState } from "react";
import "./Login.css";
import Cookies from "universal-cookie";

const Cookie = new Cookies();
const URL = "http://localhost:9000";

async function login(username, password) {
  return await fetch(`${URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"username": username, "password":password})
  })
    .then(response => {
      if (response.status === 400 || response.status === 401)
      {
        return {"user_id": -1}
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("user_id", response.user_id, {path: '/'})
      Cookie.set("username", username, {path: '/login'})
    })
}

function goto(path){
  window.location = window.location.origin + path
}

function Login() {
  const [errorMessages, setErrorMessages] = useState({});
  const error = "Contraseña o Usuario invalido";

  const handleSubmit = (event) => {
    event.preventDefault();
    var { uname, pass } = document.forms[0];

    login(uname.value, pass.value).then((data) => {
      if (Cookie.get("user_id") > -1) {
        goto("/home");
      } else if (Cookie.get("user_id") === -1) {
        setErrorMessages({ name: "default", message: error });
      }
    });
  };

  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

    return (
        <div className="form2">
    <form onSubmit={handleSubmit}>
    <div className="input-container">
    <label>Usuario </label>
    <input type="text" name="uname" required />
    {renderErrorMessage("uname")}
    </div>
    <div className="input-container">
    <label>Contraseña </label>
    <input type="password" name="pass" required />
    {renderErrorMessage("pass")}
    </div>
    <div className="button-container">
    <input type="submit" value="Iniciar Sesión"/>
    </div>
    </form>
    <div id="registerlink">
    <button id="register" onClick={() => goto("/register")}>
        Register
      </button>
      </div>
    </div>
      );
    }
    
    export default Login;