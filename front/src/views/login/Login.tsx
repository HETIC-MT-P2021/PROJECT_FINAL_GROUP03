import { useParams } from "react-router-dom";

const Login = () => {
  let params = {
    id: ""
  };
  params = useParams();
  localStorage.setItem("u_hash", params.id);
  window.location.assign("/");
  return (
    <>
      <h1>Product id: {params.id}</h1>
    </>
  );
};

export default Login;
