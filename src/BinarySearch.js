import React, {Component} from 'react';
import ReactDOM from 'react-dom';

class BinarySearch extends Component {
    constructor(props) {
        super(props)

        this.state = {
            items: '',
            searchValue: '',
        }
        this.changeHandler = this.changeHandler.bind(this)
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
        fetch('http://127.0.0.1:3001/searches/binary', config)
            .then(response => {
                return response.json()
            }).then(data => {
                ReactDOM.render(
                    <div>
                        <p>{ data["response"] }</p>
                    </div>,
                    document.getElementById("binaryAnswer")
                )
            }).catch(error => {
                console.log(error)
            })
    }

    render() {
        return (
            <div className="binarySearchExample">
                <h1>Binary Search</h1>
                <p id="binaryInfo">
                    <b>Binary Search</b> bla bla bla.
                </p>
                <form id="binarySearch" onSubmit={this.submitHandler}>
                    <label for="items">Enter an ordered list of numbers (ordering system soon):</label><br/>
                    <input type="text" value={this.state.value} name="items" onChange={this.changeHandler} /><br/>
                    <label>Enter the number you would like to search for: </label><br/>
                    <input type="text" value={this.state.value} name="searchValue" onChange={this.changeHandler}></input><br/>
                    <input type="submit" value="Submit"></input>
                </form>
                <div id="binaryAnswer"></div>
            </div>
        )
    }


}

export default BinarySearch