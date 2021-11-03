import queryString from "querystring";
import http from "./http.utils";
import {AxiosResponse} from "axios";
import TokenInterface from "./interfaces/tokens";
import {rejects} from "assert";

// Define constants


class AuthUtils {
    private serverUrl: string;
    private clientId: string;
    private clientSecret: string;
    private port: string;

    constructor() {
        this.serverUrl = process.env.REACT_APP_DISCORD_API_URL || "";
        this.clientId = process.env.REACT_APP_CLIENT_ID || "";
        this.clientSecret = process.env.REACT_APP_CLIENT_SECRET || "";
        this.port = process.env.REACT_APP_PORT || "8080";
    }

    public handleTokens() {
        const code = localStorage.getItem("code") || "";
        console.log("code is : ", code, typeof code);
        if ("" != code && "undefined" !== code)
            this.getTokenFromCode(code);
        else this.refreshTokens();
    }

    private getTokenFromCode(code: string) {
        localStorage.setItem("code", "undefined")
        const config = {
            headers: {
                'content-type': 'application/x-www-form-urlencoded'
            },
            data: {}
        }

        http.post(this.serverUrl + "/oauth2/token", new URLSearchParams({
            client_id: this.clientId,
            client_secret: this.clientSecret,
            code,
            grant_type: 'authorization_code',
            redirect_uri: `http://localhost:${this.port}`,
            scope: 'identify,guilds',
        }), config)
        .then((response: AxiosResponse<{}>) => {
            const tokens: TokenInterface = response.data as TokenInterface;
            localStorage.setItem("access_token", tokens["access_token"]);
            localStorage.setItem("refresh_token", tokens["refresh_token"]);
            setTimeout(this.refreshTokens, tokens["expires_in"] * 1000)
        }).catch(err => {
            console.log(err)
        })
    }

    private refreshTokens() {
        const refreshToken = localStorage.getItem("refresh_token");
        if (null === refreshToken) {
            console.log("no refresh token, redirecting...")
            window.location.href = "/login";
            return;
        }

        const config = {
            headers: {
                'content-type': 'application/x-www-form-urlencoded'
            },
            data: {}
        }

        http.post(this.serverUrl + "/oauth2/token", new URLSearchParams({
            client_id: this.clientId,
            client_secret: this.clientSecret,
            refresh_token: refreshToken,
            grant_type: 'refresh_token',
            redirect_uri: `http://localhost:${this.port}`,
        }), config)
            .then((response: AxiosResponse<{}>) => {
                const tokens: TokenInterface = response.data as TokenInterface;
                localStorage.setItem("access_token", tokens["access_token"]);
                localStorage.setItem("refresh_token", tokens["refresh_token"]);
                setTimeout(this.refreshTokens, tokens["expires_in"] * 1000)
            }).catch (err => {
                console.log(err);
        })
    }
}

export default AuthUtils;