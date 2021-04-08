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
        alert("A form was submitted: " + this.state)
        e.preventDefault()
        console.log(JSON.stringify(this.state))
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
                let filterData = [];
                filterData = data["steps"].filter(function(item, pos, self) {
                    console.log(item)
                    return self.indexOf(item) === pos;
                })
                let stepsList = filterData.map(step => <li key={step}>{step}</li>)
                ReactDOM.render(
                    <div>
                        <ul>
                            {stepsList}
                        </ul>
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
                    <label>Enter list of numbers to sort:
                        <input type="text" value={this.state.value} name="items" onChange={this.changeHandler} />
                    </label>
                    <input type="submit" value="Submit" />
                </form>
                <div id="bubbleAnswer"></div>
            </div>
        )
    }
}

export default BubbleSort