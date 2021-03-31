import React, {Component} from 'react';
import axios from 'axios';


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
        console.log(this.state)
        const { items } = this.state
        axios.post('http://127.0.0.1:3001/sorts/bubble', items)
            .then(response => {
                console.log(response)
                document.getElementById('bubbleAnswer').value = response
            })
            .catch(error => {
                console.log(error)
            })
    }

    render() {
        const {items} = this.state
        return (
            <div>
                <form id="bubbleInput" onSubmit={this.submitHandler}>
                    <label>Enter list of numbers to sort:</label>
                    <input 
                        id="bubbleNumbers" 
                        name="items" 
                        type="text" 
                        value={items}
                        onChange={this.changeHandler}></input>
                        <button type="submit">
                            Submit
                        </button>
                </form>
            </div>
        )
    }
}

export default BubbleSort