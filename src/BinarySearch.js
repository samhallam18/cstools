import React, {Component} from 'react';
import ReactDom from 'react-dom';

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
                
            })
    }


}