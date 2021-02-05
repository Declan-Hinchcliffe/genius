import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
} from 'react-router-dom';
// import Footer from './Components/Footer';
// import Header from './Components/Header';
import Home from './views/Home';
import OneSong from './views/OneSong'
import AllSongs from './views/AllSongs';
import SearchSongs from './views/SearchSongs';


function App() {
  return (
    <div >
      <Router>
        <h2>What do you want to do?</h2>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/onesong">Search for one song</Link>
            </li>
            <li>
              <Link to="/allsongs">Get artists top songs</Link>
            </li>
            <li>
              <Link to="/searchsongs">Search for songs</Link>
            </li>
         </ul>
        </nav>

        <Switch>
          <Route path="/onesong">
            <OneSong />
          </Route>

          <Route path="/allsongs">
            <AllSongs />
          </Route>

          <Route path="/searchsongs">
            <SearchSongs />
          </Route>

          <Route path="/">
            <Home />
          </Route>

        </Switch>
        
      </Router>
    </div>
  );
}

export default App;
