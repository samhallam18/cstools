import BubbleSort from './BubbleSort';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";
import Home from './Home';
import InsertionSort from './InsertionSort';
import BinarySearch from './BinarySearch';
import LinearSearch from './LinearSearch';

export default function App() {
    return(
        <Router>
            <div className="sidenav">
                <Link to="/">Home</Link>
                <Link to="/bubble">Bubble Sort</Link>
                <Link to="/insertion">Insertion Sort</Link>
                <Link to="/binary">Binary Search</Link>
                <Link to="/linear">Linear Search</Link>
                Useful Links:
                <a href="https://github.com/samhallam18/cstools">GitHub Repo</a>
            </div>
            <div className="main">
                <Switch>
                    <Route path="/bubble">
                        <BubbleSort />
                    </Route>
                    <Route path="/insertion">
                        <InsertionSort />
                    </Route>
                    <Route path="/binary">
                        <BinarySearch />
                    </Route>
                    <Route paht="/linear">
                        <LinearSearch />
                    </Route>
                    <Route path="/">
                        <Home />
                    </Route>
                </Switch>
            </div>
        </Router>
    )
}