import Server from "../../interfaces/Server";
import "./Server.css";

const Servers = (props: { servers: Server[] }) => {
  const servers = props.servers;
  const redirectToServerPage = (id: string) => {
    window.location.assign("/servers/" + id)
  }
  const serverItems = servers.map((server, index) => <div className="server-card" onClick={() => redirectToServerPage(server.discord_id)} key={index}>{ server.name }</div>);
  return (
    <section className="servers-container">
      {serverItems}
    </section>
  );
}
export default Servers;