import {useEffect, useState} from "react";
import Server from "../../interfaces/server";
import http from "../../http.utils";
import {AxiosResponse} from "axios";
import Servers from "../../components/servers";

function Dashboard() {
    const [servers, setServers] = useState<Server[]>([]);
    const serversURL = process.env.REACT_APP_API_URL  + "/servers";

    useEffect(() => {
        http
            .get(serversURL, {
                headers: {
                    Authorization: localStorage.getItem("access_token") || ""
                }
            })
            .then((response: AxiosResponse) => {
                console.log(response);
                setServers(response.data);
            }).catch(e => {console.log("error : ", e)});
    }, []);

    return (
        <div><Servers servers={servers} /></div>
    )
}

export default Dashboard;