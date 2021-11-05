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
    const [forbiddenWordsErrorMessageText, setForbiddenWordsErrorMessageText] = useState("");

    let params = {
        id: ""
    };
    params = useParams();
    const serverResourceUrl = process.env.REACT_APP_API_URL + "/servers/" + params.id;
    let counter = 0;

    const fetchServer = () => {
        if (server?.discord_id != undefined) {return}
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
                if (counter < 5)
                    setTimeout(fetchServer, 1000);
            });
        counter++;
    }

    useEffect(fetchServer, [serverResourceUrl, fetchServer]);

    const changeWelcomeMessage = async (newMessage: string) => {
        if (newMessage === (server?.welcome_message || "")) return;
        let response = await http.patch(serverResourceUrl, {
            "welcome_message": newMessage
        });

        setErrorMessageText(
            response.status === 200
                ? "nouveau message sauvegardé"
                : "erreur lors du changement de message"
        );

        setTimeout(() => {
            setErrorMessageText("");
        }, 2000);
    };

    const changeForbiddenWords = async (forbiddenWords: string) => {
        let response = await http.patch(serverResourceUrl, {
            "forbidden_words": forbiddenWords
        });

        setForbiddenWordsErrorMessageText(
            response.status === 200
                ? "nouvelle liste sauvegardée"
                : "erreur lors du changement de liste de mots interdits"
        );

        setTimeout(() => {
            setForbiddenWordsErrorMessageText("");
        }, 2000);
    }

    if (undefined === server?.name) {
        return (
            <div className="server-view">Loading server</div>
        )
    }

    return (
        <section className="server-view">
            <h1>{ server.name }</h1>
            <br/>
            <WelcomeMessageForm
                welcome_message={server.welcome_message}
                onvalidate={changeWelcomeMessage}
            />
            <p>{errorMessageText}</p>
            <br/>
            <ForbiddenWordsForm forbidden_words={server.forbidden_words.split(",")} onvalidate={changeForbiddenWords} />
            <p>{forbiddenWordsErrorMessageText}</p>
        </section>
    )
};

export default Server;