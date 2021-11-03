import "./sidebar.css";
import {MdHome, MdLogout} from "react-icons/md";
import Server from "../../interfaces/server";

const Sidebar = () => {

    const redirectTo = (link : string) => {
        window.location.href = link;
    }

    return (
        <section className="sidebar">
            <div className="sidebar__link" onClick={() => redirectTo("/dashboard")}><MdHome color="orange" size="30px"/></div>
            <div className="sidebar__link" onClick={() => redirectTo("/logout")}><MdLogout color="orange" size="30px"/></div>
        </section>
    )
}

export default Sidebar;