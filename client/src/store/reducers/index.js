import { combineReducers } from "redux"
import currentUser from "./currentUser";
import posts from "./posts";
import errors from "./errors";

const rootReducer = combineReducers({
    currentUser,
    posts,
    errors
});

export default rootReducer;