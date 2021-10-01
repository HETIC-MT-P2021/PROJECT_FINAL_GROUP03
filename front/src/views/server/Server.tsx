import { useParams } from "react-router";
import WelcomeMessageForm from "../../components/welcomeMessage/Form";
import "./Server.css";
import http from "../../http.utils";
import { useEffect, useState } from "react";
import { AxiosResponse } from "axios";
import ServerInterface from "../../interfaces/Server";

const Server = () => {
  const [server, setServer] = useState<ServerInterface>();
  const [errorMessageText, setErrorMessageText] = useState("");
  let params = {
    id: ""
  };
  params = useParams();
  const serverURL = process.env.REACT_APP_API_URL + "/servers/" + params.id;

  useEffect(() => {
    http
      .get(serverURL, {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("u_hash")
        }
      })
      .then((response: AxiosResponse) => {
        setServer(response.data);
      });
  }, []);

  const changeWelcomeMessage = async (newMessage: string) => {
    if (newMessage == server?.welcome_message) return;
    let response = await http.put(serverURL + "/welcome-message", {
      "welcome-message": newMessage
    });

    setErrorMessageText(
      response.status == 200
        ? "nouveau message sauvegardé"
        : "erreur lors du changement de message"
    );
    setTimeout(() => {
      setErrorMessageText("");
    }, 2000);
  };

  if (undefined == server) return <div>Loading data</div>;

  return (
    <div className="server-container">
      <h1>{server.name}</h1>
      <WelcomeMessageForm
        welcome_message={server.welcome_message}
        onvalidate={changeWelcomeMessage}
      />
      <p>{errorMessageText}</p>
    </div>
  );
};

export default Server;
