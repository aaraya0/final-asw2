import React, { useState } from "react";
import "./Register.css";
import Cookies from "universal-cookie";

const Cookie = new Cookies();
const URL = "http://localhost:9000";

async function register(username, password, first_name, last_name, email) {
  return await fetch(`${URL}/user`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "username": username,
      "password": password,
      "first_name": first_name,
      "last_name": last_name,
      "email": email
    })
  })
    .then(response => {
      if (response.status === 400 || response.status === 401) {
        return {"user_id": -1};
      }
      return response.json();
    })
    .then(response => {
      Cookie.set("user_id", response.user_id, { path: '/' });
      Cookie.set("username", username, { path: '/login' });
    });
}

function goto(path) {
  window.location = window.location.origin + path;
}

function Register() {
  // React States
  const [errorMessages, setErrorMessages] = useState({});

  const error = "Las contraseñas deben coincidir";

  const handleSubmit = (event) => {
    // Prevent page reload
    event.preventDefault();
    var { uname, pass, cpass, first_name, last_name, email } = document.forms[0];

    // Validate user registration TODO: Check if the username is not taken
    register(uname.value, pass.value, first_name.value, last_name.value, email.value).then(data => {
      if (pass.value === cpass.value) {
        goto("/");
      } else {
        setErrorMessages({ name: "default", message: error });
      }
    });
  };

  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && <div id="error">{errorMessages.message}</div>;

  // JSX code for login form
  const renderForm = (
    <div id="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Usuario: </label>
          <input type="text" name="uname" id="uname" required />
        </div>
        <div className="input-container">
          <label>Contraseña: </label>
          <input type="password" name="pass" id="pass" required />
        </div>
        <div className="input-container">
          <label>Confirmar Contraseña: </label>
          <input type="password" name="cpass" id="cpass" required />
        </div>

        {renderErrorMessage("default")}

        <div className="input-container">
          <label>Nombre(s): </label>
          <input type="text" name="first_name" id="name" required />
        </div>
        <div className="input-container">
          <label>Apellido(s): </label>
          <input type="text" name="last_name" id="last-name" required />
        </div>
        <div className="input-container">
          <label>Email: </label>
          <input type="email" name="email" id="e-mail" required />
        </div>
        <div className="button-container">
          <input type="submit" id="guardar" value="Guardar Datos" />
        </div>
      </form>
    </div>
  );

  const logreg = (
    <div>
      <a id="login" onClick={() => goto("/login")}>
        Login
      </a>
      <a id="register" className="clicked" onClick={() => goto("/register")}>
        Register
      </a>
    </div>
  );

  return (
    <div>
      <div className="home">
        <div className="topnavHOME"></div>
      </div>

      <div className="app">
        <div className="login-form">{renderForm}</div>
      </div>
    </div>
  );
}

export default Register;
