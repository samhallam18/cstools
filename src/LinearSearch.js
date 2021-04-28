import React, {Component} from 'react';
import ReactDOM from 'react-dom';

class LinearSearch extends Component {
    constructor(props) {
        super(props)

        this.state = {
            items: '',
            searchValue: ''
        }

        this.changeHandler = this.changeHandler.bind(this)
        this.submitHandler = this.submitHandler.bind(this)
    }

    changeHandler = (e) => {
        this.setState({[e.target.name]: e.target.value})
    }

    submitHandler = (e) => {
        e.preventDefault()
        let config = {
            method: 'POST'
        }
        config.body = new FormData()
        for (let key in this.state) {
            config.body.append(key, this.state[key])
        }
        fetch('http://127.0.0.1/searches/linear', config)
            .then(response => {
                return response.json()
            }).then(data => {
                let stepsList = data["steps"].map(step => <tr key={step}><td>{step}</td></tr>)
                ReactDOM.render(
                    <div>
                        <table className="linearTable">
                            <thead>
                                <tr>
                                    <th>Steps</th>
                                </tr>
                            </thead>
                            <tbody>
                                {stepsList}
                            </tbody>
                        </table>
                    </div>,
                    document.getElementById("linearAnswer")
                )
            }).catch(error => {
                console.log(error)
            })
    }

    render() {
        return (
            <div className="linearSearchExample">
                <h1>Linear Search</h1>
                <p id="linearInfo">
                    <b>Linear Search</b> bla bla bla.
                </p>
                <form id="linearSearch" onSubmit={this.submitHandler}>
                    <label for="items">Enter a list of numbers:</label><br/>
                    <input type="text" value={this.state.value} name="items" onChange={this.changeHandler} /><br/>
                    <label>Enter the number you would like to search for: </label><br/>
                    <input type="text" value={this.state.value} name="searchValue" onChange={this.changeHandler}></input><br/>
                    <input type="submit" value="Submit"></input>
                </form>
                <div id="linearAnswer"></div>
            </div>
        )
    }
}

export default LinearSearch