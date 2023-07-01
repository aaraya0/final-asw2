import React, {useEffect, useState} from "react";
import "./styles/Items.css";
import logo from "./images/logo.png"
import Cookies from "universal-cookie";

import { ITEMSHOST, USERSHOST, MESSAGESHOST, ITEMSPORT, USERSPORT, MESSAGESPORT} from "./config/config";
import Comment from "./Comment"
import usersvg from "./images/user.svg"

const Cookie = new Cookies();
const URLITEMS = `${ITEMSHOST}:${ITEMSPORT}`
const URLUSERS = `${USERSHOST}:${USERSPORT}`
const URLMESSAGES = `${MESSAGESHOST}:${MESSAGESPORT}`

function logout() {
    Cookie.set("user_id", -1, { path: "/" })
    document.location.reload()
}

async function getUserById(id) {
    return fetch(`${URLUSERS}/users/` + id, {
        method: "GET",
        headers: {
            "Content-Type": "comment/json"
        }
    }).then(response => response.json())
}


async function deleteComment(id) {
    return await fetch(`${URLMESSAGES}/messages/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        response.status > 200 ? goto("/mycomments") : alert("error deleting comment");
    })
}

async function getCommentsByUserId(id) {
    return fetch(`${URLMESSAGES}/users/${id}/messages`, {
        method: "GET",
        headers: {
            "Content-Type": "comment/json"
        }
    }).then(response => response.json())
}


function goto(path) {
    window.location = window.location.origin + path
}

function showComments(comments, titles) {
    return comments.map((comment) =>
        <div id={comment.message_id} className="comentario">
            <div className="comentario-texto">
                <Comment
                    key={comment.message_id}
                    comment={comment}
                    first_name={comment.first_name}
                />
            </div>
            
            <div className="comentario-eliminar">
                {!comment.system ? <button className="eliminar-comentario" onClick={() => deleteComment(comment.message_id)}> X </button> : void(0)}
            </div>

            <div className="comentario-donde">
                {titles[comment.item_id] ? titles[comment.item_id] : "error"}
            </div>

        </div>
    )
}



async function setComments(setUserComments, userId) {
    await getCommentsByUserId(userId).then(response => {
        response != null ? setUserComments(response) : setUserComments([]);
    })
}

async function getItemById(id, system = false){
    if (system) return {titulo:"SYSTEM"}
    return fetch(URLITEMS + "/items/" + id, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())

}


async function getTitles(comments){
    let titles = [];
    for(let i = 0; i < comments.length; i++){
        let title = await getItemById(comments[i].item_id, comments[i].system).then(response=> {
            return response.titulo
        });
        titles[comments[i].item_id] = title;
    }
    return titles
}

function MyComments() {
    const [user, setUser] = useState({});
    const [isLogged, setIsLogged] = useState(false);
    const [userComments, setUserComments] = useState([])
    const [needComments, setNeedComments] = useState(true)
    const [titles, setTitles] = useState([])

    if (Cookie.get("user_id") > -1 && !isLogged) {
        getUserById(Cookie.get("user_id")).then(response => setUser(response))
        setIsLogged(true)

    }

    const login = (
        <span>
            <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
        </span>
    )

    useEffect(() => {
        if (userComments.length <= 0 && Cookie.get("user_id") > -1 && needComments) {
            setComments(setUserComments, Cookie.get("user_id"))
        }
        setNeedComments(false);
    }, [userComments.length])
    const error = (
        <div>
            <div> No hay comentarios todavia </div>
        </div>
    )

    useEffect(() =>
    {
        if(userComments.length > 0){
            getTitles(userComments).then(response => setTitles(response))
        }
    }, [userComments.length])

    const logreg = (
        <div>
            <a id="login" onClick={() => goto("/login")}>Login</a>
            <a id="register" onClick={() => goto("/register")}>Register</a>
        </div>
    )

    const loggedout = (
        <div>
            <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
        </div>
    )

    return (
        <div className="comments">
            <div className="topnavHOME">
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> HouseHunter </p>
                {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
            </div>

            <div id="main">
                {userComments.length > 0 ? showComments(userComments, titles) : error}
            </div>
        </div>
    );
}

export default MyComments;