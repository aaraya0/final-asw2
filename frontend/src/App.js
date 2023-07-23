import React from 'react';
import './App.css';
import Search from './components/Search';
import Results from './components/Results';
import NavBar from './components/NavBar';
import Login from './components/Login'
import Register from './components/Register'

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
function App() {
  return (
    <>
    <Router>
    <NavBar/>
    <Routes>
    <Route exact path="/" element={<Login/>}/>
    <Route exact path="/home" element={<Search/>}/>
    <Route exact path="/results" element={<Results/>}/>
    <Route exact path="/register" element={<Register/>}/>

    </Routes>
 
    
    
    </Router>
      
  
    </>
  );
}

export default App;