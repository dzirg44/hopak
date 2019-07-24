import { apiCall } from "../../services/api";
import { addError, removeError } from "./errors";
import { LOAD_POSTS, REMOVE_POST } from "../actionTypes";


export const loadPosts = posts => ({
    type: LOAD_POSTS,
    posts
});


export const fetchPosts = (user) => {
    return dispatch => {
        return apiCall("get", `/api/users/${user.user_id}/posts`)
        .then(res => {
            dispatch(loadPosts(res))
        })
        .catch(err => {
            addError(err.message);
        });
    };
};

