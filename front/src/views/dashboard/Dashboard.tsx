import {useEffect, useState} from "react";
import Server from "../../interfaces/server";
import http from "../../http.utils";
import {AxiosResponse} from "axios";
import Servers from "../../components/servers/servers";
import "./dashboard.css";

function Dashboard() {
    const [servers, setServers] = useState<Server[]>([]);
    const serversURL = process.env.REACT_APP_API_URL  + "/servers";

    const fetchServers = () => {
        http
            .get(serversURL, {
                headers: {
                    Authorization: localStorage.getItem("access_token") || ""
                }
            })
            .then((response: AxiosResponse) => {
                setServers(response.data);
            }).catch(e => {setTimeout(fetchServers, 500)});
    };

    useEffect(fetchServers, []);

    return (
        <section className="dashboard-view">
            <h1>Choix du serveur</h1>
            <Servers servers={servers} />
        </section>
    )
}

export default Dashboard;