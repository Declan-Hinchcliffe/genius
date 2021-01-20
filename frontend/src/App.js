import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';
import Songs from './views/Songs'


function App() {
  return (
    <div>
      <Router>

        <Route path="/lyrics/:name">
          <Songs />
        </Route>
        
      </Router>
    </div>
  );
}

export default App;
