import React from "react";
import { Link } from "react-router-dom";
import PostsList from "../containers/PostsList";

const Homepage = ({ currentUser }) => {
if (!currentUser.isAuthenticated) {
    return (
        <div className="home-hero">
            <h1>Welcome to SHARE-IT</h1>
            <h3>New to SHARE-IT?</h3>
            <Link to="/signup" className="waves-effect waves-light btn" style={{textTransform: "uppercase"}}>
                Sign up here
            </Link>
        </div>
    );
}
    return (
        <div>
            <PostsList />
        </div>
    );
};

export default Homepage;