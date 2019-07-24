import { LOAD_POSTS, REMOVE_POST } from "../actionTypes";


export default (state = [], action) => {
    switch (action.type) {
        case LOAD_POSTS:
            return [...action.posts];
        default:
            return state;
    }
};

