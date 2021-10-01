import axios from "axios";
export default axios.create({
  headers: {
    "Content-type": "application/json",
    Authorization: "Bearer " + localStorage.getItem("u_hash")
  }
});
