import React from "react";
import "./App.css";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// views
import Home from "./views/home/Home";
import Auth from "./views/auth/Auth";

function App() {
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
