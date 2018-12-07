import * as React from "react";
import * as ReactDOM from "react-dom";
import * as axios from "axios";
import LoginPage from "./login.js"
import WalletLanding from "./wallet_landing.js"

const {dialog} = window.require('electron').remote;

class MainPage extends React.Component {
	constructor() {
		super();
		this.state = {
			loggedIn: false
		}

		this.toggle_logged_in = this.toggle_logged_in.bind(this);
	}

	toggle_logged_in() {
		this.setState(prevState => ({ loggedIn: !prevState.loggedIn }));
	}

	render() {
		return (
			<div class="main-container">
				{ this.state.loggedIn ? <WalletLanding/> : <LoginPage toggle_logged_in={this.toggle_logged_in}/> }
			</div>
		)
	}
}

const main = document.getElementById("container");
ReactDOM.render(
					<MainPage/>, main
				)
