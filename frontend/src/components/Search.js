import React from "react";
import "./Search.css";
import Cookies from 'universal-cookie';

const cookies = new Cookies();

function gopath(path){
    window.location = window.location.origin + path
  }

  
async function SearchByQuery(){

    let current= cookies.get('busqueda')
    let chain=''
    let a= current.split("?")
    let item = a[1];
    let b=item.split("=")
    item=b[1]
    cookies.set("busqueda_limpia", item)
    /*let c= item.split("+")

    for (let i = 0; i < c.length; i++){
        if (i!=0){
          chain =`${chain}`+`%20`+`${c[i]}`;
        }
        cookies.set("busqueda_limpia", chain)

    }*/
    gopath("/results")
}
function Search(){

   
  
  
const renderForm = (

  
    <div id="cover">
  
  <form method="get" action="">
    <div class="tb">
      <div class="td">
        <input type="text" id="search_query" placeholder="Buscar" name="search" required /></div>
      <div class="td" id="s-cover">
      <button  id="lupa" onClick={SearchByQuery(cookies.set("busqueda", window.location.search))} type="input">
         <img src="img/lupa.png"></img>
        </button>
      </div>
    </div>
  </form>
</div>
);

      return (
      <div className="app">
      <div className="search-form">
      <div class="container">
  <div class="row">
    <div class="col-md-12 text-center">
      <h3 class="animate-charcter">Bienes Raíces</h3>
      <img id="logo" src="img/tree-and-roots.png"></img>
    </div>
  </div>
</div>
      {renderForm}

      </div>
      </div>
      );
    
}export default Search; 