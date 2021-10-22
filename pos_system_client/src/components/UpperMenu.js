import React from 'react';
import Navbar from 'react-bootstrap/Navbar'
import Nav from 'react-bootstrap/Nav'


class UpperMenu extends React.Component {
    render() {
        return (
            <>
                <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
                        <Navbar.Brand href="/home">POS System</Navbar.Brand>
                        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                        <Navbar.Collapse id="responsive-navbar-nav">

                            <Nav className="me-auto">
                            <Nav.Link href="/promos">Promos</Nav.Link>
                            <Nav.Link href="/settings">Settings</Nav.Link>
                            </Nav>
                        </Navbar.Collapse>
                </Navbar>
            </>
        );
    }
}

export default UpperMenu;
