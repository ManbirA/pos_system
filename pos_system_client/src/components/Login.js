import React from 'react';
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { withRouter } from 'react-router';
class Login extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            usernameValue: '',
            passwordValue: ''
        };
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleUsernameChange = this.handleUsernameChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handlePasswordChange(event) {
        this.setState({
            passwordValue: event.target.value
        });
    }

    handleUsernameChange(event) {
        this.setState({
            usernameValue: event.target.value
        });
    }

    handleSubmit(event) {
        alert('Form submitted: ' + this.state.usernameValue + ' ' + this.state.passwordValue);
        // Make API call 
        // Navigate to Home page
        this.props.history.push('/home');
        event.preventDefault();
    }

    render() {
        return (
            <Form onSubmit={this.handleSubmit}>
                <Form.Group className="mb-3" controlId="formEmail">
                    <Form.Label>Username</Form.Label>
                    <Form.Control type="text" placeholder="Enter Username" onChange={this.handleUsernameChange} value={this.state.usernameValue}/>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="Password" onChange={this.handlePasswordChange} value={this.state.passwordValue}/>
                </Form.Group>
                <Form.Group>
                    <Button variant="primary" type="submit">
                        Login
                    </Button>
                </Form.Group>
            </Form>
        );
    }


}



export default withRouter(Login);