import React, {Component} from 'react';
import ReactDOM from 'react-dom';

class InsertionSort extends Component {
    constructor(props) {
        super(props)

        this.state = {
            items: '',
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
            method: 'POST',
        };
        config.body = new FormData();
        for (let key in this.state) {
            config.body.append(key, this.state[key])
        }
        fetch('http://127.0.0.1:3001/sorts/insertion', config)
            .then(response => {
                return response.json()
            }).then(data => {
                let stepsList = data["steps"].map(step => <tr key={step}><td>{step.join(", ")}</td></tr>)
                ReactDOM.render(
                    <div>
                        <table className="insertionTable">
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
                    document.getElementById("insertionAnswer")
                )
            }).catch(error => {
                console.log(error)
            })
    }

    render() {
        return (
            <div className="insertionSortExample">
                <h1>Insertion Sort</h1>
                <p id="insertionInfo">
                    <b>Insertion Sort</b> bla bla bla.
                </p>
                <form id="insertionInput" onSubmit={this.submitHandler}>
                    <label htmlFor="items">Enter list of numbers to sort:</label><br></br>
                    <input type="text" value={this.state.value} name="items" onChange={this.changeHandler}/><br></br>
                    <input type="submit" value="Submit" />
                </form>
                <br></br>
                <div id="insertionAnswer"></div>
            </div>
        )
    }
}

export default InsertionSort