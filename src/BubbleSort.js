import React, {Component} from 'react';
import ReactDOM from 'react-dom';

class BubbleSort extends Component {
    constructor(props) {
        super(props)

        this.state = {
            items: '',
        }
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
        fetch('http://127.0.0.1:3001/sorts/bubble', config)
            .then(response => {
                return response.json()
            }).then(data => {
                let stepsList = data["steps"].map(step => <tr key={step}><td>{step.join(", ")}</td></tr>)
                ReactDOM.render(
                    <div>
                        <table className="bubbleTable">
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
                    document.getElementById("bubbleAnswer")
                )
            }).catch(error => {
                console.log(error)
            })
    }

    render() {
        return (
            <div className="bubbleSortExample">
                <h1>Bubble Sort</h1>
                <p id="bubbleInfo">
                    <b>Bubble Sort</b> is a sorting algorithm that sorts a list of data by moving values to the end over successive passes.
                </p>
                <form id="bubbleInput" onSubmit={this.submitHandler}>
                    <label>Enter list of numbers to sort:<br></br>
                        <input type="text" value={this.state.value} name="items" onChange={this.changeHandler} /><br></br>
                    </label>
                    <input type="submit" value="Submit" />
                </form>
                <div id="bubbleAnswer"></div>
            </div>
        )
    }
}

export default BubbleSort