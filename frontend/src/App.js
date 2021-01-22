import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom';
import Footer from './Components/Footer';
import Header from './Components/Header';
import Songs from './views/Songs'


function App() {
  return (
    <div className="relative pb-10 min-h-screen">
      <Router>
        <Header/>

        <Route path="/songs/lyrics/:name">
          <Songs />
        </Route>
        
        <Footer/>
      </Router>
    </div>
  );
}

export default App;
