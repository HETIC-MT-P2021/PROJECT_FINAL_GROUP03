const Auth = () => {
  const redirectURL =
    "https://discord.com/api/oauth2/authorize?client_id=882205244170334218&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=code&scope=identify";

  window.location.href = redirectURL;

  return <div>Redirecting</div>;
};

export default Auth;
