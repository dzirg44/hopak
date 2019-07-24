import React, { Component } from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import { logout } from "../store/actions/auth";

class Navbar extends Component {
    logout = e => {
        e.preventDefault();
        this.props.logout();
    };
    render() {
        return (
            <nav className="nav-wrapper">
                <a href="/" className="brand-logo">SHARE-IT</a>
                {
                this.props.currentUser.isAuthenticated ? 
                    <ul className="right">
                        <li>
                            <Link to={`/users/${this.props.currentUser.user.user_id}/posts/new`}>New Post</Link>
                        </li>
                        <li>
                            <a onClick={this.logout}>Log out</a>
                        </li>
                    </ul> :
                    <ul className="right">
                        <li>
                            <Link to="/signup">Sign up</Link>
                        </li>
                        <li>
                            <Link to="/signin">Log in</Link>
                        </li>
                    </ul>
                }
            </nav>
        );
    }
}

function mapStateToProps(state) {
    return {
        currentUser: state.currentUser
    };
}

export default connect(mapStateToProps, { logout }) (Navbar);