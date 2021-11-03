
const Logout = () => {
    localStorage.clear();
    window.location.href = "/login";

    return (
        <div>Loging out...</div>
    )
}

export default Logout;