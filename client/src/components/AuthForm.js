import React, { Component } from "react";

class AuthForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            username: "",
            email: "",
            password: ""
        }
    }

    handleChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        });
    };

    handleSubmit = (e) => {
        e.preventDefault();
        const authType = this.props.buttonText === "Log in" ? "signin" : "signup"
        this.props.onAuth(authType, this.state)
        .then(() => {
            this.props.history.push("/")
        })
        .catch((err) => {
            console.log(err)
            return;
        });
    };

    render() {
        const {username, email, password} = this.state;
        const {heading, buttonText, errors, history, removeError} = this.props;

        history.listen(() => {
            removeError()
        });

        return (
            <div className="row" style={{width: "70%"}}>
                <h4 className="center">{heading}</h4>
                <form className="col s12" onSubmit={this.handleSubmit}>
                    {
                        buttonText === "Log in" ? <div></div> : (
                            <div className="row">
                                <div className="input-field col s12">
                                    <i className="material-icons prefix">account_circle</i>
                                    <input type="text" id="username" name="username" className="validate" placeholder="Username" onChange={this.handleChange} value={username} />
                                </div>
                            </div>
                        )
                    }
                    {errors.message && <div className="msg msg-error z-depth-3 scale-transition">{errors.message}</div>}
                    <div className="row">
                        <div className="input-field col s12">
                            <i className="material-icons prefix">email</i>
                            <input  id="email" name="email" type="email" className="validate" placeholder="Email" onChange={this.handleChange} value={email} />
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s12">
                            <i className="material-icons prefix">vpn_key</i>
                            <input type="password" id="password" name="password" className="validate" placeholder="Password" onChange={this.handleChange} value={password} />
                        </div>
                    </div>
                    <div className="center"><button className="waves-effect btn" type="submit" name="action"><i className="material-icons left">person_add</i>{buttonText}></button></div>
                </form>
            </div>
        );
    };
}

export default AuthForm;