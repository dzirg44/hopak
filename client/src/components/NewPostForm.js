import React, { useState } from "react";

const NewPostForm = (props) => {
    // Declare new state variables
    const [title, content, imageUrl, setTitle, setContent, setImageUrl] = useState('');
    console.log("Title: " + title);
    console.log("Content: " + content);
    console.log("Image URL: " + imageUrl);
    return (
        <div className="row" style={{width: "70%"}}>
                <h4 className="center">Add New Post</h4>
                <form className="col s12">
                    <div className="row">
                        <div className="input-field col s12">
                            <i className="material-icons prefix">title</i>
                            <input id="title" type="text" className="validate" placeholder="Post title" value={title} onChange={e => setTitle(e.currentTarget.value)} />
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <i className="material-icons prefix">description</i>
                            <textarea placeholder="Post content" value={content} onChange={(e) => setContent(e.currentTarget.value)} />
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <i className="material-icons prefix">photo_camera</i>
                            <input id="imageUrl" name="imageUrl" type="text" className="validate" placeholder="Image URL" value={imageUrl} onChange={(e) => setImageUrl(e.currentTarget.value)}/>
                        </div>
                    </div>
                    <div className="center"><button className="waves-effect btn" type="submit" name="action"><i className="material-icons left">person_add</i>New Post</button></div>
                </form>
        </div>
    );
};

export default NewPostForm;