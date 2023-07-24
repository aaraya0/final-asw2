import React from 'react';
import './App.css';
import Search from './components/Search';
import Results from './components/Results';
import NavBar from './components/NavBar';
import Login from './components/Login';
import Register from './components/Register';
import Message from './components/Message';
import Publication from './components/Publication';
import Item from './components/Item';
import ContainerList from './admin/ContainerList';
import CreateContainer from './admin/CreateContainer';

import Account from './components/Account';

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
    <Route exact path="/info" element={<Message/>}/>
    <Route exact path="/publish" element={<Publication/>}/>
    <Route exact path="/user" element={<Account/>}/>
    <Route exact path="/messages" element={<Item/>}/>
    <Route exact path="/admin" element={<ContainerList/>}/>
    <Route exact path="/admin/create" element={<CreateContainer/>}/>

    </Routes>
 
    
    
    </Router>
      
  
    </>
  );
}

export default App;