import "./Home.css";
import queryString from "querystring";

function Home() {
    const params = queryString.parse(window.location.search);
    const code = params["?code"]?.toString();
    localStorage.setItem("code", code);

    window.location.href = "/dashboard";

  return <div>Home</div>;
}

export default Home;
