import './App.css';
import { BrowserRouter as Router, Route } from "react-router-dom";
import Promo from './components/Promo.js'
import Settings from './components/Settings.js'
import Home from './components/Home.js'
import UpperMenu from './components/UpperMenu.js'
import Login from './components/Login.js'
function App() {
  return (
    <>
      <UpperMenu></UpperMenu>
      <Router>
        <Route exact path="/" component={Login} />
        <Route path="/home" component={Home} />
        <Route path="/settings" component={Settings} />
      </Router>
    </>
  );
}

export default App;
