import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import authUtils from "./authUtil";

// views
import Home from "./views/home/Home";
import Auth from "./views/auth/Auth";
import Dashboard from "./views/dashboard/Dashboard";
import AuthUtils from "./authUtil";

function App() {
  // handle tokens
  let auth = new AuthUtils();
  auth.handleTokens();

  return (
    <Router>
      <div>
        <header className="App-header">
        </header>
        <section className="App-body">
          <Switch>
            <Route path="/login">
              <Auth />
            </Route>
            <Route path="/dashboard">
              <Dashboard />
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
