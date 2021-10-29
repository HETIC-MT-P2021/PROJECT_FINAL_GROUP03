import "./Home.css";
import queryString from "querystring";

function Home() {
  const params = queryString.parse(window.location.search);
  const code = params["?code"]?.toString();
  
  if (code) localStorage.setItem("auth", code);
  else if ("undefined" == localStorage.getItem("auth"))
    window.location.href = "/login";

  return <div>Home</div>;
}

export default Home;
