import React, { useState } from "react";
import axios from "axios";
import "./CreateContainer.css"
const CreateContainer = () => {
  const [image, setImage] = useState("");
  const [name, setName] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    // Send the data to create a new container
    axios
      .post(`http://localhost:8070/container/${image}/${name}`)
      .then((response) => console.log("Container created:", response.data))
      .catch((error) => console.error("Error creating container:", error));
  };

  return (
    <div>
      <h2>Create New Container:</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Image:
          <input type="text" id="image"  value={image} onChange={(e) => setImage(e.target.value)} required />
        </label>
        <label>
          Name:
          <input type="text" id="nombre"value={name} onChange={(e) => setName(e.target.value)} required />
        </label>
        <button type="submit">Create</button>
      </form>
    </div>
  );
};

export default CreateContainer;
