import './App.css';
import Home from './components/Home.js'
import { BrowserRouter as Router, Route } from "react-router-dom";
import Promo from './components/Promo.js'
import Settings from './components/Settings.js'
function App() {
  return (
    <>
      <Router>
        <Route exact path="/home" component={Home} />
        <Route path="/promos" component={Promo} />
        <Route path="/settings" component={Settings} />
    </Router>
    </>
  );
}

export default App;
