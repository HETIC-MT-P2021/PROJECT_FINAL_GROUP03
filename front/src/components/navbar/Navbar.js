import { Navbar, Container, Nav } from "react-bootstrap";
import "./Navbar.css";

function AppNavbar() {
  return (
    <div className="Navbar">
      <Navbar bg="custom-navbar" expand="lg">
        <Container>
          <Navbar.Brand style={{color: "white"}} href="/">Admin panel</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link style={{color: "white"}} href="/">
                Home
              </Nav.Link>
              <Nav.Link style={{color: "white"}} href="link">Link</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </div>
  );
}

export default AppNavbar;
