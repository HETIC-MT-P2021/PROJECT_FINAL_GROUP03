import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// views
import Home from "./views/home/Home";
import Auth from "./views/auth/Auth";
import Logout from "./views/logout/Logout";
import Server from "./views/server/Server";

import Dashboard from "./views/dashboard/Dashboard";
// nav

import Sidebar from "./components/sidebar/sidebar";
import AuthUtils from "./authUtil";

function App() {
  // handle tokens
  setTimeout(() => {
    let auth = new AuthUtils();
    auth.handleTokens();
  }, 1500);

  return (
    <Router>
      <div>
        <header className="App-header">
          <Sidebar></Sidebar>
        </header>
        <section className="App-body">
          <Switch>
            <Route path="/login">
              <Auth />
            </Route>
            <Route path="/logout">
              <Logout />
            </Route>
            <Route path="/dashboard">
              <Dashboard />
            </Route>
            <Route path="/servers/:id">
              <Server/>
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
