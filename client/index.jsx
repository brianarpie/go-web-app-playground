import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import PropTypes from 'prop-types';

import LoginForm from './authentication/LoginForm';

import "bootstrap-css-only/css/bootstrap.css";

class App extends Component {
	render() {
		return (
			<div className="container-fluid">
				<h1>Go and React!</h1>
				<LoginForm />
			</div>
		)
	}
}

ReactDOM.render(<App/>, document.getElementById('app'));
