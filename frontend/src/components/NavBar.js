import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import './NavBar.css';
import image from '../tree-and-roots.png';
import Cookies from 'universal-cookie';

const cookies = new Cookies();

function logout() {
  // Eliminar las cookies de user_id y username
  cookies.remove('user_id', { path: '/' });
  cookies.remove('username', { path: '/login' });
}

function NavBar() {
  const [click, setClick] = useState(false);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    // Verificar si el usuario ha iniciado sesión
    setIsLoggedIn(cookies.get('user_id') !== undefined);
  }, []); // El array vacío [] asegura que este efecto se ejecute solo una vez al montar el componente.

  const handleClick = () => setClick(!click);
  const closeMobileMenu = () => setClick(false);

  const showButtonIfLoggedIn = () => {
    if (window.innerWidth <= 960 || !isLoggedIn) {
      setClick(false); // Asegurar que el menú esté cerrado en pantallas más pequeñas o cuando no se haya iniciado sesión.
    }
  };

  window.addEventListener('resize', showButtonIfLoggedIn);

  return (
    <>
        {!isLoggedIn && ( // Mostrar el logo y el texto si el usuario no ha iniciado sesión
        <div className="nav-links">
          <img src={image} className="imagen" alt="Logo" />
          Bienes Raíces
        </div>
      )}

      {isLoggedIn && ( // Mostrar la barra de navegación solamente si el usuario ha iniciado sesión
        <nav className="NavBar">
          <div className="navbar-container">
            <Link to="/home" className="navbar-logo">
              <img src={image} className="imagen" alt="Logo" />
            </Link>
            <div className="menu-icon" onClick={handleClick}>
              <i className={click ? 'fas fa-times' : 'fas fa-bars'} />
            </div>
            <ul id="lista" className={click ? 'nav-menu active' : 'nav-menu'}>
              <li className="nav-item">
                <Link to="/publish" className="nav-links" onClick={closeMobileMenu}>
                  Publicar
                </Link>
              </li>
              <li className="nav-item">
                <Link to="/user" className="nav-links" onClick={closeMobileMenu}>
                  Mi Cuenta
                </Link>
              </li>
              <li className="nav-item">
                <Link to="/" className="nav-links" onClick={logout}>
                  Cerrar Sesión
                </Link>
              </li>
            </ul>
          </div>
        </nav>
      )} 
    </>
  );
}

export default NavBar;
