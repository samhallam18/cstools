import BubbleSort from './BubbleSort';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";
import Home from './Home';

export default function App() {
    return(
        <Router>
            <div class="sidenav">
                <Link to="/">Home</Link>
                <Link to="/bubble">Bubble Sort</Link>
            </div>
            <div class="main">
                <Switch>
                    <Route path="/bubble">
                        <BubbleSort />
                    </Route>
                    <Route path="/">
                        <Home />
                    </Route>
                </Switch>
            </div>
        </Router>
    )
}