import { Navbar, Container, Nav } from "react-bootstrap";
import "./Navbar.css";

function AppNavbar() {
  return (
    <div className="Navbar">
      <Navbar bg="custom-navbar" expand="lg">
        <Container>
          <Navbar.Brand style={{color: "#EC9F05"}} href="#home">Admin panel</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link style={{color: "#EC9F05"}} href="#home">
                Home
              </Nav.Link>
              <Nav.Link style={{color: "#EC9F05"}} href="#link">Link</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </div>
  );
}

export default AppNavbar;
