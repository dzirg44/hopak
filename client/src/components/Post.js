import React from "react";

const Post = (props) => (
    <div className="row post">
        <div className="col s6 post-image" style={{backgroundImage: `url(${props.imageUrl})`}}/>
        <div className="col s6 post-content">
            <h2 className="title">{ props.title }</h2>
            <p className="content">{props.content}</p>
            <span>Date: {props.createdTime}</span>
        </div>
    </div>
);

export default Post;