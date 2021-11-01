import "./Home.css";
import queryString from "querystring";
import http from "../../http.utils";
import {useEffect, useState} from "react";
import TokenInterface from "../../interfaces/tokens";
import {AxiosResponse} from "axios";

function Home() {
  const params = queryString.parse(window.location.search);
  let code = params["?code"]?.toString();
  const [token, setToken] = useState<TokenInterface>();

  const serverURL = process.env.REACT_APP_DISCORD_API_URL;
  const clientId = process.env.REACT_APP_CLIENT_ID || "";
  const clientSecret = process.env.REACT_APP_CLIENT_SECRET || "";
  const port = process.env.REACT_APP_PORT || "8080";

  if (code) localStorage.setItem("auth", code);
  else if ("undefined" === localStorage.getItem("auth"))
    window.location.href = "/login";
  else code = localStorage.getItem("auth") || ""

    const getDiscordTokens = () => {
        const config = {
            headers: {
                'content-type': 'application/x-www-form-urlencoded'
            },
            data: {}
        }

        http.post(serverURL + "/oauth2/token", new URLSearchParams({
            client_id: clientId,
            client_secret: clientSecret,
            code,
            grant_type: 'authorization_code',
            redirect_uri: `http://localhost:${port}`,
            scope: 'identify,guilds',
        }), config)
            .then((response: AxiosResponse<{}>) => {
                const tokens : TokenInterface = response.data as TokenInterface;
                localStorage.setItem("access_token", tokens["access_token"]);
                localStorage.setItem("refresh_token", tokens["refresh_token"]);
                setTimeout(() => getDiscordTokens(), tokens["expires_in"])
            })
    }

    useEffect(() => {
      getDiscordTokens();
  }, [getDiscordTokens]);


  return <div>Home</div>;
}

export default Home;
