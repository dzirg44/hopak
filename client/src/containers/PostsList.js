import React, { Component } from "react";
import { connect } from "react-redux";
import { fetchPosts } from "../store/actions/posts";
import jwtDecode from "jwt-decode";
import Post from "../components/Post";

class PostsList extends Component {

    componentDidMount() {
        console.log(jwtDecode(localStorage.jwtToken));
        console.log(this.props.fetchPosts(jwtDecode(localStorage.jwtToken)));
    }

    render() {
        const { posts } = this.props;
        console.log(posts)
        let postsList = posts.map(post => 
            <Post 
                key={post.post_id}
                imageUrl={post.image_url}
                title={post.title}
                content={post.content}
                createdTime={post.created_time}
            />
        );
        return postsList;
    }
}

function mapStateToProps(state) {
    return {
        posts: state.posts
    };
}

export default connect(mapStateToProps, { fetchPosts })(PostsList);