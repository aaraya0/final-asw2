import React, { useEffect, useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom"; // Import Link
import "./ContainerList.css"
const ContainerList = () => {
  const [containers, setContainers] = useState([]);

  useEffect(() => {
    // Fetch containers from backend
    axios
      .get("http://localhost:8070/containers")
      .then((response) => setContainers(response.data))
      .catch((error) => console.error("Error fetching containers:", error));
  }, []);

  return (
    <div>
    <Link to="/admin/create">Create Container</Link> {/* New "Create Container" button */}
      <h2>Containers:</h2>
      <ul>
        {containers.map((container) => (
          <li key={container.id}>
            <strong>Container ID:</strong> {container.id}, <strong>Image:</strong> {container.image},{" "}
            <strong>Name:</strong> {container.name}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ContainerList;
