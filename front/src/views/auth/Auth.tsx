const Auth = () => {
  window.location.href = process.env.REACT_APP_DISCORD_REDIRECT_URI || "";

  return <div>Redirecting</div>;
};

export default Auth;
