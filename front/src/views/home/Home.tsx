import http from "../../http.utils";
import Servers from "../../components/servers/Servers";
import { useEffect, useState } from "react";
import { AxiosResponse } from "axios";
import Server from "../../interfaces/Server";
import "./Home.css";

const serversURL = process.env.REACT_APP_API_URL + "/servers";
function Home() { 
  const [servers, setServers] = useState<Server[]>([]);

  useEffect(() => {
    http
      .get(serversURL, {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("u_hash")
        }
      })
      .then((response: AxiosResponse) => {
        setServers(response.data);
      });
  }, []);
  return (
    <div>
      <Servers servers={servers} />
    </div>
  );
}

export default Home;
