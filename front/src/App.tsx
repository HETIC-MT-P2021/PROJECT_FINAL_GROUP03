import React from "react";
import AppNavbar from "./components/navbar/Navbar";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// views
import Home from "./views/home/Home";
import Community from "./views/community/Community";

function App() {
  return (
    <Router>
      <div>
        <header className="App-header">
          <AppNavbar />
        </header>
        <section className="App-body">
          <Switch>
            <Route path="/link">
              <Community />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </section>
      </div>
    </Router>
  );
}

export default App;
