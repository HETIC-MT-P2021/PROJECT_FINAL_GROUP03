import {useEffect, useState} from "react";
import ServerInterface from "../../interfaces/server";
import {useParams} from "react-router-dom";
import http from "../../http.utils";
import {AxiosResponse} from "axios";
import "./server.css"
import WelcomeMessageForm from "../../components/forms/welcomeMessageForm";
import ForbiddenWordsForm from "../../components/forms/forbiddenWordsForm";


const Server = () => {
    const [server, setServer] = useState<ServerInterface>();
    const [errorMessageText, setErrorMessageText] = useState("");
    let params = {
        id: ""
    };
    params = useParams();
    const serverResourceUrl = process.env.REACT_APP_API_URL + "/servers/" + params.id;
    let counter = 0;
    const fetchServer = () => {
        http
            .get(serverResourceUrl, {
                headers: {
                    Authorization: localStorage.getItem("access_token") || ""
                }
            })
            .then((response: AxiosResponse) => {
                setServer(response.data);
            })
            .catch(e => {
                console.log(e, " retrying...");
                if (counter < 5)
                    setTimeout(fetchServer, 1000);
            });
        counter++;
    }

    useEffect(fetchServer, []);

    const changeWelcomeMessage = async (newMessage: string) => {
        if (newMessage == (server?.welcome_message || "")) return;
        let response = await http.patch(serverResourceUrl, {
            "welcome_message": newMessage
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

    if (undefined == server?.name) {
        return (
            <div className="server-view">Loading server</div>
        )
    }

    return (
        <section className="server-view">
            <h1>{ server.name }</h1>
            
            <WelcomeMessageForm
                welcome_message={server.welcome_message}
                onvalidate={changeWelcomeMessage}
            />
            <p>{errorMessageText}</p>

            <ForbiddenWordsForm forbidden_words={server.forbidden_words} onvalidate={(e) => console.log(e)} />
        </section>
    )
};

export default Server;